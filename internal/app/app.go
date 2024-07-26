package app

import (
	grpcapp "example.com/pet_auth/internal/app/grpc"
	"log/slog"
	"time"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(log *slog.Logger, grpcPort int, storagePath string, tokenTTL time.Duration) *App {
	// TODO ИНИЦИАЛИЩИРОВАТЬ ХРАНИЛИЩЕ
	// TODO ИНИЦИАЛИЩИРОВАТЬ AUTHSERVICE

	grpcApp := grpcapp.New(log, grpcPort)

	return &App{
		GRPCSrv: grpcApp,
	}
}
