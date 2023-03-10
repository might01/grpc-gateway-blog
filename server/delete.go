package main

import (
	pb "bitbucket.com/mightnvi/grpc-blog/proto/blog/v1"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

func (s *Server) DeleteBlog(ctx context.Context, in *pb.DeleteBlogRequest) (*emptypb.Empty, error) {
	log.Printf("DeleteBlog was invoked with: %v\n", in)

	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Cannot Parse ID")
	}

	res, err := s.collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Cannot Delete: %v", err)
	}

	if res.DeletedCount == 0 {
		return nil, status.Errorf(codes.NotFound, "Blog was not found")
	}

	return &emptypb.Empty{}, nil
}
