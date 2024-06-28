/*
@author: NanYan
*/
package user

import "github.com/gin-gonic/gin"

type RouterUser struct {
}

func (*RouterUser) Register(r *gin.Engine) {
	h := &HandlerUser{}
	r.POST("/project/login/getCaptcha", h.getCaptcha)
}
