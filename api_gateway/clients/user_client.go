package clients

import (
	"context"
	"log"

	"github.com/MRaihanZ/Simple-Microservice/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ErrorMsg() {
	log.Fatalf("Error Happen in user_client")
	msg := recover()
	log.Fatalf("ERROR MESSAGE: %s", msg)
}

func ConnectUserService() (*grpc.ClientConn, user.UserServiceClient) {
	defer ErrorMsg()
	conn, err := grpc.NewClient("127.0.0.1:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client := user.NewUserServiceClient(conn)
	return conn, client
}

func GetUsers(conn *grpc.ClientConn, client user.UserServiceClient) *user.Response {
	defer ErrorMsg()
	defer conn.Close()
	response, err := client.GetUsers(context.Background(), &user.Empty{})
	if err != nil {
		panic(err)
	}
	return response
}

// func CreateUsers(conn *grpc.ClientConn, client user.UserServiceClient, name string, email string, password string) *user.Response {
// 	defer ErrorMsg()
// 	defer conn.Close()
// 	response, err := client.CreateUser(context.Background(), &user.User{Name: name, Email: email, Password: password})
// 	if err != nil {
// 		panic(err)
// 	}
// 	return response
// }
