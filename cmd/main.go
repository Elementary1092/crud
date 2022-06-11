package main

import (
	"fmt"
	"github.com/Elementary1092/crud/internal/adapters/api"
	"github.com/Elementary1092/crud/internal/adapters/db"
	"github.com/Elementary1092/crud/internal/config"
	services "github.com/Elementary1092/crud/pkg/client/grpc"
	"github.com/Elementary1092/crud/pkg/client/postgre"
	"github.com/Elementary1092/crud/pkg/logging"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"time"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("Starting the CRUD service")

	cfg := config.GetConfiguration()

	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancelFunc()

	dbClient, err := postgre.NewClient(ctx, cfg.DBCfg)
	if err != nil {
		logger.Fatalf("Unable to start server due to: %v", err)
	}
	logger.Info("Successfully connected to database")

	storage := db.NewPostgreSQLStorage(dbClient, logger)

	service := api.NewService(storage)

	logger.Infof("Initializing server on port: %s", cfg.Listen.Port)
	server := api.NewServer(service, logger)
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.Address, cfg.Listen.Port))
	if err != nil {
		logger.Fatalf("Unable to start server due to: %v", err)
	}

	grpcServer := grpc.NewServer()
	services.RegisterCRUDServiceServer(grpcServer, server)

	reflection.Register(grpcServer)

	logger.Fatalf("Server is down due to: %v", grpcServer.Serve(listener))
}
