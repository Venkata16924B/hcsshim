package uvm

import (
	"context"
	"fmt"
	"sync/atomic"

	"github.com/Microsoft/hcsshim/internal/requesttype"
	hcsschema "github.com/Microsoft/hcsshim/internal/schema2"
)

const (
	resourcePathVpciFmt = "VirtualMachine/Devices/VirtualPci/%s"
)

func (uvm *UtilityVM) GpuCounter() uint64 {
	return atomic.AddUint64(&uvm.gpuCounter, 1)
}

func (uvm *UtilityVM) AssignDevice(ctx context.Context, id string, device hcsschema.VirtualPciDevice) error {
	uvm.m.Lock()
	defer uvm.m.Unlock()

	return uvm.modify(ctx, &hcsschema.ModifySettingRequest{
		ResourcePath: fmt.Sprintf(resourcePathVpciFmt, id),
		RequestType:  requesttype.Add,
		Settings:     device,
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
