package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"toolkit/jwtutil"

	_ "github.com/lib/pq"
	db "github.com/lightsaid/booking-gapi/db/postgres"
	"github.com/lightsaid/booking-gapi/pb"
	"github.com/lightsaid/booking-gapi/server"
	"google.golang.org/grpc"
)

const (
	jwtSecret = "abcdefg123456abcdefg123456abcdefg123456"
	address   = "0.0.0.0:5800"
)

func main() {
	var db_source = "postgresql://postgres:postgres_booking@localhost:5555/db_bookings?sslmode=disable"
	conn, err := sql.Open("postgres", db_source)
	fatalOnerror(err, "sql.Open")
	defer conn.Close()

	store := db.NewStore(conn)
	jwtMaker, err := jwtutil.NewMaker(jwtSecret)
	fatalOnerror(err, "jwtutil.NewMaker")

	authServer := server.NewAuthServer(store, jwtMaker)

	userServer := server.NewUserServer(store)

	movieServer := server.NewMovieServer(store)

	smsServer := server.NewSMSServer()

	lis, err := net.Listen("tcp", address)
	fatalOnerror(err, "net.Listen")
	defer lis.Close()

	srv := grpc.NewServer()
	pb.RegisterAuthServiceServer(srv, authServer)
	pb.RegisterUserServiceServer(srv, userServer)
	pb.RegisterMovieServiceServer(srv, movieServer)
	pb.RegisterSMSServiceServer(srv, smsServer)

	fmt.Println("Start server on ", address)
	err = srv.Serve(lis)
	fatalOnerror(err, "srv.Serve")

}

func fatalOnerror(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err.Error())
	}
}
