package uvm

import (
	"context"
	"fmt"

	"github.com/Microsoft/hcsshim/internal/guestrequest"
	"github.com/Microsoft/hcsshim/internal/requesttype"
	hcsschema "github.com/Microsoft/hcsshim/internal/schema2"
	"github.com/google/uuid"
)

const (
	resourcePathVpciFmt = "VirtualMachine/Devices/VirtualPci/%s"
)

func (uvm *UtilityVM) AssignDevice(ctx context.Context, device hcsschema.VirtualPciDevice) (string, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	id := uuid.String()

	uvm.m.Lock()
	defer uvm.m.Unlock()
	return id, uvm.modify(ctx, &hcsschema.ModifySettingRequest{
		ResourcePath: fmt.Sprintf(resourcePathVpciFmt, id),
		RequestType:  requesttype.Add,
		Settings:     device,
		GuestRequest: guestrequest.GuestRequest{
			ResourceType: guestrequest.ResourceTypeVPciDevice,
			RequestType:  requesttype.Add,
			Settings: guestrequest.LCOWMappedVPciDevice{
				VmbusUuid: id,
			},
		},
	})
}

func (uvm *UtilityVM) RemoveDevice(ctx context.Context, id string) error {
	uvm.m.Lock()
	defer uvm.m.Unlock()
	return uvm.modify(ctx, &hcsschema.ModifySettingRequest{
		ResourcePath: fmt.Sprintf(resourcePathVpciFmt, id),
		RequestType:  requesttype.Remove,
	})
}
