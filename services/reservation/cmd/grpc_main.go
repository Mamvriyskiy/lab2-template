package main

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"reservation/internal/config"
	"reservation/internal/delivery"
	"reservation/internal/repo"
	"reservation/internal/usecase"
	"reservation/proto/reservation"
)

const (
	grpcPort = "50052"
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
	usecase := usecase.NewUsecase(repo)

	grpcServer := grpc.NewServer()
	reservationService := delivery.NewGRPCHandler(usecase)
	reservation.RegisterReservationServiceServer(grpcServer, reservationService)

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
