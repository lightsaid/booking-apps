package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
	dbrepo "github.com/lightsaid/booking-sys/dbrepo/postgres"
)

func main() {
	var db_source = "postgresql://postgres:postgres_booking@localhost:5555/db_bookings?sslmode=disable"
	db, err := sql.Open("postgres", db_source)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	queries := dbrepo.New(db)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	arg := dbrepo.GetRolesParams{Limit: 10, Offset: 0}
	roles, err := queries.GetRoles(ctx, arg)
	if err != nil {
		log.Fatal(err)
	}
	js, err := json.MarshalIndent(roles, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(js))
}
