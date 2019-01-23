package ftdi

// #include "ftd2xx.h"
import "C"

// CreateDeviceInfoList This function builds a device information list and returns the number of D2XX devices connected to the system. The list contains information about both unopen and open devices.
func CreateDeviceInfoList() (int, Status) {
	var n C.uint
	status := Status(C.FT_CreateDeviceInfoList(&n))
	return int(n), status
}

// GetDeviceInfoList This function returns a device information list and the number of D2XX devices in the list.
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

// GetDeviceInfoDetail This function returns an entry from the device information list.
func GetDeviceInfoDetail(index int, pDev *Device) Status {
	cindex := C.uint(index)
	return Status(C.FT_GetDeviceInfoDetail(cindex, &pDev.Flags, &pDev.Type, &pDev.ID, &pDev.LocId, (C.PVOID)(&pDev.SerialNumber[0]), (C.PVOID)(&pDev.Description[0]), &pDev.ftHandle))
}
