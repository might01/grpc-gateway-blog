package main

import (
	"context"
	"log"

	pb "bitbucket.com/mightnvi/grpc-blog/proto/blog/v1"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("createBlog was invoked")

	blog := &pb.CreateBlogRequest{AuthorId: "test", Title: "test_title", Content: "test_content"}
	res, err := c.CreateBlog(context.Background(), blog)
	if err != nil {
		log.Fatalf("error with %v\n", err)
	}

	log.Printf("Blog created with id %s\n", res.Id)
	return res.Id
}
