package utils

import (
	"context"
	"errors"
	"fmt"
	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25/mo"
	//"reflect"
)

//VM struct
type VM struct {
	Object     object.VirtualMachine
	Name       string
	Folder     object.Folder
	Host       object.HostSystem
	Path       string
	IsOrphaned bool
}

//GetObjectFromName returns interface from Name
func GetObjectFromName(name string, vimType []string, c *govmomi.Client, ctx context.Context, ati interface{}) (ato interface{}, err error) {

	m := view.NewManager(c.Client)
	v, err := m.CreateContainerView(ctx, c.ServiceContent.RootFolder, vimType, true)
	if err != nil {
		return
	}
	defer v.Destroy(ctx)
	switch objs := ati.(type) {
	case []mo.HostSystem:
		err = v.Retrieve(ctx, vimType, nil, &objs)
		if err != nil {
			return
		}
		for _, ato := range objs {
			if ato.Name == name {
				return ato, nil
			}
		}
	case []mo.VirtualMachine:
		err = v.Retrieve(ctx, vimType, nil, &objs)
		if err != nil {
			return
		}
		for _, ato := range objs {
			if ato.Name == name {
				return ato, nil
			}
		}
	case []mo.Datastore:
		err = v.Retrieve(ctx, vimType, nil, &objs)
		if err != nil {
			return
		}
		for _, ato := range objs {
			if ato.Name == name {
				return ato, nil
			}
		}
	case []mo.Network:
		err = v.Retrieve(ctx, vimType, nil, &objs)
		if err != nil {
			return
		}
		for _, ato := range objs {
			if ato.Name == name {
				return ato, nil
			}
		}
	case []mo.ComputeResource:
		err = v.Retrieve(ctx, vimType, nil, &objs)
		if err != nil {
			return
		}
		for _, ato := range objs {
			if ato.Name == name {
				return ato, nil
			}
		}
	case []mo.Folder:
		err = v.Retrieve(ctx, vimType, nil, &objs)
		if err != nil {
			return
		}
		for _, ato := range objs {
			if ato.Name == name {
				return ato, nil
			}
		}
	case []mo.ClusterComputeResource:
		err = v.Retrieve(ctx, vimType, nil, &objs)
		if err != nil {
			return
		}
		for _, ato := range objs {
			if ato.Name == name {
				return ato, nil
			}
		}
	}
	return ato, fmt.Errorf("No object found for %s", name)
}

// BlockingRegisterVM registers VM with blocking function
func BlockingRegisterVM(folder object.Folder, path string, host object.HostSystem, c *govmomi.Client, ctx context.Context) error {

	//find the root ResourcePool of the host's parent cluster
	//First need to check that host is in a cluster.
	//If not, return error
	var mh mo.HostSystem
	err := host.Properties(ctx, host.Reference(), []string{"parent"}, &mh)
	if err != nil {
		return err
	}
	if mh.Parent.Type == "ClusterComputeResource" {
		objCluster := object.NewClusterComputeResource(c.Client, *mh.Parent)
		pool, err := objCluster.ResourcePool(ctx)
		if err != nil {
			return err
		}
		task, err := folder.RegisterVM(ctx, path, "", false, pool, &host)
		if err := task.Wait(ctx); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("Host Parent is not a Cluster Compute Resource, it is %s", mh.Parent.Type)
	}
	return nil

}

// GetVMWithStatus returns list of orphaned VMs
func GetVMWithStatus(clusterName string, status string, c *govmomi.Client, ctx context.Context) (vmList []VM, err error) {

	var clusters []mo.ClusterComputeResource
	i, err := GetObjectFromName(clusterName, []string{"ClusterComputeResource"}, c, ctx, clusters)
	if err != nil {
		return
	}
	clusterMo, ok := i.(mo.ClusterComputeResource)
	if !ok {
		err = errors.New("Can't convert interface to Cluster Compute Resource")
		return
	}
	cluster := clusterMo.Reference()
	resourcePool, err := object.NewClusterComputeResource(c.Client, cluster).ResourcePool(ctx)
	m := view.NewManager(c.Client)
	v, err := m.CreateContainerView(ctx, c.ServiceContent.RootFolder, []string{"VirtualMachine"}, true)
	if err != nil {
		return
	}
	defer v.Destroy(ctx)
	var vmMos []mo.VirtualMachine
	err = v.Retrieve(ctx, []string{"VirtualMachine"}, nil, &vmMos)
	if err != nil {
		return
	}
	var vmResourcePool *object.ResourcePool
	var vmResourcePoolName, resourcePoolName string
	for _, vm := range vmMos {
		vmObj := object.NewVirtualMachine(c.Client, vm.Reference())
		vmResourcePool, err = vmObj.ResourcePool(ctx)
		if err != nil {
			if err.Error() == "VM doesn't have a resourcePool" {
				continue
			}
			return
		}
		vmResourcePoolName, err = vmResourcePool.ObjectName(ctx)
		if err != nil {
			return
		}
		resourcePoolName, err = resourcePool.ObjectName(ctx)
		if err != nil {
			return
		}
		var hostPtr *object.HostSystem
		if vmResourcePoolName == resourcePoolName && vm.Summary.Runtime.ConnectionState == status {
			vmFolder := object.NewFolder(c.Client, *vm.Parent)
			e := VM{}
			e.Object = *vmObj
			e.Name = vm.Name
			e.Folder = *vmFolder
			hostPtr, err = vmObj.HostSystem(ctx)
			if err != nil {
				return
			}
			e.Host = *hostPtr
			e.Path = vm.Config.Files.VmPathName
			e.IsOrphaned = true
			vmList = append(vmList, e)
		}
	}
	return
}
