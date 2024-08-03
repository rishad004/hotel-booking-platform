package main

import (
	"log"
	"net"

	"github.com/rishad004/hotel-booking-platform/hotel-service/internal/handler"
	"github.com/rishad004/hotel-booking-platform/hotel-service/internal/repository"
	"github.com/rishad004/hotel-booking-platform/hotel-service/internal/service"
	pb "github.com/rishad004/hotel-booking-platform/hotel-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	repo := repository.NewHotelRepo()
	hotelService := service.NewHotelService(repo)
	hotelHandler := handler.NewHotelHandler(hotelService)

	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal("failed to listen, err: ", err)
	}

	g := grpc.NewServer()
	pb.RegisterHotelServiceServer(g, hotelHandler)

	reflection.Register(g)

	log.Println("Server listening on port :8081")
	if err := g.Serve(listen); err != nil {
		log.Fatal("failed to listen - err: ", err)
	}

}
