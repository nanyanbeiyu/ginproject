/*
@author: NanYan
*/
package user

import "github.com/gin-gonic/gin"

type RouterLogin struct {
}

func (*RouterLogin) Register(r *gin.Engine) {
	InitUserRpc()
	h := New()
	r.POST("/project/login/getCaptcha", h.GetCaptcha)
	r.POST("/project/login/register", h.register)
}
