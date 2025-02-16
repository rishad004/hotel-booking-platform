package main

import (
	"log"
	"net"

	"github.com/rishad004/hotel-booking-platform/user-service/internal/handler"
	"github.com/rishad004/hotel-booking-platform/user-service/internal/repository"
	"github.com/rishad004/hotel-booking-platform/user-service/internal/service"
	pb "github.com/rishad004/hotel-booking-platform/user-service/proto"
	hotel_pb "github.com/rishad004/hotel-booking-platform/user-service/proto/client/hotel"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	connHotel, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to user service: %v", err)
	}
	defer connHotel.Close()
	hotelService := hotel_pb.NewHotelServiceClient(connHotel)

	repo := repository.NewUserRepo()
	userService := service.NewUserService(repo, hotelService)
	userHandler := handler.NewUserHandler(userService)

	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("failed to listen, err: ", err)
	}

	g := grpc.NewServer()
	pb.RegisterUserServiceServer(g, userHandler)

	reflection.Register(g)

	log.Println("Server listening on port :8080")
	if err := g.Serve(listen); err != nil {
		log.Fatal("failed to listen - err: ", err)
	}

}
