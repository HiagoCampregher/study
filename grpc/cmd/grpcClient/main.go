package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"study/grpc/pb"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	address := flag.String("address", "localhost:50051", "gRPC server address")
	name := flag.String("name", "gRPC", "category name")
	description := flag.String("description", "Remote procedure calls", "category description")
	flag.Parse()

	conn, err := grpc.NewClient(
		*address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client := pb.NewCategoryServiceClient(conn)
	category, err := client.CreateCategory(ctx, &pb.CategoryRequest{
		Name:        *name,
		Description: *description,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("created: id=%s name=%q description=%q\n", category.Id, category.Name, category.Description)
}
