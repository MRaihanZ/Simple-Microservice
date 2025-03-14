package entities

import pb "github.com/MRaihanZ/Simple-Microservice/proto/user"

type UserService struct {
	pb.UnimplementedUserServiceServer
}
