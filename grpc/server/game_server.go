package server

import (
	"context"

	models "github.com/gglzc/fishMachine/models/user"
	pb "github.com/gglzc/fishMachine/proto"
	bulletServices "github.com/gglzc/fishMachine/service/bullet"
	userServices "github.com/gglzc/fishMachine/service/user"
)

type GameServer struct {
	pb.UnimplementedGameServiceServer
	BulltetService *bulletServices.GameService
	UserService    *userServices.UserService
}

func (s *GameServer) ShootBullet(ctx context.Context, req *pb.ShootRequest) (*pb.ShootResponse, error) {
	err := s.BulltetService.ShootBullet(req.UserId, req.Cost)
	if err != nil {
		return &pb.ShootResponse{
			Captured: false,
			Reward:   0,
			Error:    err.Error(),
		}, err
	}
	return &pb.ShootResponse{
		Captured: true,
		Reward:   req.Cost,
		Error:    "",
	}, nil
}
func (s *GameServer) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	user, err := s.UserService.GetUserByID(req.UserId)
	if err != nil {
		return &pb.UserResponse{
			Error: err.Error(),
		}, nil
	}
	return &pb.UserResponse{
		Id:       user.ID,
		Username: user.Username,
		Balance:  user.Balance,
		Error:    "",
	}, nil
}
func (s *GameServer)CreateUser(ctx context.Context , req *pb.CreateUserRequest)(*pb.CreateUserResponse,error){
	err := s.UserService.CreateUser(models.User{
		ID:       1,
		Username: req.Username,
		Password: req.Password,
		Balance:  1000000000,
	})
	if err != nil {
		return &pb.CreateUserResponse{
			Status: "Fail Create User",
			Error: err.Error(),
		}, nil
	}
	// 返回成功响应
	return &pb.CreateUserResponse{
		Status: "Success Create User",
		Error: "",
	}, nil
}
