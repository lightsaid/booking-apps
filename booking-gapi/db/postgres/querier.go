// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"context"
)

type Querier interface {
	CreateUser(ctx context.Context, arg CreateUserParams) (*TbUser, error)
	DeleteMovie(ctx context.Context, id int64) (*TbMovie, error)
	GetMovie(ctx context.Context, id int64) (*TbMovie, error)
	GetUser(ctx context.Context, id int64) (*TbUser, error)
	GetUserByPhone(ctx context.Context, phoneNumber string) (*TbUser, error)
	ListMovies(ctx context.Context, arg ListMoviesParams) ([]*TbMovie, error)
	UpdateMovie(ctx context.Context, arg UpdateMovieParams) (*TbMovie, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (*TbUser, error)
}

var _ Querier = (*Queries)(nil)
