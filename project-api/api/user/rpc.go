/*
@author: NanYan
*/
package user

import (
	"carrygpc.com/project-api/config"
	"carrygpc.com/project-common/discovery"
	"carrygpc.com/project-common/logs"
	loginv1 "carrygpc.com/project-user/pkg/service/login.service.v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"log"
)

var UserClient loginv1.LoginServiceClient

func InitUserRpc() {
	etcdRegister := discovery.NewResolver(config.C.EC.Addrs, logs.LG)
	resolver.Register(etcdRegister)
	conn, err := grpc.Dial(etcdRegister.Scheme()+":///project-user", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	UserClient = loginv1.NewLoginServiceClient(conn)
}
