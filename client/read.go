package main

import (
	pb "bitbucket.com/mightnvi/grpc-blog/proto"
	"context"
	"log"
)

func readBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	log.Printf("readBlog was invoked with %s\n", id)

	req := &pb.BlogId{Id: id}

	res, err := c.ReadBlog(context.Background(), req)

	if err != nil {
		log.Printf("Error while reading: %v\n", err)
	}

	return res
}
