package api

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/MRaihanZ/Simple-Microservice/models"
	pb "github.com/MRaihanZ/Simple-Microservice/proto/user"
	"google.golang.org/grpc"
)

// type server struct {
// 	pb.UnimplementedUserServiceServer
// }

//	func ConvertToProtoUsers(users []models.User) []*pb.User {
//		var protoUsers []*pb.User
//		for _, u := range users {
//			protoUsers = append(protoUsers, &pb.User{
//				Name:     u.Name,
//				Email:    u.Email,
//				Password: u.Password,
//			})
//		}
//		return protoUsers
//	}
func (*server) CreateLog(ctx context.Context, empty *pb.Empty) (*pb.Response, error) {
	fmt.Println("Getting users...")
	result, err := models.CreateLog()
	if err != nil {
		fmt.Println("Error:", err)
		return &pb.Response{Code: 500, Body: nil}, err
	}
	protoUsers := ConvertToProtoUsers(result)
	return &pb.Response{Code: 200, Body: protoUsers}, nil
}

func Run() {

	fmt.Println("Go gRPC Beginners Tutorial!")

	lis, err := net.Listen("tcp", "127.0.0.1:9001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, &server{})
	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
