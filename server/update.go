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

func (s *Server) UpdateBlog(ctx context.Context, in *pb.UpdateBlogRequest) (*emptypb.Empty, error) {
	log.Printf("ReadBlog was invoked with %v\n", in)

	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Cannot Parse ID")
	}

	data := &BlogItem{AuthorID: in.AuthorId, Title: in.Title, Content: in.Content}

	res, err := s.collection.UpdateOne(
		ctx,
		bson.M{"_id": oid},
		bson.M{"$set": data},
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error: %v", err)
	}

	if res.MatchedCount == 0 {
		return nil, status.Errorf(codes.NotFound, "Cannot find blog with Id")
	}
	return &emptypb.Empty{}, nil
}
