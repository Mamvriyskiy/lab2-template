package main

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"library/internal/config"
	"library/internal/delivery"
	"library/internal/repo"
	"library/internal/usecase"
	"library/proto/library"
	"log"
	"net"
	"os"
)

const (
	grpcPort = "50051"
)

func main() {

	if err := runGRPC(); err != nil {
		log.Print(err)
		os.Exit(1)
	}
}

func runGRPC() error {
	conn, err := config.GetConnectionString()
	if err != nil {
		return err
	}

	pool, err := pgxpool.Connect(context.Background(), conn)
	if err != nil {
		return err
	}
	repo := repo.NewLibRepo(pool)
	usecase := usecase.NewLibUsecase(repo)

	grpcServer := grpc.NewServer()
	libraryService := delivery.NewGRPCHandler(usecase)
	library.RegisterLibraryServiceServer(grpcServer, libraryService)

	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return err
	}

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to start gRPC server: %v", err)
		return err
	}
	return nil
}
