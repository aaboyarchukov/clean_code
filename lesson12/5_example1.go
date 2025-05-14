package lesson12

func (bf *BloomFilter) Hash1(s string) int {
	sum := 0
	const HASH_1_KOEFF int = 17
	for _, char := range s {
		code := int(char)
		sum += code * HASH_1_KOEFF
	}
	sum %= bf.filter_len
	return sum
}

func (bf *BloomFilter) Hash2(s string) int {
	sum := 0
	const HASH_2_KOEFF int = 223
	for _, char := range s {
		code := int(char)
		sum += code * HASH_2_KOEFF
	}
	sum %= bf.filter_len
	return sum
}

package main

import (
	"log/slog"
	"os"
	auth_app "system_of_monitoring_statistics/services/auth/app"
	auth_config "system_of_monitoring_statistics/services/auth/config"

	"github.com/joho/godotenv"
)

func main() {
	// load env variables
	MustSetUpEnv()

	// TODO: инициализировать объект loger
	log := SetupLoger("local")

	// инициализировать config для auth
	authConfig := auth_config.MustLoad()
	// инициализировать auth
	authApllication := auth_app.New(log, authConfig.GRPC.Port)
	authApllication.GRPCServer.MustRun()
	// TODO: все упаковать в gorutines и добавить канал,
	// который ожидает сигнала по завершению
	// затем GracefulShotDowmn
}

func MustSetUpEnv() {
	if err := godotenv.Load(".env"); err != nil {
		panic("failed with load env file")
	}
}

func SetupLoger(enviroment string) *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
}

func (s *authServerApi) Login(
	ctx context.Context,
	in *auth_v1.LoginRequest,
) (*auth_v1.LoginResponse, error) {

	// validate data
	errValidate := ValidateStringData(in.GetEmail(), in.GetPassword())
	if errValidate != nil {
		return nil, status.Error(codes.InvalidArgument, "email and password are required")
	}

	token, errLogin := s.auth.Login(ctx, in.GetEmail(), in.GetPassword())
	if errLogin != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid email or password")
	}

	return &auth_v1.LoginResponse{
		Token: token,
	}, nil
}
