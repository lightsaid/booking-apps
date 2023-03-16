package server

import (
	"context"
	"database/sql"
	"log"
	"time"

	db "github.com/lightsaid/booking-gapi/db/postgres"
	"github.com/lightsaid/booking-gapi/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserServer struct {
	store db.Store
	pb.UnimplementedUserServiceServer
}

func NewUserServer(store db.Store) *UserServer {
	return &UserServer{store: store}
}

func (srv *UserServer) GetProfile(ctx context.Context, tmp *emptypb.Empty) (*pb.GetProfileResponse, error) {
	// TODO： 从ctx获取id
	var uid int64 = 1
	user, err := srv.store.GetUser(ctx, uid)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "用户不存在")
		}
		return nil, status.Errorf(codes.Internal, "查询用户失败")
	}

	rsp := &pb.GetProfileResponse{User: &pb.User{Id: user.ID, Name: user.Name, PhoneNumber: user.PhoneNumber, Avatar: *user.Avatar}}
	return rsp, nil
}

func (srv *UserServer) UpdateProfile(ctx context.Context, req *pb.UpdateProfileRequest) (*pb.UpdateProfileResponse, error) {
	// TODO: validator input
	user, err := srv.store.GetUser(ctx, req.Id)
	if err != nil {
		log.Println("GetUser failed ", err)
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "用户不存在")
		}
		return nil, status.Errorf(codes.Internal, "查找用户失败")
	}

	if req.Name != "" {
		user.Name = req.Name
	}

	if req.Avatar != "" {
		user.Avatar = &req.Avatar
	}

	user, err = srv.store.UpdateUser(ctx, db.UpdateUserParams{
		ID:        user.ID,
		Name:      user.Name,
		Avatar:    user.Avatar,
		Openid:    user.Openid,
		Unionid:   user.Unionid,
		UpdatedAt: time.Now(),
	})
	if err != nil {
		log.Println("UpdateUser failed ", err)
		status.Errorf(codes.Internal, "更新失败")
	}
	rsp := pb.UpdateProfileResponse{User: &pb.User{
		Id:          user.ID,
		RoleId:      *user.RoleID,
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
		Avatar:      *user.Avatar,
	}}

	return &rsp, nil
}
