/*
@author: NanYan
*/
package user

import (
	login "carrygpc.com/project-api/api/user/user_grpc"
	"carrygpc.com/project-api/config"
	"carrygpc.com/project-common/discovery"
	"carrygpc.com/project-common/logs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"log"
)

var UserClient login.LoginServiceClient

func InitUserRpc() {
	etcdRegister := discovery.NewResolver(config.C.EC.Addrs, logs.LG)
	resolver.Register(etcdRegister)
	conn, err := grpc.Dial(etcdRegister.Scheme()+":///project-user", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	UserClient = login.NewLoginServiceClient(conn)
}
