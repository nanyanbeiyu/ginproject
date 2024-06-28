/*
@author: NanYan
@doc: 路由注册
*/
package routers

import (
	"carrygpc.com/project-user/api/user"
	"github.com/gin-gonic/gin"
)

type Router interface {
	Register(r *gin.Engine)
}

type RegisterRouter struct {
}

func New() *RegisterRouter {
	return &RegisterRouter{}
}

func (*RegisterRouter) Router(router Router, r *gin.Engine) {
	router.Register(r)
}

func Inirouter(r *gin.Engine) {
	router := New()
	// 路由注册位置
	router.Router(&user.RouterUser{}, r)
}
