/*
@author: NanYan
*/
package login_service_v1

import (
	common "carrygpc.com/project-common"
	"carrygpc.com/project-user/pkg/dao/redis"
	"carrygpc.com/project-user/pkg/model"
	"carrygpc.com/project-user/pkg/repo"
	"context"
	"fmt"
	"go.uber.org/zap"
	codes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type LoginService struct {
	UnimplementedLoginServiceServer
	cache repo.Cache
}

func NewLoginService() *LoginService {
	return &LoginService{
		cache: redis.Rc,
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
	return &CaptchaResp{}, nil
}
