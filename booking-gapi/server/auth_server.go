package server

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"
	"toolkit/jwtutil"
	"toolkit/random"

	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v4"
	db "github.com/lightsaid/booking-gapi/db/postgres"
	"github.com/lightsaid/booking-gapi/pb"
	"github.com/lightsaid/booking-gapi/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type AuthServer struct {
	store db.Store
	pb.UnimplementedAuthServiceServer
	jwtMaker *jwtutil.Maker
	smsStore SMSStore
}

func NewAuthServer(store db.Store, jwtMaker *jwtutil.Maker) *AuthServer {
	return &AuthServer{
		store:    store,
		jwtMaker: jwtMaker,
		smsStore: NewSMSStore(),
	}
}

func createJWTPayload(uid int64, expires time.Time) (*jwtutil.JWTPayload, error) {
	jid, err := uuid.NewV4()
	if err != nil {
		return nil, errors.New("生成 JWT ID 失败")
	}
	payload := jwtutil.NewJWTPayload(uid, jwt.RegisteredClaims{
		ID:        jid.String(),
		ExpiresAt: jwt.NewNumericDate(expires),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Issuer:    "booking-gapi",
	})
	return payload, nil
}

func (srv *AuthServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	v := utils.NewValidator()
	v.CheckPhone(req.PhoneNumber, "phone_number", "手机号码不正确")
	v.Check(req.Code >= 1000, "code", "验证码至少是4为数字")

	if !v.Valid() {
		return nil, status.Errorf(codes.InvalidArgument, "参数不对: %s", v.String())
	}

	// sms, ok := mocksms.GetMockSMS(req.PhoneNumber)
	// if !ok || sms != nil && sms.Code() != int64(req.Code) {
	// 	return nil, status.Errorf(codes.InvalidArgument, "验证码不正确")
	// }

	if ok := srv.smsStore.Valid(req.PhoneNumber, req.Code); !ok {
		return nil, status.Errorf(codes.InvalidArgument, "验证码不正确")
	}

	user, err := srv.store.GetUserByPhone(ctx, req.PhoneNumber)
	if err != nil && err != sql.ErrNoRows {
		log.Println("Register GetUserByPhone failed1: ", err)
		return nil, status.Error(codes.Internal, "查询用户失败")
	}
	if user.ID > 0 {
		log.Println("Register GetUserByPhone failed2: ", user.ID, user.RoleID, user.Avatar)
		return nil, status.Error(codes.AlreadyExists, "用户已存在")
	}

	newUser, err := srv.store.CreateUser(ctx, db.CreateUserParams{PhoneNumber: req.PhoneNumber, Name: random.RandomString(5)})
	if err != nil {
		log.Println("Register CreateUser failed1: ", err)
		return nil, status.Error(codes.Internal, "注册失败")
	}
	var avatar string
	if user.Avatar != nil {
		avatar = *user.Avatar
	}
	var roleId int64
	if user.RoleID != nil {
		roleId = *user.RoleID
	}
	rsp := pb.RegisterResponse{User: &pb.User{
		Id:          newUser.ID,
		Name:        newUser.Name,
		RoleId:      roleId,
		Avatar:      avatar,
		PhoneNumber: newUser.PhoneNumber,
	}}
	return &rsp, nil
}

func (srv *AuthServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	v := utils.NewValidator()
	v.CheckPhone(req.PhoneNumber, "phone_number", "手机号码不正确")
	v.Check(req.Code >= 1000, "code", "验证码至少是4为数字")

	if !v.Valid() {
		return nil, status.Errorf(codes.InvalidArgument, "参数不对: %s", v.String())
	}

	// sms, ok := mocksms.GetMockSMS(req.PhoneNumber)
	// if !ok || sms != nil && sms.Code() != int64(req.Code) {
	// 	return nil, status.Errorf(codes.InvalidArgument, "验证码不正确")
	// }

	if ok := srv.smsStore.Valid(req.PhoneNumber, req.Code); !ok {
		return nil, status.Errorf(codes.InvalidArgument, "验证码不正确")
	}

	user, err := srv.store.GetUserByPhone(ctx, req.PhoneNumber)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "用户不存在")
		}
		return nil, status.Error(codes.Internal, "查询用户失败")
	}

	aPayload, err := createJWTPayload(user.ID, time.Now().Add(30*time.Minute))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	access_token, err := srv.jwtMaker.GenToken(aPayload)
	if err != nil {
		return nil, status.Error(codes.Internal, "生成Token失败")
	}

	rPayload, err := createJWTPayload(user.ID, time.Now().Add(30*time.Minute))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	refresh_token, err := srv.jwtMaker.GenToken(rPayload)
	if err != nil {
		return nil, status.Error(codes.Internal, "生成Token失败")
	}
	var avatar string
	if user.Avatar != nil {
		avatar = *user.Avatar
	}
	var roleId int64
	if user.RoleID != nil {
		roleId = *user.RoleID
	}
	rsp := &pb.LoginResponse{
		AccessToken:  access_token,
		RefreshToken: refresh_token,
		User:         &pb.User{Id: user.ID, Name: user.Name, Avatar: avatar, RoleId: roleId, PhoneNumber: user.PhoneNumber},
	}

	return rsp, nil
}

func (srv *AuthServer) Refresh(ctx context.Context, req *wrapperspb.StringValue) (*wrapperspb.StringValue, error) {
	// TODO 拦截器，从ctx查找userid
	var uid int64 = 1

	payload, err := createJWTPayload(uid, time.Now().Add(30*time.Minute))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	token, err := srv.jwtMaker.GenToken(payload)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return wrapperspb.String(token), nil
}
