package main

import (
	"context"
	//"errors"
	"fmt"
	"github.com/vfiftyfive/vmware/utils"
	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/object"
	//"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	"log"
	//"reflect"
)

const vcURL = "https://administrator@vsphere.local:C!5co123@nvermand-vc-01.uktme.cisco.com/sdk"
const cluster = "pod-02"
const host = "nvermand-esxi-05.uktme.cisco.com"
const hostToRemove = "nvermand-esxi-03.ukmte.cisco.com"

func main() {

	//Parse vCenter URL and log error if problem
	url, err := soap.ParseURL(vcURL)
	if err != nil {
		log.Fatal(err)
	}
	//Create context with cancel option
	ctx, cancel := context.WithCancel(context.Background())
	//Terminate channel before returning
	defer cancel()
	//Create new session instance to vCenter
	c, err := govmomi.NewClient(ctx, url, true)
	if err != nil {
		log.Fatal(err)
	}
	//logout when before returning
	defer c.Logout(ctx)
	vl, err := utils.GetVMWithStatus(cluster, "disconnected", c, ctx)
	if (err != nil) {
		log.Fatal(err)
	}
	//Remove VMs from inventory
	//done := make(chan bool)
	for _, vm := range vl {
	//	go func(vm utils.VM) {
			_, err := vm.Folder.ObjectName(ctx)
			hostMo, err := utils.GetObjectFromName(host, []string{"HostSystem"}, c, ctx, []mo.HostSystem{})
			hostMoToRemove, err := utils.GetObjectFromName(hostToRemove, []string{"HostSystem"}, c, ctx, []mo.HostSystem{})
			h, ok := hostMoToRemove.(mo.HostSystem); 
			if !ok {
				log.Fatalf("%v is no type HostSystem", h)
			}
			h1, ok := hostMo.(mo.HostSystem); 
			if !ok {
				log.Fatalf("%v is no type HostSystem", h1)
			}
			hObj := object.NewHostSystem(c.Client, h.Reference())
			h1Obj := object.NewHostSystem(c.Client, h1.Reference())
			if err != nil {
				log.Fatal(err)
			}
			if err := utils.RemoveHost(ctx, *hObj); err == nil {
				if err := utils.BlockingRegisterVM(vm.Folder, vm.Path, *h1Obj, c, ctx); err != nil {
					fmt.Printf("Can't register VM %s: %s", vm.Name, err)
				}
			}
	//		done <- true
	//	}(v)
	}
}