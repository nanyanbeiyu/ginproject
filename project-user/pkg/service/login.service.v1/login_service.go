/*
@author: NanYan
*/
package login_service_v1

import (
	common "carrygpc.com/project-common"
	"carrygpc.com/project-common/encrypts"
	"carrygpc.com/project-common/snowflakeID"
	"carrygpc.com/project-user/internal/dao/memberDao"
	"carrygpc.com/project-user/internal/dao/organizationDao"
	"carrygpc.com/project-user/internal/dao/redis"
	"carrygpc.com/project-user/internal/data/member"
	"carrygpc.com/project-user/internal/data/organization"
	"carrygpc.com/project-user/internal/repo"
	"carrygpc.com/project-user/pkg/model"
	"context"
	"fmt"
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
}

func NewLoginService() *LoginService {
	return &LoginService{
		cache:            redis.Rc,
		memberRepo:       memberDao.NewMemberDao(),
		organizationRepo: organizationDao.NewOrganizationDao(),
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
		//5. 存储验证码 redis 时间15分钟
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
		LastLoginTime: time.Now().UnixMilli(),
		Status:        1,
	}
	err = ls.memberRepo.SaveMember(ctx, mem)
	if err != nil {
		zap.L().Error("register save member db err:", zap.Error(err))
		return &RegisterResp{}, err
	}
	org := &organization.Organization{
		Name:     mem.Name + "个人组织",
		MemberId: mem.Account,
		Personal: 1,
		Avatar:   "https://profile-avatar.csdnimg.cn/30ff829b21d8422db18289abaf318685_qq_46103376.jpg!1",
	}
	err = ls.organizationRepo.SaveOrganization(ctx, org)
	if err != nil {
		zap.L().Error("register save organization db err:", zap.Error(err))
		return &RegisterResp{}, err
	}
	err = ls.cache.Delete("REGISTER_" + req.Mobile)
	if err != nil {
		zap.L().Error("Register captchaCode delete error:", zap.Error(err))
		return nil, status.Error(codes.Code(model.DBError), model.GetMsg(model.DBError))
	}
	// 5.返回
	return nil, nil
}
