package main

import (
	"context"
	"fmt"
	"github.com/vfiftyfive/cisco/goucs"
	"github.com/vfiftyfive/cisco/goucs/mo"
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

	fv := mo.FabricVlan{
		Id:      "179",
		Name:    "api_test",
		Sharing: "none",
		Status:  "created",
	}

	Dn := "fabric/lan/A"

	resp, err := cl.ConfigConfMo(ctx, Dn, fv)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("HTTP Response Body:\n %s\n", string(resp))

	defer cl.Logout(ctx)
}
