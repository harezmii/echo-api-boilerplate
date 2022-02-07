package ent

import (
	"context"
	"fmt"
	_ "github.com/lib/pq"
)

func ConnectionEnt() *Client {
	client, _ := Open("postgres", "host=192.168.8.162 port=5432 user=postgres dbname=development password=123456 sslmode=disable")
	if err := client.Schema.Create(context.Background()); err != nil {
		fmt.Println(err)
	}
	return client
}
