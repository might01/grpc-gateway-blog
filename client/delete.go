package main

import (
	"context"
	"log"

	pb "bitbucket.com/mightnvi/grpc-blog/proto/blog/v1"
)

func deleteBlog(c pb.BlogServiceClient, id string) {
	log.Printf("deleteBlog was invoked with %s\n", id)

	_, err := c.DeleteBlog(context.Background(), &pb.DeleteBlogRequest{Id: id})
	if err != nil {
		log.Fatalf("error with %v\n", err)
	}

	log.Println("Blog was deleted")
}
