package service

import (
	"auth_svc/pkg/pb"
	"auth_svc/pkg/repository/interfaces"
	"context"
	"net/http"
)

type authService struct {
	Repo interfaces.AuthRepo
	pb.UnimplementedAuthServiceServer
}

func NewAuthService(repo interfaces.AuthRepo) pb.AuthServiceServer {
	return &authService{
		Repo: repo,
	}
}

func (cr *authService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	userId, err := cr.Repo.Register(context.Background(), req)
	if err != nil {
		return &pb.RegisterResponse{
			Status: http.StatusInternalServerError,
		}, err
	}
	return &pb.RegisterResponse{
		Status:   http.StatusOK,
		Response: "SignUp Success",
		ID:       userId,
	}, nil
}

func (cr *authService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, err := cr.Repo.GetUser(context.Background(), req.ID)
	if err != nil {
		return &pb.GetUserResponse{
			Status: http.StatusInternalServerError,
		}, err
	}

	data := &pb.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
	}
	return &pb.GetUserResponse{
		Status: http.StatusOK,
		User:   data,
	}, nil
}

func (cr *authService) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	user, err := cr.Repo.UpdateUser(context.Background(), req)
	if err != nil {
		return &pb.UpdateResponse{
			Status: http.StatusInternalServerError,
		}, err
	}
	data := &pb.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
	}
	return &pb.UpdateResponse{
		Status:   http.StatusOK,
		Response: "Successfully updated",
		User:     data,
	}, nil
}

func (cr *authService) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	err := cr.Repo.DeleteUser(context.Background(), req.ID)
	if err != nil {
		return &pb.DeleteResponse{
			Status: http.StatusInternalServerError,
		}, err
	}
	return &pb.DeleteResponse{
		Status:   http.StatusOK,
		Response: "Successfully deleted",
	}, nil
}

func (cr *authService) GetAll(ctx context.Context, req *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	users, err := cr.Repo.GetAll(context.Background(), req)
	if err != nil {
		return &pb.GetAllResponse{
			Status: http.StatusInternalServerError,
		}, err
	}

	// Convert FirstNames slice to []*pb.MethodUser
	var methodUsers []*pb.MethodUser
	for _, firstName := range users.FirstNames {
		methodUser := &pb.MethodUser{
			FirstName: firstName,
		}
		methodUsers = append(methodUsers, methodUser)
	}
	return &pb.GetAllResponse{
		Status:   http.StatusOK,
		Response: "Success",
		Count:    users.Count,
		Users:    methodUsers,
	}, nil
}
