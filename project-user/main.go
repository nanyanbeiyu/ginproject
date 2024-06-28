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
	//从配置中读取日志配置，初始化日志
	config.C.InitZapLog()
	routers.Inirouter(r)
	srv.Run(r, config.C.ServeConf.Host, config.C.ServeConf.Port)
}
