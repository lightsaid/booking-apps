package server

import (
	"context"
	"log"
	"toolkit/random"

	"github.com/lightsaid/booking-gapi/pb"
	"github.com/lightsaid/booking-gapi/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type SMSServer struct {
	pb.UnimplementedSMSServiceServer
	smsStore SMSStore
}

func NewSMSServer() *SMSServer {
	store := NewSMSStore()
	return &SMSServer{
		smsStore: store,
	}
}

func (srv *SMSServer) Send(ctx context.Context, req *wrapperspb.StringValue) (*wrapperspb.Int32Value, error) {
	v := utils.NewValidator()
	v.CheckPhone(req.Value, "phone", "手机号码不正确")
	if !v.Valid() {
		return nil, status.Errorf(codes.InvalidArgument, v.String())
	}

	code := random.RandomInt(1000, 9999)
	err := srv.smsStore.Save(req.Value, int32(code))
	if err != nil {
		log.Println("rpc send failed ", err)
		return nil, status.Error(codes.Internal, "发送短信验证码失败")
	}

	return wrapperspb.Int32(int32(code)), nil
}
