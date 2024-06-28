/*
@author: NanYan
*/
package main

import (
	"carrygpc.com/project-api/config"
	"carrygpc.com/project-api/routers"
	srv "carrygpc.com/project-common"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routers.InitRouter(r)
	srv.Run(r, config.C.ServeConf.Name, config.C.ServeConf.Addr, nil)
}
