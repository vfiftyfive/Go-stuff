package main

import (
	"testing"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vfiftyfive/vmware/utils"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25/soap"
	"context"
	"github.com/vmware/govmomi"
	

)

const vcURL = "https://administrator@vsphere.local:C!5co123@nvermand-vc-01.uktme.cisco.com/sdk"
const cluster = "pod-02"

type Helper struct {
	ctx context.Context  
	c *govmomi.Client
	err error
	cancel context.CancelFunc 
}

func login (vcURL string) (h Helper) {
	//Parse vCenter URL and log error if problem
	url, err := soap.ParseURL(vcURL)
	if err != nil {
		h.err = err
		return h
	}
	//Create context with cancel option
	h.ctx, h.cancel = context.WithCancel(context.Background())
	//Create new session instance to vCenter
	h.c, h.err = govmomi.NewClient(h.ctx, url, true)
	if err != nil {
		h.cancel()
		return h
	}
	return h
}

//Test object View Container
func TestGetVMEntityView (t *testing.T) {
	h := login (vcURL)
	if (h.err != nil){
		t.Fatal(h.err)
	} else {
		t.Log("Login to vCenter succeedeed")
	}
	//Define new VM container view and destroy view when returning
	m := view.NewManager(h.c.Client)
 	v, err := m.CreateContainerView(h.ctx, h.c.ServiceContent.RootFolder, []string{"VirtualMachine"}, true)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log("Containerview created")
	}
	defer v.Destroy(h.ctx)
	var vms []mo.VirtualMachine
	err = v.Retrieve(h.ctx, []string{"VirtualMachine"}, nil, &vms)
	if err != nil {
		t.Log(h.err)
	} else {
		for _, v:= range vms {
		t.Logf("%v", v.Name)
		}
	}
}

//Test find disconnected VM in cluster
func TestFindDisconnectedVM(t *testing.T) ([]mo.VirtualMachine) {
	h := login (vcURL)
	if (h.err != nil){
		t.Fatal(h.err)
	} else {
		t.Log("Login to vCenter succeedeed")
	}
	vl, err := utils.GetVMWithStatus(cluster, "disconnected", h.c, h.ctx)
	if (err != nil) {
		t.Fatal(err)
	} else {
		for _, v := range vl {
		t.Logf("%v", v.Name)	
		}
		return vl
	}

}
