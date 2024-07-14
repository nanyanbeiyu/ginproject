/*
@author: NanYan
*/
package user

import (
	login "carrygpc.com/project-api/api/user/user_grpc"
	"carrygpc.com/project-api/pkg/model/user"
	common "carrygpc.com/project-common"
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
	"net/http"
)

type HandlerLogin struct {
}

func New() *HandlerLogin {
	return &HandlerLogin{}
}

func (*HandlerLogin) GetCaptcha(c *gin.Context) {
	result := &common.Result{}
	mobile := c.PostForm("mobile")
	ctx := context.Background()
	captchaResp, err := UserClient.GetCaptcha(ctx, &login.CaptchaReq{Mobile: mobile})
	if err != nil {
		fromError, _ := status.FromError(err)
		c.JSON(200, result.Fail(common.BusinessCode(fromError.Code()), fromError.Message()))
		return
	}
	c.JSON(200, result.Success(captchaResp.Code))
}

func (l *HandlerLogin) register(c *gin.Context) {
	// 1.接收参数 参数模型
	result := common.Result{}
	var req user.RegisterReq
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "参数格式错误"))
		return
	}
	// 2.校验参数 判断参数是否合法
	if err = req.Verify(); err != nil {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, err.Error()))
		return
	}
	// 3.业务处理 调用grpc服务 获取响应
	ctx := context.Background()
	msg := &login.RegisterReq{
		Email:    req.Email,
		Name:     req.Name,
		Password: req.Password,
		Mobile:   req.Mobile,
		Captcha:  req.Captcha,
	}
	_, err = UserClient.Register(ctx, msg)
	// 4.返回结果
	if err != nil {
		fromError, _ := status.FromError(err)
		c.JSON(200, result.Fail(common.BusinessCode(fromError.Code()), fromError.Message()))
		return
	}
	c.JSON(200, result.Success(nil))
}
