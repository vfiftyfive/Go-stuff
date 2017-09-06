package main

import (
	"context"
	//	"fmt"
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
		Status:  "created,modified",
	}
	fv2 := mo.FabricVlan{
		Id:        "379",
		Name:      "test-379",
		Sharing:   "isolated",
		PubNwName: "test-279",
		Status:    "created,modified",
	}

	p1 := mo.Pair{
		Key: "fabric/lan/A/net-" + fv1.Name,
		Mo:  fv1,
	}
	p2 := mo.Pair{
		Key: "fabric/lan/A/net-" + fv2.Name,
		Mo:  fv2,
	}
	p3 := mo.Pair{
		Key: "fabric/lan/B/net-" + fv1.Name,
		Mo:  fv1,
	}
	p4 := mo.Pair{
		Key: "fabric/lan/B/net-" + fv2.Name,
		Mo:  fv2,
	}

	p := []mo.Pair{p1, p2, p3, p4}
	_, err = cl.ConfigConfMos(ctx, p)
	// for _, m := range *cms.OutConfigs {
	// 	fmt.Println(m)
	// }

	if err != nil {
		log.Fatalf("Error when creating object: %s", err)
	}

	//ConfigResolveChildren Test
	crc, err := cl.ConfigResolveChildren(ctx, "vnicEther", "org-root/ls-sp-01", "false")
	if err != nil {
		log.Fatal(err)
	}

	//Create vnicEtherIf with corresponding VLANs
	vif1 := mo.VnicEtherIf{Name: fv1.Name}
	vif2 := mo.VnicEtherIf{Name: fv2.Name}
	v := []mo.VnicEtherIf{vif1, vif2}

	//Add vif to vnic
	vp := make([]mo.Pair, len(*crc.OutConfigs)*len(v))
	j := 0
	for j < len(*crc.OutConfigs)*len(v) {
		for i, m := range *crc.OutConfigs {
			//Create the corresponding pair
			//fmt.Println(reflect.TypeOf(m).Elem())
			if vnic, ok := m.(*mo.VnicEther); ok {
				t := 0
				for t < len(v) {
					vp[j] = mo.Pair{
						Key: vnic.Dn + "/if-" + v[i].Name,
						Mo:  v[t],
					}
					t++
					j++
				}
			}
		}
	}

	_, err = cl.ConfigConfMos(ctx, vp)
	if err != nil {
		log.Fatal(err)
	}

	defer cl.Logout(ctx)
}
