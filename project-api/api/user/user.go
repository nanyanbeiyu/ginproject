/*
@author: NanYan
*/
package user

import (
	common "carrygpc.com/project-common"
	loginv1 "carrygpc.com/project-user/pkg/service/login.service.v1"
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
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
	_, err := UserClient.GetCaptcha(ctx, &loginv1.CaptchaReq{Mobile: mobile})
	if err != nil {
		fromError, _ := status.FromError(err)
		c.JSON(200, result.Fail(common.BusinessCode(fromError.Code()), fromError.Message()))
		return
	}
	c.JSON(200, result.Success(nil))
}
