package main

import (
	"log"

	"github.com/PatriciaChebet/rocket_service_gRPC/internal/db"
	"github.com/PatriciaChebet/rocket_service_gRPC/internal/rocket"
)

func Run() error {
	//responsible for initialising and starting our gRPC server
	rocketStore, err := db.New()
	if err != nil {
		return err
	}
	_ = rocket.New(rocketStore)
	return nil
}
func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
