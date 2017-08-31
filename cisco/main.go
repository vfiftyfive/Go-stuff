package main

import (
	"context"
	"github.com/vfiftyfive/cisco/goucs"
	"github.com/vfiftyfive/cisco/goucs/mo"
	"log"
)

var (
	url  = "10.51.48.11"
	user = "admin"
	pwd  = "C1sco123"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	u, err := goucs.ParseURL(url)
	if err != nil {
		log.Fatal(err)
	}
	cl, err := goucs.NewClient(ctx, u, true, user, pwd)
	if err != nil {
		log.Fatal(err)
	}

	fv := mo.FabricVlan{
		Id:      "179",
		Name:    "api_testsdfds",
		Sharing: "none",
		Status:  "created",
	}

	Dn := "fabric/lan/A"

	_, err = cl.ConfigConfMo(ctx, Dn, fv)
	if err != nil {
		log.Fatal(err)
	}

	defer cl.Logout(ctx)
}
