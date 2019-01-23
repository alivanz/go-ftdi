package ftdi

// #include "ftd2xx.h"
import "C"

type device C.FT_DEVICE_LIST_INFO_NODE
type Device struct {
	device
}

func (d Device) Flags() int {
	return int(d.device.Flags)
}
func (d Device) SerialNumber() string {
	return C.GoString(&d.device.SerialNumber[0])
}
func (d Device) Description() string {
	return C.GoString(&d.device.Description[0])
}
func (d Device) String() string {
	return d.SerialNumber()
}
