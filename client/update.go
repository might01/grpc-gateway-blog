package main

import (
	pb "bitbucket.com/mightnvi/grpc-blog/proto/blog/v1"
	"context"
	"log"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Printf("updateBlog was invoked with %s\n", id)

	newBlog := &pb.UpdateBlogRequest{
		Id:       id,
		AuthorId: "test_update",
		Title:    "test_title_update",
		Content:  "test_content_update",
	}
	_, err := c.UpdateBlog(context.Background(), newBlog)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Blog was updated")
}
