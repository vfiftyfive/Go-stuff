package utils

import (
	"context"
	"errors"
	//"fmt"
	"github.com/vmware/govmomi"
	//"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25/mo"
)

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
	}
	return ato, errors.New("No object found")
}

// func BlockingRegisterVM(folder object.folder, path string, host object.HostSystem, c *govmomi.Client, ctx context.Context) error {

// 	var folders []mo.Folder
// 	found := false
// 	i, err := GetObjectFromName(folder.Name, []string{"Folder"}, c, ctx, folders)
// 	// for _, dstFolder := range c.ServiceContent.RootFolder {
// 	// 	if folder.Name == dstFolder {
// 	// 		found = true
// 	// 		break
// 	// 	}
// 	// }
// 	if found == false {
// 		return errors.New("Folder not Found")
// 	}

// 	task, err := folder.RegisterVM(ctx, path, "", false, nil, &host)
// 	if err != nil {
// 		return err
// 	}
// 	return task.Wait(ctx)
// }
