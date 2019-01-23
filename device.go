package ftdi

// #include "ftd2xx.h"
import "C"

type Device C.FT_DEVICE_LIST_INFO_NODE

func (d Device) Flag() int {
	return int(d.Flags)
}
func (d Device) SerialNum() string {
	return C.GoString(&d.SerialNumber[0])
}
func (d Device) Desc() string {
	return C.GoString(&d.Description[0])
}
func (d Device) String() string {
	return d.SerialNum()
}
