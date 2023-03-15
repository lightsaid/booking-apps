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

func (srv *UserServer) GetProfile(ctx context.Context, tmp *emptypb.Empty) (*pb.GetProfileResponse, error) {
	// TODO：
	rsp := &pb.GetProfileResponse{User: &pb.User{Id: 1, Name: "张三", PhoneNumber: "18765432101"}}
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
		Id:          user.RoleID,
		RoleId:      user.RoleID,
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
		Avatar:      *user.Avatar,
	}}

	return &rsp, nil
}
