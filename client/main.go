package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "bitbucket.com/mightnvi/grpc-blog/proto/blog/v1"
)

var addr = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Did not connect: %v\n", err)
	}

	defer conn.Close()
	c := pb.NewBlogServiceClient(conn)

	id := createBlog(c)
	log.Println(readBlog(c, id))

	updateBlog(c, id)
	log.Println(readBlog(c, id))

	listBlog(c)

	deleteBlog(c, id)
	log.Println(readBlog(c, id))
}
