package server

import (
	"context"

	db "github.com/lightsaid/booking-gapi/db/postgres"
	"github.com/lightsaid/booking-gapi/pb"
)

type MovieServer struct {
	store db.Store
	pb.UnimplementedAuthServiceServer
}

func (srv *MovieServer) ListMovie(rq *pb.ListMovieRequest, stream pb.MovieService_ListMovieServer) error {
	return nil
}

func (srv *MovieServer) GetMovie(context.Context, *pb.GetMovieRequest) (*pb.GetMovieResponse, error) {
	return nil, nil
}
