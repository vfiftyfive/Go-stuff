package main

import (
	"context"
	"github.com/vfiftyfive/cisco/goucs"
	"log"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	url := "10.51.48.11"
	user := "admin"
	pwd := "C1sco123"

	u, err := goucs.ParseURL(url)
	if err != nil {
		log.Fatal(err)
	}
	cl, err := goucs.NewClient(ctx, u, true, user, pwd)
	if err != nil {
		log.Fatal(err)
	}
	defer cl.Logout(ctx)
}
