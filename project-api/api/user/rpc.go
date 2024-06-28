/*
@author: NanYan
*/
package user

import (
	loginv1 "carrygpc.com/project-user/pkg/service/login.service.v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var UserClient loginv1.LoginServiceClient

func InitUserRpc() {
	conn, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	UserClient = loginv1.NewLoginServiceClient(conn)
}
