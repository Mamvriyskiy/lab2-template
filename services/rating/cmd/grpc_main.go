package main

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"rating/internal/config"
	"rating/internal/delivery"
	"rating/internal/repo"
	"rating/internal/usecase"
	"rating/proto/rating"
)

const (
	grpcPort = "50053"
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
	repo := repo.NewRepo(pool)
	usecase := usecase.NewRatingUsecase(repo)

	grpcServer := grpc.NewServer()
	ratingService := delivery.NewGRPCHandler(usecase)
	rating.RegisterRatingServiceServer(grpcServer, ratingService)

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
