/*
@author: NanYan
*/
package login_service_v1

import (
	common "carrygpc.com/project-common"
	"carrygpc.com/project-common/encrypts"
	"carrygpc.com/project-common/snowflakeID"
	"carrygpc.com/project-user/internal/dao/Transaction"
	"carrygpc.com/project-user/internal/dao/gredis"
	"carrygpc.com/project-user/internal/dao/memberDao"
	"carrygpc.com/project-user/internal/dao/organizationDao"
	"carrygpc.com/project-user/internal/data/member"
	"carrygpc.com/project-user/internal/data/organization"
	"carrygpc.com/project-user/internal/database"
	"carrygpc.com/project-user/internal/database/tran"
	"carrygpc.com/project-user/internal/repo"
	"carrygpc.com/project-user/pkg/model"
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	codes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type LoginService struct {
	UnimplementedLoginServiceServer
	cache            repo.Cache
	memberRepo       repo.MemberRepo
	organizationRepo repo.OrganizationRepo
	transaction      tran.Transaction
}

func NewLoginService() *LoginService {
	return &LoginService{
		cache:            gredis.Rc,
		memberRepo:       memberDao.NewMemberDao(),
		organizationRepo: organizationDao.NewOrganizationDao(),
		transaction:      Transaction.NewTransaction(),
	}
}

func (ls *LoginService) GetCaptcha(ctx context.Context, req *CaptchaReq) (*CaptchaResp, error) {
	//1. 获取参数
	mobile := req.Mobile
	//2. 校验参数
	if !common.VerifyMobile(mobile) {
		return nil, status.Error(codes.Code(model.NoLegalMobile), model.GetMsg(model.NoLegalMobile))
	}
	//3. 生成随机验证码（随机4位1000-9999或者6位100000-999999）
	code := common.GenerateRandomCode(6)
	//4. 调用短信平台发送验证码（三方短信平台 放入协程执行 接口可以快速响应）
	go func() {
		time.Sleep(2 * time.Second)
		zap.L().Info(fmt.Sprintf("发送验证码：REGISTER_%s : %s", mobile, code))
		//5. 存储验证码 gredis 时间15分钟
		err := ls.cache.Put("REGISTER_"+mobile, code, 15*time.Minute)
		if err != nil {
			zap.L().Error(fmt.Sprintf("将验证码存入redis失败：REGISTER_%s : %s cause by: %v", mobile, code, err))
			return
		}
		zap.L().Info(fmt.Sprintf("将验证码存入redis成功：REGISTER_%s : %s", mobile, code))
	}()
	return &CaptchaResp{Code: code}, nil
}

func (ls *LoginService) Register(ctx context.Context, req *RegisterReq) (*RegisterResp, error) {
	// 1.可以校验参数
	// 2.校验验证码
	captchaCode, err := ls.cache.Get("REGISTER_" + req.Mobile)

	if errors.Is(err, redis.Nil) {
		return nil, status.Error(codes.Code(model.CaptcgaNotExist), model.GetMsg(model.CaptcgaNotExist))
	}

	if err != nil {
		zap.L().Error("Register captchaCode get error:", zap.Error(err))
		return nil, status.Error(codes.Code(model.NoLegalCaptcha), model.GetMsg(model.NoLegalCaptcha))
	}
	if req.Captcha != captchaCode {
		return nil, status.Error(codes.Code(model.NoLegalCaptcha), model.GetMsg(model.NoLegalCaptcha))
	}
	// 3.校验业务逻辑（邮箱、用户名、手机号是否被注册）
	// 3.1 校验邮箱是否注册
	exist, err := ls.memberRepo.GetMemberEmail(ctx, req.Email)
	if err != nil {
		zap.L().Error("Register get member email error:", zap.Error(err))
		return nil, status.Error(codes.Code(model.DBError), model.GetMsg(model.DBError))
	}
	if exist {
		return nil, status.Error(codes.Code(model.EmailExist), model.GetMsg(model.EmailExist))
	}
	// 3.2 校验用户名是否注册
	exist, err = ls.memberRepo.GetMemberByName(ctx, req.Name)
	if err != nil {
		zap.L().Error("Register get member name error:", zap.Error(err))
		return nil, status.Error(codes.Code(model.DBError), model.GetMsg(model.DBError))
	}
	if exist {
		return nil, status.Error(codes.Code(model.NameExist), model.GetMsg(model.NameExist))
	}
	// 3.3 校验手机号是否注册
	exist, err = ls.memberRepo.GetMemberByMobile(ctx, req.Mobile)
	if err != nil {
		zap.L().Error("Register get member mobile error:", zap.Error(err))
		return nil, status.Error(codes.Code(model.DBError), model.GetMsg(model.DBError))
	}
	if exist {
		return nil, status.Error(codes.Code(model.MobileExist), model.GetMsg(model.MobileExist))
	}
	// 4.执行业务（将数据存入member表 生成一个数据存入organization表）
	pwd := encrypts.Md5(req.Password)
	account := snowflakeID.SnowflakeID()
	mem := &member.Member{
		Account:       account,
		Password:      pwd,
		Name:          req.Name,
		Mobile:        req.Mobile,
		Email:         req.Email,
		LastLoginTime: time.Now(),
		Status:        1,
	}
	err = ls.transaction.Action(func(conn database.DbConn) error {
		err = ls.memberRepo.SaveMember(conn, ctx, mem)
		if err != nil {
			zap.L().Error("register save member db err:", zap.Error(err))
			return status.Error(codes.Code(model.DBError), model.GetMsg(model.DBError))
		}
		org := &organization.Organization{
			Name:     mem.Name + "个人组织",
			MemberId: mem.Account,
			Personal: 1,
			Avatar:   "https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fc-ssl.dtstatic.com%2Fuploads%2Fblog%2F202103%2F31%2F20210331160001_9a852.thumb.1000_0.jpg&refer=http%3A%2F%2Fc-ssl.dtstatic.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=auto?sec=1673017724&t=ced22fc74624e6940fd6a89a21d30cc5",
		}
		err = ls.organizationRepo.SaveOrganization(conn, ctx, org)
		if err != nil {
			zap.L().Error("register save organization db err:", zap.Error(err))
			return status.Error(codes.Code(model.DBError), model.GetMsg(model.DBError))
		}
		err = ls.cache.Delete("REGISTER_" + req.Mobile)
		if err != nil {
			zap.L().Error("Register captchaCode delete error:", zap.Error(err))
			return status.Error(codes.Code(model.DBError), model.GetMsg(model.DBError))
		}
		return nil
	})

	// 5.返回
	return nil, err
}
