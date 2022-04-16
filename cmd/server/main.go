package main

import (
	"log"

	"github.com/PatriciaChebet/rocket_service_gRPC/internal/db"
	"github.com/PatriciaChebet/rocket_service_gRPC/internal/rocket"
	"github.com/PatriciaChebet/rocket_service_gRPC/internal/transport/grpc"
)

func Run() error {
	//responsible for initialising and starting our gRPC server
	rocketStore, err := db.New()
	if err != nil {
		return err
	}

	err = rocketStore.Migrate()
	if err != nil {
		log.Println("Failed to run migrations")
		return err
	}

	rktService := rocket.New(rocketStore)
	rktHandler := grpc.New(rktService)

	if err := rktHandler.Serve(); err != nil {
		return err
	}
	return nil
}
func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
