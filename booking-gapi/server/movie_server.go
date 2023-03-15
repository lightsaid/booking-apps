package server

import (
	"context"
	"database/sql"
	"log"

	db "github.com/lightsaid/booking-gapi/db/postgres"
	"github.com/lightsaid/booking-gapi/pb"
	"github.com/lightsaid/booking-gapi/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type MovieServer struct {
	store db.Store
	pb.UnimplementedAuthServiceServer
}

func toPbMovie(item *db.TbMovie) *pb.Movie {
	return &pb.Movie{
		Id:          item.ID,
		Title:       item.Title,
		ReleaseDate: timestamppb.New(item.ReleaseDate),
		Poster:      item.Poster,
		Director:    item.Director,
		Description: *item.Description,
		Genre:       *item.Genre,
		Star:        *item.Star,
		Duration:    item.Duration,
	}
}

func (srv *MovieServer) ListMovie(req *pb.ListMovieRequest, stream pb.MovieService_ListMovieServer) error {
	limit, offset := utils.GetOffset(req.Pager)
	list, err := srv.store.ListMovies(stream.Context(), db.ListMoviesParams{Limit: limit, Offset: offset})
	if err != nil {
		log.Println("ListMovie failed: ", err)
		return status.Errorf(codes.Internal, "查询失败")
	}

	for _, item := range list {
		err = stream.Send(&pb.ListMovieResponse{
			Movie: toPbMovie(item),
		})
		if err != nil {
			log.Println("ListMovie Send failed: ", err)
			return status.Errorf(codes.Unknown, "传输数据错误")
		}
	}
	return nil
}

func (srv *MovieServer) GetMovie(ctx context.Context, req *pb.GetMovieRequest) (*pb.GetMovieResponse, error) {
	if req.Id <= 0 {
		return nil, status.Errorf(codes.NotFound, "id 不存在")
	}
	movie, err := srv.store.GetMovie(ctx, req.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "数据不存在")
		}
		return nil, status.Errorf(codes.Internal, "查询失败")
	}

	return &pb.GetMovieResponse{Movie: toPbMovie(movie)}, nil
}
