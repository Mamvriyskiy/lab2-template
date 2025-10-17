package main

import (
	"fmt"
	"gateway/internal/config"
	"gateway/internal/delivery"
	"gateway/internal/usecase"
	"gateway/middleware"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"
)

func main() {
	if err := run(); err != nil {
		log.Print(err)
		os.Exit(1)
	}
}
func run() error {
	r := mux.NewRouter()
	//os.Setenv("LIB_SERVICE_URL", "127.0.0.1:50051")
	//os.Setenv("RATING_SERVICE_URL", "127.0.0.1:50053")
	//os.Setenv("RESERVATION_SERVICE_URL", "127.0.0.1:50052")

	log.Printf("lib: %s", os.Getenv("LIB_SERVICE_URL"))
	log.Printf("rating: %s", os.Getenv("RATING_SERVICE_URL"))
	log.Printf("reservation: %s", os.Getenv("RESERVATION_SERVICE_URL"))
	srv := http.Server{Handler: r, Addr: fmt.Sprintf(":%s", "8080")}
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	clients, err := config.GetConnectionString()
	if err != nil {
		return err
	}
	libConn, err := grpc.Dial(clients.LibraryURL, opts...)
	if err != nil {
		return err
	}
	defer libConn.Close()
	reservationConn, err := grpc.Dial(clients.ReservationURL, opts...)
	if err != nil {
		return err
	}
	defer reservationConn.Close()
	ratingConn, err := grpc.Dial(clients.RatingURL, opts...)
	if err != nil {
		return err
	}
	defer ratingConn.Close()
	apiClient := delivery.NewGRPCClient(ratingConn, reservationConn, libConn)
	usecase := usecase.NewGatewayUsecase(apiClient)
	gatewayHandler := delivery.NewGatewayHandler(usecase)

	r.Use(middleware.CORSMiddleware)
	api := r.PathPrefix("/api/v1").Subrouter()
	{
		api.HandleFunc("/libraries", gatewayHandler.GetLibraries).Methods("GET")
		api.HandleFunc("/libraries/{libraryUid}/books", gatewayHandler.GetBooks).Methods("GET")
		api.HandleFunc("/reservations", gatewayHandler.GetReservations).Methods("GET")
		api.HandleFunc("/reservations", gatewayHandler.GetBook).Methods("POST")
		api.HandleFunc("/reservations/{reservationUid}/return", gatewayHandler.ReturnBook).Methods("POST")
		api.HandleFunc("/rating", gatewayHandler.GetRating).Methods("GET")
	}

	http.Handle("/", r)
	log.Print("main running on: ", srv.Addr)
	return srv.ListenAndServe()
}
