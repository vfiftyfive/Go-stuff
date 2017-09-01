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

	//ConfigConfMo Test
	// fv := mo.FabricVlan{
	// 	Id:      "179",
	// 	Name:    "api_testsdfds",
	// 	Sharing: "none",
	// 	Status:  "created",
	// }

	// Dn := "fabric/lan/A"

	// _, err = cl.ConfigConfMo(ctx, Dn, fv)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	//ConfigConfMos Test - private VLANs
	fv1 := mo.FabricVlan{
		Id:      "279",
		Name:    "test-279",
		Sharing: "primary",
		Status:  "created",
	}
	fv2 := mo.FabricVlan{
		Id:        "379",
		Name:      "test-379",
		Sharing:   "isolated",
		PubNwName: "test-279",
		Status:    "created",
	}
	p1 := mo.Pair{
		Key: "fabric/lan/A/net-" + fv1.Name,
		Mo:  fv1,
	}
	p2 := mo.Pair{
		Key: "fabric/lan/A/net-" + fv2.Name,
		Mo:  fv2,
	}

	p := []mo.Pair{p1, p2}
	_, err = cl.ConfigConfMos(ctx, p)

	if err != nil {
		log.Fatalf("Error when creating object: %s", err)
	}

	defer cl.Logout(ctx)
}
