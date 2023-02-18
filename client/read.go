package main

import (
	pb "bitbucket.com/mightnvi/grpc-blog/proto/blog/v1"
	"context"
	"log"
)

func readBlog(c pb.BlogServiceClient, id string) *pb.ReadBlogResponse {
	log.Printf("readBlog was invoked with %s\n", id)

	req := &pb.ReadBlogRequest{Id: id}

	res, err := c.ReadBlog(context.Background(), req)

	if err != nil {
		log.Printf("Error while reading: %v\n", err)
	}

	return res
}
