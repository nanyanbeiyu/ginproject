/*
@author: NanYan
@doc: 路由注册
*/
package routers

import (
	"carrygpc.com/project-common/discovery"
	"carrygpc.com/project-common/logs"
	"carrygpc.com/project-user/api/user"
	"carrygpc.com/project-user/config"
	login_service_v1 "carrygpc.com/project-user/pkg/service/login.service.v1"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"log"
	"net"
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

func Initrouter(r *gin.Engine) {
	router := New()
	// 路由注册位置
	router.Router(&user.RouterUser{}, r)
}

type GcConfig struct {
	Addr         string
	RegisterFunc func(server *grpc.Server)
}

func RegisterGrpc() *grpc.Server {
	c := GcConfig{
		Addr: config.C.GC.Addr,
		RegisterFunc: func(server *grpc.Server) {
			login_service_v1.RegisterLoginServiceServer(server, login_service_v1.NewLoginService())
		},
	}

	s := grpc.NewServer()
	c.RegisterFunc(s)
	lis, err := net.Listen("tcp", c.Addr)
	if err != nil {
		log.Println("cannot listen")
	}
	go func() {
		err = s.Serve(lis)
		if err != nil {
			log.Println("server started error", err)
			return
		}
	}()
	return s
}

func RegisterEtcdServer() {
	etcdRegister := discovery.NewResolver(config.C.EC.Addrs, logs.LG)
	resolver.Register(etcdRegister)
	info := discovery.Server{
		Name:    config.C.GC.Name,
		Addr:    config.C.GC.Addr,
		Version: config.C.GC.Version,
		Weight:  config.C.GC.Weight,
	}
	r := discovery.NewRegister(config.C.EC.Addrs, logs.LG)
	_, err := r.Register(info, 2)
	if err != nil {
		log.Fatalln(err)
	}
}
