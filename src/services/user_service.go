package services

import (
	"context"
	"database/sql"
	"log"

	v1 "github.com/kalpg69/backend-admin/src/api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	queryCreateUpdateDeleteUser = "CreateUpdateDeleteUser"
	queryGetUser                = "SELECT UserName, [Password], Role FROM Users WHERE UserName = @UserName"
)

type userServiceServer struct {
	v1.UnimplementedUserServiceServer
	db *sql.DB
}

func NewUserServiceServer(db *sql.DB) v1.UserServiceServer {
	return &userServiceServer{db: db}
}

func (u *userServiceServer) CreateUser(ctx context.Context, req *v1.CreateUserRequest) (*v1.CreateUserResponse, error) {
	if apiVersion != req.GetApi() {
		log.Printf("Error: API version for Create User request not implemented")
		return nil, status.Error(codes.Unimplemented, "Error: API version not supported")
	}

	newUser := req.GetUser()

	_, err := u.db.ExecContext(ctx, queryCreateUpdateDeleteUser,
		sql.Named("UserName", newUser.GetUsername()),
		sql.Named("Password", newUser.GetPassword()),
		sql.Named("Role", newUser.GetRole()),
		sql.Named("OperationType", createOperationType))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error: Failed to create user")
	}

	return &v1.CreateUserResponse{
		Response: &v1.Response{
			Api:           "v1",
			Status:        "OK",
			StatusMessage: "User created successfully",
		},
	}, nil
}

func (u *userServiceServer) UpdateUser(ctx context.Context, req *v1.UpdateUserRequest) (*v1.UpdateUserResponse, error) {
	return nil, nil
}

func (u *userServiceServer) DeleteUser(ctx context.Context, req *v1.DeleteUserRequest) (*v1.DeleteUserResponse, error) {
	return nil, nil
}

func (u *userServiceServer) GetUser(ctx context.Context, req *v1.GetUserRequest) (*v1.GetUserResponse, error) {

	var currentUser v1.User
	if err := u.db.QueryRowContext(ctx, queryGetUser, sql.Named("UserName", req.GetUsername())).Scan(&currentUser.Username, &currentUser.Password, &currentUser.Role); err != nil {
		return nil, status.Errorf(codes.NotFound, "Error: User not found")
	}

	return &v1.GetUserResponse{
		Api: "v1",
		User: &v1.User{
			Username: currentUser.Username,
			Password: currentUser.Password,
			Role:     currentUser.Role,
		},
	}, nil
}
