package main

import (
	"database/sql"
	"flag"
	"log"
	"net"
	"study/grpc/internal/database"
	"study/grpc/internal/service"
	"study/grpc/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	address := flag.String("address", ":50051", "address the gRPC server listens on")
	flag.Parse()

	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	if err := database.Migrate(db); err != nil {
		log.Fatal(err)
	}

	categoryDB := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDB)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)

	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", *address)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("gRPC server listening on %s", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
