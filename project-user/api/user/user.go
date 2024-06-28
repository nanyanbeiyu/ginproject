/*
@author: NanYan
*/
package user

import (
	common "carrygpc.com/project-common"
	"carrygpc.com/project-user/pkg/dao/redis"
	"carrygpc.com/project-user/pkg/model"
	"carrygpc.com/project-user/pkg/repo"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type HandlerUser struct {
	cache repo.Cache
}

func New() *HandlerUser {
	return &HandlerUser{
		cache: redis.Rc,
	}
}

func (*HandlerUser) getCaptcha(c *gin.Context) {
	h := New()
	rsp := &common.Result{}
	//1. 获取参数
	mobile := c.PostForm("mobile")
	//2. 校验参数
	if !common.VerifyMobile(mobile) {
		rsp.Fail(model.NoLegalMobile, model.GetMsg(model.NoLegalMobile))
		c.JSON(http.StatusOK, rsp)
		return
	}
	//3. 生成随机验证码（随机4位1000-9999或者6位100000-999999）
	code := common.GenerateRandomCode(6)
	//4. 调用短信平台发送验证码（三方短信平台 放入协程执行 接口可以快速响应）
	go func() {
		time.Sleep(2 * time.Second)
		zap.L().Info(fmt.Sprintf("发送验证码：REGISTER_%s : %s", mobile, code))
		//5. 存储验证码 redis 时间15分钟
		err := h.cache.Put("REGISTER_"+mobile, code, 15*time.Minute)
		if err != nil {
			zap.L().Error(fmt.Sprintf("将验证码存入redis失败：REGISTER_%s : %s cause by: %v", mobile, code, err))
			return
		}
		zap.L().Info(fmt.Sprintf("将验证码存入redis成功：REGISTER_%s : %s", mobile, code))
	}()
	c.JSON(200, rsp.Success(model.GetMsg(model.VerifyCodeSuccess)))
}
