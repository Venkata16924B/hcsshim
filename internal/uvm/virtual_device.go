package uvm

import (
	"context"
	"fmt"

	"github.com/Microsoft/hcsshim/internal/guestrequest"
	"github.com/Microsoft/hcsshim/internal/requesttype"
	hcsschema "github.com/Microsoft/hcsshim/internal/schema2"
)

const (
	resourcePathVpciFmt = "VirtualMachine/Devices/VirtualPci/%s"
)

func (uvm *UtilityVM) AssignDevice(ctx context.Context, id string, device hcsschema.VirtualPciDevice) error {
	uvm.m.Lock()
	defer uvm.m.Unlock()

	return uvm.modify(ctx, &hcsschema.ModifySettingRequest{
		ResourcePath: fmt.Sprintf(resourcePathVpciFmt, id),
		RequestType:  requesttype.Add,
		Settings:     device,
		GuestRequest: guestrequest.GuestRequest{
			ResourceType: guestrequest.ResourceTypeVPciDevice,
			RequestType:  requesttype.Add,
			Settings:     nil, /* TODO - Guest request */
		},
	})
}

func (uvm *UtilityVM) RemoveDevice(ctx context.Context, id string) error {
	uvm.m.Lock()
	defer uvm.m.Unlock()

	return uvm.modify(ctx, &hcsschema.ModifySettingRequest{
		ResourcePath: fmt.Sprintf(resourcePathVpciFmt, id),
		RequestType:  requesttype.Remove,
		GuestRequest: guestrequest.GuestRequest{
			ResourceType: guestrequest.ResourceTypeVPciDevice,
			RequestType:  requesttype.Remove,
			Settings:     nil, /* TODO - Guest request */
		},
	})
}
