package user

import (
	"context"
	"errors"
	"gRPC_Assignment/model"
	pb "gRPC_Assignment/proto"
)

type UserService struct {
	users []model.User
}

func NewUserService() *UserService {

	users := []model.User{
		{1, "Steve", "LA", 1234567890, 5.8, true},
	}
	return &UserService{users: users}
}

func (s *UserService) GetUserById(ctx context.Context, req *pb.UserRequest) (*pb.User, error) {
	for _, user := range s.users {
		if user.ID == req.Id {
			return &pb.User{
				Id:      user.ID,
				Fname:   user.Fname,
				City:    user.City,
				Phone:   user.Phone,
				Height:  user.Height,
				Married: user.Married,
			}, nil
		}
	}
	return nil, errors.New("user not found")
}

func (s *UserService) GetUsersByIds(req *pb.UsersRequest, stream pb.UserService_GetUsersByIdsServer) error {
	var userResp []*pb.User

	for _, id := range req.Ids {
		for _, user := range s.users {
			if user.ID == id {
				userResp = append(userResp, &pb.User{
					Id:      user.ID,
					Fname:   user.Fname,
					City:    user.City,
					Phone:   user.Phone,
					Height:  user.Height,
					Married: user.Married,
				})
			}
		}
	}
	for i := range userResp {
		if err := stream.Send(userResp[i]); err != nil {
			return err
		}
	}
	return nil
}
