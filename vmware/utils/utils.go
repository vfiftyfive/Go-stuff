package utils

import (
	"context"
	"errors"
	"fmt"
	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25/mo"
	"reflect"
)

type VM struct {
	Name       string
	Folder     object.Folder
	Host       object.HostSystem
	IsOrphaned bool
}

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

func BlockingRegisterVM(folder string, path string, host string, c *govmomi.Client, ctx context.Context) error {

	var folders []mo.Folder
	i, err := GetObjectFromName(folder, []string{"Folder"}, c, ctx, folders)
	if err != nil {
		return err
	}
	vmFolder, ok := i.(mo.Folder)
	if !ok {
		return errors.New("Can't convert interface to Virtual Machine Folder")
	}
	objVmFolder := object.NewFolder(c.Client, vmFolder.Reference())

	//find the root ResourcePool of the host's parent cluster
	//First need to check that host is in a cluster.
	//If not, return error
	var hosts []mo.HostSystem
	i, err = GetObjectFromName(host, []string{"HostSystem"}, c, ctx, hosts)
	if err != nil {
		return errors.New("Host not found!")
	}
	hostMo, ok := i.(mo.HostSystem)
	if !ok {
		return errors.New("Can't convert interface to Host System")
	}
	parentType := reflect.ValueOf(hostMo.Parent.Type)
	if parentType.String() == "ClusterComputeResource" {
		objCluster := object.NewClusterComputeResource(c.Client, *hostMo.Parent)
		pool, err := objCluster.ResourcePool(ctx)
		if err != nil {
			return err
		}
		objHost := object.NewHostSystem(c.Client, hostMo.Reference())
		task, err := objVmFolder.RegisterVM(ctx, path, "", false, pool, objHost)
		if err := task.Wait(ctx); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("Host Parent is not a Cluster Compute Resource, it is %s", parentType.String())
	}
	return nil
}

func GetClusterOrphanedVMs(clusterName string, c *govmomi.Client, ctx context.Context) (vmList []VM, err error) {

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
		if vmResourcePoolName == resourcePoolName && vm.Summary.Runtime.ConnectionState == "orphaned" {
			vmFolder := object.NewFolder(c.Client, *vm.Parent)
			e := VM{}
			e.Name = vm.Name
			e.Folder = *vmFolder
			hostPtr, err = vmObj.HostSystem(ctx)
			if err != nil {
				return
			}
			e.Host = *hostPtr
			e.IsOrphaned = true
			vmList = append(vmList, e)
		}
	}
	return
}
