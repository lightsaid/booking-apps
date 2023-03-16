package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"time"

	"github.com/lightsaid/booking-gapi/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

const (
	address = "localhost:5800"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	fatalOnerror(err, "grpc.Dial")
	defer conn.Close()

	// smsClient := pb.NewSMSServiceClient(conn)
	// authClient := pb.NewAuthServiceClient(conn)
	// registerUser(authClient, smsClient)
	// loginUser(authClient, smsClient)

	// ===========
	// userClient := pb.NewUserServiceClient(conn)

	// =========
	movieClient := pb.NewMovieServiceClient(conn)
	listMovie(movieClient)
	getMovie(movieClient)

}

func registerUser(client pb.AuthServiceClient, smsClient pb.SMSServiceClient) {
	log.Println(">>>>>>>> TEST REGISTER >>>>>>>")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	phone := "18765432104"
	// sms := mocksms.NewMockSMS(phone)
	code, err := smsClient.Send(ctx, wrapperspb.String(phone))
	if err != nil {
		log.Println("registerUser smsClient.Send: ", err)
		return
	}
	resp, err := client.Register(ctx, &pb.RegisterRequest{PhoneNumber: phone, Code: code.Value})
	if err != nil {
		log.Println("client.Register failed ", err.Error())
		return
	}
	log.Println("注册成功：", resp.User.Id, resp.User.Name, resp.User.PhoneNumber)
}

func loginUser(client pb.AuthServiceClient, smsClient pb.SMSServiceClient) {
	log.Println(">>>>>>>> TEST LOGIN >>>>>>>")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	phone := "18765432103"
	// sms := mocksms.NewMockSMS(phone)

	code, err := smsClient.Send(ctx, wrapperspb.String(phone))
	if err != nil {
		log.Println("loginUser smsClient.Send: ", err)
		return
	}

	rsp, err := client.Login(ctx, &pb.LoginRequest{PhoneNumber: phone, Code: code.Value})
	if err != nil {
		log.Println("client.Register failed ", err.Error())
		return
	}
	buf, _ := json.Marshal(rsp)
	log.Println("登录成功：", string(buf))
}

func fatalOnerror(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err.Error())
	}
}

func listMovie(client pb.MovieServiceClient) {
	log.Println(">>>>>>>> TEST ListMovie >>>>>>>")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	stream, err := client.ListMovie(ctx, &pb.ListMovieRequest{Pager: &pb.Pagation{Page: 1, Size: 10}})
	if err != nil {
		log.Println("client.ListMovie failed: ", err)
		return
	}

	for {
		movie, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("stream.Recv() error: ", err)
			return
		}
		log.Println("获取到Title： ", movie.Movie.Title)
	}
}

func getMovie(client pb.MovieServiceClient) {
	log.Println(">>>>>>>> TEST getMovie >>>>>>>")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	resp, err := client.GetMovie(ctx, &pb.GetMovieRequest{Id: 2})
	if err != nil {
		log.Println("GetMovie() error: ", err)
		return
	}

	buf, _ := json.MarshalIndent(resp, "", " ")
	log.Println(string(buf))
}
