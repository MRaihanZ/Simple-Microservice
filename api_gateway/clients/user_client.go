package clients

import (
	"context"
	"fmt"
	"log"

	pb "github.com/MRaihanZ/Simple-Microservice/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ConnectUserService() (*grpc.ClientConn, pb.UserServiceClient) {
	fmt.Println("Connecting to user service...")
	conn, err := grpc.NewClient("127.0.0.1:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	fmt.Println("Connected to user service")

	client := pb.NewUserServiceClient(conn)
	return conn, client
}

func GetUsers(conn *grpc.ClientConn, client pb.UserServiceClient) *pb.Response {
	fmt.Println("Getting users...")
	defer conn.Close()
	response, err := client.GetUsers(context.Background(), &pb.Empty{})
	if err != nil {
		log.Fatalf("error getting data Users: %s", err)
	}
	return response
}

// func CreateUsers(conn *grpc.ClientConn, client user.UserServiceClient, name string, email string, password string) *user.Response {
// 	defer conn.Close()
// 	response, err := client.CreateUser(context.Background(), &user.User{Name: name, Email: email, Password: password})
// 	if err != nil {
// 		panic(err)
// 	}
// 	return response
// }
