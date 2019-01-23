package ftdi

// #include "ftd2xx.h"
import "C"

func CreateDeviceInfoList() (int, Status) {
	var n C.uint
	status := Status(C.FT_CreateDeviceInfoList(&n))
	return int(n), status
}
func GetDeviceInfoList() ([]Device, Status) {
	ndev, status := CreateDeviceInfoList()
	if status != FT_OK {
		return nil, status
	}
	if ndev <= 0 {
		return nil, status
	}
	out := make([]Device, ndev)
	var n C.uint
	n = C.uint(ndev)
	status = Status(C.FT_GetDeviceInfoList((*C.FT_DEVICE_LIST_INFO_NODE)(&out[0]), &n))
	return out, status
}
