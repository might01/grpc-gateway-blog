package main

import (
	pb "bitbucket.com/mightnvi/grpc-blog/proto"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
)

var addr = "0.0.0.0:50051"

type Server struct {
	collection *mongo.Collection

	pb.BlogServiceServer
}

func main() {
	db := getEnv("MONGO_URI", "localhost:27017")
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@" + db + "/"))
	if err != nil {
		log.Fatalf("Error while try to initial db con: %v\n", err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("blogdb").Collection("blog")

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}
	defer lis.Close()

	server := &Server{collection: collection}

	log.Printf("Listening at %s\n", addr)
	s := grpc.NewServer()
	pb.RegisterBlogServiceServer(s, server)

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
