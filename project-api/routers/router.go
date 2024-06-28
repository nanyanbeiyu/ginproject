/*
@author: NanYan
*/
package routers

import (
	"carrygpc.com/project-api/api/user"
	"github.com/gin-gonic/gin"
)

// Router 接口
type Router interface {
	Register(engine *gin.Engine)
}

type RegisterRouter struct {
}

func New() *RegisterRouter {
	return &RegisterRouter{}
}

func (*RegisterRouter) Router(router Router, r *gin.Engine) {
	router.Register(r)
}

func InitRouter(r *gin.Engine) {
	router := New()
	router.Router(&user.RouterLogin{}, r)
}
