package main

import (
	"context"
	"github.com/vfiftyfive/cisco/goucs"
	//	"github.com/vfiftyfive/cisco/goucs/mo"
	"fmt"
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
	// fv1 := mo.FabricVlan{
	// 	Id:      "279",
	// 	Name:    "test-279",
	// 	Sharing: "primary",
	// 	Status:  "created,modified",
	// }
	// fv2 := mo.FabricVlan{
	// 	Id:        "379",
	// 	Name:      "test-379",
	// 	Sharing:   "isolated",
	// 	PubNwName: "test-279",
	// 	Status:    "created,modified",
	// }

	// p1 := mo.Pair{
	// 	Key: "fabric/lan/A/net-" + fv1.Name,
	// 	Mo:  fv1,
	// }
	// p2 := mo.Pair{
	// 	Key: "fabric/lan/A/net-" + fv2.Name,
	// 	Mo:  fv2,
	// }
	// p3 := mo.Pair{
	// 	Key: "fabric/lan/B/net-" + fv1.Name,
	// 	Mo:  fv1,
	// }
	// p4 := mo.Pair{
	// 	Key: "fabric/lan/B/net-" + fv2.Name,
	// 	Mo:  fv2,
	// }

	// p := []mo.Pair{p1, p2, p3, p4}
	// _, err = cl.ConfigConfMos(ctx, p)

	// if err != nil {
	// 	log.Fatalf("Error when creating object: %s", err)
	// }

	//ConfigResolveChildren Test
	crc, err := cl.ConfigResolveChildren(ctx, "vnicEther", "org-root/ls-sp-01")
	if err != nil {
		log.Fatal(err)
	}
	for _, m := range *crc.OutConfigs {
		fmt.Println(m)
	}

	defer cl.Logout(ctx)
}
