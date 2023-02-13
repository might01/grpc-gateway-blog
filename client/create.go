package main

import (
	pb "bitbucket.com/mightnvi/grpc-blog/proto"
	"context"
	"log"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("createBlog was invoked")

	blog := &pb.CreatBlogRequest{AuthorId: "test", Title: "test_title", Content: "test_content"}
	res, err := c.CreateBlog(context.Background(), blog)
	if err != nil {
		log.Fatalf("error with %v\n", err)
	}

	log.Printf("Blog created with id %s\n", res.Id)
	return res.Id
}
