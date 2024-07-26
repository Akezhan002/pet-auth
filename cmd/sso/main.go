package main

import (
	"example.com/pet_auth/internal/app"
	"example.com/pet_auth/internal/config"
	"example.com/pet_auth/internal/lib/logger/handlers/slogpretty"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info("starting application", slog.Any("config", cfg))

	applictaion := app.New(log, cfg.GRPC.Port, cfg.StoragePath, cfg.TokenTTL)

	go applictaion.GRPCSrv.MustRun()

	// TODO: иницилизировать APP

	// TODO: запустить grpc сервер

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop
	applictaion.GRPCSrv.Stop()
	log.Info("Application stop")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = setupPrettySlog()
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}
	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
