/*
@author: NanYan
*/
package main

import (
	srv "carrygpc.com/project-common"
	"carrygpc.com/project-user/config"
	"carrygpc.com/project-user/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routers.Initrouter(r)
	// 注册grpc服务
	grpc := routers.RegisterGrpc()
	// 注册etcd服务
	routers.RegisterEtcdServer()
	stop := func() {
		grpc.Stop()
	}
	srv.Run(r, config.C.ServeConf.Host, config.C.ServeConf.Port, stop)
}
