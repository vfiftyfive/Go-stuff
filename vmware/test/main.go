package main

import (
	"context"
	//"errors"
	"fmt"
	"github.com/vfiftyfive/vmware/utils"
	"github.com/vmware/govmomi"
	//"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	"log"
	//"reflect"
)

const vcURL = "https://root:C!5co123C!5co123@nvermand-vc-01.uktme.cisco.com/sdk"
const clusterName = "pod-02"
const path = "[nvermand_esxi_nfs_datastore] DLR-01-0/DLR-01-0.vmx"

var hostName = "nvermand-esxi-03.uktme.cisco.com"
var hostMos []mo.HostSystem

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
	//Define new VM container view and destroy view when returning
	m := view.NewManager(c.Client)
	v, err := m.CreateContainerView(ctx, c.ServiceContent.RootFolder, []string{"VirtualMachine"}, true)
	if err != nil {
		log.Fatal(err)
	}
	defer v.Destroy(ctx)
	var vms []mo.VirtualMachine
	err = v.Retrieve(ctx, []string{"VirtualMachine"}, nil, &vms)
	if err != nil {
		log.Fatal(err)
	}

	if err := utils.BlockingRegisterVM("ACI", path, hostName, c, ctx); err != nil {
		fmt.Println(err)
	}

	//Test folders and VM register
	// var rootFolder = object.NewRootFolder(c.Client)
	// datacenterList, err := rootFolder.Children(ctx)
	// if datacenter, ok := datacenterList[0].(*object.Datacenter); ok {
	// 	dcFolder, err := datacenter.Folders(ctx)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	i, err := utils.GetObjectFromName(hostName, []string{"HostSystem"}, c, ctx, hostMos)
	// 	if hostMo, ok := i.(mo.HostSystem); ok {
	// 		hostFolderChildren, err := dcFolder.HostFolder.Children(ctx)
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}
	// 		var resources *object.ResourcePool
	// 		found := false
	// 		for _, c := range hostFolderChildren {
	// 			if cl, ok := c.(*object.ClusterComputeResource); ok {
	// 				clName, err := cl.ObjectName(ctx)
	// 				if err != nil {
	// 					log.Fatal(err)
	// 				}
	// 				if clName == clusterName {
	// 					found = true
	// 					resources, err = cl.ResourcePool(ctx)
	// 					if err != nil {
	// 						log.Fatal(err)
	// 					}
	// 					break
	// 				}
	// 			}
	// 		}
	// 		if found == false {
	// 			log.Fatal(errors.New("no cluster found"))
	// 		}
	// 		pool := object.NewResourcePool(c.Client, resources.Reference())
	// 		task, err := dcFolder.VmFolder.RegisterVM(ctx, path, "", false, pool, object.NewHostSystem(c.Client, hostMo.Reference()))
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}
	// 		if err := task.Wait(ctx); err != nil {
	// 			log.Fatal(err)
	// 		}
	// 	}
	// }
}
