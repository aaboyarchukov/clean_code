package lesson2

import (
	"context"
	"fmt"
	"slices"
	auth_v1 "system_of_monitoring_statistics/services/auth/gen/protos"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Auth interface {
	Login(ctx context.Context,
		email string,
		password string,
	) (string, error)
	Register(ctx context.Context,
		email string,
		password string,
		name string,
		surname string,
		sex string,
		height int64,
		weight int64,
		birth_date int64,
	) (int64, error)
}

type authServerApi struct {
	auth_v1.UnimplementedAuthServer
	auth Auth
}

func RegisterServer(gRPC *grpc.Server, auth Auth) {
	auth_v1.RegisterAuthServer(gRPC, &authServerApi{
		auth: auth,
	})
}

func ValidateStringData(args ...string) error {
	if slices.Contains(args, "") {
		return fmt.Errorf("string data is empty")
	}
	return nil
}

func (s *authServerApi) Login(
	ctx context.Context,
	in *auth_v1.LoginRequest,
) (*auth_v1.LoginResponse, error) {

	// validate data
	// 6.3
	errValidate := ValidateStringData(in.GetEmail(), in.GetPassword())
	if errValidate != nil {
		return nil, status.Error(codes.InvalidArgument, "email and password are required")
	}

	// 6.3
	token, errLogin := s.auth.Login(ctx, in.GetEmail(), in.GetPassword())
	if errLogin != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid email or password")
	}

	return &auth_v1.LoginResponse{
		Token: token,
	}, nil
}

func (s *authServerApi) Register(ctx context.Context, in *auth_v1.RegisterRequest) (*auth_v1.RegisterResponse, error) {
	// validate data
	// 6.3
	errValidate := ValidateStringData(in.GetEmail(), in.GetPassword(), in.GetName())
	if errValidate != nil {
		return nil, status.Error(codes.InvalidArgument, "email, password and name are required")
	}

	// 6.3
	userID, errRegister := s.auth.Register(ctx, in.GetEmail(), in.GetPassword(), in.GetName(),
		in.GetSurname(), in.GetSex(), in.GetHeight(), in.GetWeight(), in.GetDateBirth())
	if errRegister != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid email or password")
	}
	return &auth_v1.RegisterResponse{
		UserId: userID,
	}, nil
}
