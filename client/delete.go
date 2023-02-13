package main

import (
	pb "bitbucket.com/mightnvi/grpc-blog/proto"
	"context"
	"log"
)

func deleteBlog(c pb.BlogServiceClient, id string) {
	log.Printf("deleteBlog was invoked with %s\n", id)

	_, err := c.DeleteBlog(context.Background(), &pb.BlogId{Id: id})
	if err != nil {
		log.Fatalf("error with %v\n", err)
	}

	log.Println("Blog was deleted")
}
