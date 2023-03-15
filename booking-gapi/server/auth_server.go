package server

import (
	"context"

	db "github.com/lightsaid/booking-gapi/db/postgres"
	"github.com/lightsaid/booking-gapi/pb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type AuthServer struct {
	store db.Store
	pb.UnimplementedAuthServiceServer
}

func (srv *AuthServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {

	srv.store.CreateUser(ctx, db.CreateUserParams{})

	return nil, nil
}

func (srv *AuthServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	return nil, nil
}

func (srv *AuthServer) Refresh(ctx context.Context, req *wrapperspb.StringValue) (*wrapperspb.StringValue, error) {
	return nil, nil
}
