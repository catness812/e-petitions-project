package controller

import (
	"context"
	"errors"

	"github.com/catness812/e-petitions-project/user_service/internal/models"
	"github.com/catness812/e-petitions-project/user_service/internal/pb"
	"github.com/gookit/slog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type IUserService interface {
	Create(user *models.User) error
	UpdatePasswordByEmail(user *models.User) error
	Delete(userEmail string) error
	GetUserByEmail(userEmail string) (*models.User, error)
	AddAdmin(userEmail string) error
	GetUserEmailById(userID uint) (string, error)
}

type UserController struct {
	userservice IUserService
}

func NewUserController(userService IUserService) *UserController {
	return &UserController{
		userservice: userService,
	}
}

func (ctrl *UserController) CreateUser(ctx context.Context, req *pb.UserRequest) (*wrapperspb.StringValue, error) {
	if req.Email == "" || req.Password == "" {
		return nil, errors.New("Email and Password cannot be empty")
	}

	user := &models.User{
		Email:    req.Email,
		Password: req.Password,
		Status:   req.Status,
	}

	err := ctrl.userservice.Create(user)

	if err != nil {
		slog.Errorf("Error adding user:%v", err.Error())
		return &wrapperspb.StringValue{Value: "Error adding user"}, err
	}
	slog.Info("User added successfully")
	return &wrapperspb.StringValue{Value: "User added successfully"}, nil

}

func (ctrl *UserController) UpdateUser(ctx context.Context, req *pb.UserRequest) (*wrapperspb.StringValue, error) {
	if req.Email == "" || req.Password == "" {
		return nil, errors.New("Email and Password cannot be empty")
	}
	user := &models.User{
		Email:    req.Email,
		Password: req.Password,
	}
	err := ctrl.userservice.UpdatePasswordByEmail(user)

	if err != nil {
		slog.Errorf("Error updating user: %v", err.Error())
		return &wrapperspb.StringValue{Value: "Error updating user"}, err
	}

	slog.Info("User updated successfully")
	return &wrapperspb.StringValue{Value: "User updated successfully"}, nil
}

func (ctrl *UserController) GetUserByEmail(ctx context.Context, req *pb.GetUserByEmailRequest) (*pb.GetUserByEmailResponse, error) {
	userEmail := req.GetEmail()
	user, err := ctrl.userservice.GetUserByEmail(userEmail)
	if err != nil {
		slog.Errorf("User not found: %v", err)
		return nil, status.Error(codes.NotFound, "User not found")
	}

	userResponse := &pb.GetUserByEmailResponse{
		Email:    req.Email,
		Id:       user.Id,
		Password: user.Password,
		Role:     user.Role,
	}

	slog.Info("Get User successful")
	return userResponse, nil
}

func (ctrl *UserController) GetUserEmailById(ctx context.Context, req *pb.GetUserEmailByIdRequest) (*wrapperspb.StringValue, error) {
	userId := req.Id
	userEmail, err := ctrl.userservice.GetUserEmailById(uint(userId))
	if err != nil {
		slog.Errorf("User email not found: %v", err)
		return nil, status.Error(codes.NotFound, "User email not found")
	}
	return &wrapperspb.StringValue{Value: userEmail}, nil
}

func (ctrl *UserController) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*wrapperspb.StringValue, error) {
	userEmail := req.GetEmail()

	if userEmail == "" {
		return nil, status.Error(codes.InvalidArgument, "Email field cannot be empty")
	}

	err := ctrl.userservice.Delete(userEmail)
	if err != nil {
		slog.Errorf("Couldn't delete: %v", err.Error())
		return nil, status.Error(codes.NotFound, "Couldn't delete")
	}

	return &wrapperspb.StringValue{Value: "User deleted successfully"}, nil
}

func (ctrl *UserController) AddAdmin(ctx context.Context, req *pb.AddAdminRequest) (*wrapperspb.StringValue, error) {
	userEmail := req.GetEmail()
	err := ctrl.userservice.AddAdmin(userEmail)
	if err != nil {
		slog.Errorf("Couldn't update role: %v", err.Error())
		return nil, status.Error(codes.NotFound, "Couldn't update role")
	}
	return &wrapperspb.StringValue{Value: "User role updated successfully"}, nil
}
