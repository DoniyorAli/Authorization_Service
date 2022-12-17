package main

import (
	"UacademyGo/Blogpost/Authorization_Service/config"
	"UacademyGo/Blogpost/Authorization_Service/protogen/blogpost"

	"UacademyGo/Blogpost/Authorization_Service/services/authorization"

	"UacademyGo/Blogpost/Authorization_Service/storage"
	"UacademyGo/Blogpost/Authorization_Service/storage/postgres"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// * @license.name  Apache 2.0
// * @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

func main() {

	cfg := config.Load()
	psqlConfString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)

	var err error
	var stg storage.StorageInter
	stg, err = postgres.InitDB(psqlConfString)
	if err != nil {
		panic(err)
	}

	println("gRPC server tutorial in Go")

	listener, err := net.Listen("tcp", cfg.GRPCPort)
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()

	authService := authorization.NewAuthService(cfg, stg)
	blogpost.RegisterAuthServiceServer(srv, authService)

	reflection.Register(srv)

	if err := srv.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
