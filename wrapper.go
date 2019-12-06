package ftdi

// #ifdef _WIN32
// #define FTD2XX_STATIC
// #endif
// #include "ftd2xx.h"
import "C"

// CreateDeviceInfoList This function builds a device information list and returns the number of D2XX devices connected to the system. The list contains information about both unopen and open devices.
func CreateDeviceInfoList() (int, error) {
	var n C.DWORD
	if err := errorStatus(C.FT_CreateDeviceInfoList(&n)); err != nil {
		return 0, err
	}
	return int(n), nil
}

// GetDeviceInfoList This function returns a device information list and the number of D2XX devices in the list.
func GetDeviceInfoList() ([]Device, error) {
	ndev, err := CreateDeviceInfoList()
	if err != nil {
		return nil, err
	}
	if ndev <= 0 {
		return nil, nil
	}
	out := make([]Device, ndev)
	n := C.DWORD(ndev)
	if err := errorStatus(C.FT_GetDeviceInfoList((*C.FT_DEVICE_LIST_INFO_NODE)(&out[0].device), &n)); err != nil {
		return nil, err
	}
	return out, nil
}

// GetDeviceInfoDetail This function returns an entry from the device information list.
func GetDeviceInfoDetail(index int, pDev *Device) error {
	cindex := C.DWORD(index)
	return errorStatus(C.FT_GetDeviceInfoDetail(
		cindex,
		&pDev.device.Flags,
		&pDev.Type,
		&pDev.ID,
		&pDev.LocId,
		(C.LPVOID)(&pDev.device.SerialNumber[0]),
		(C.LPVOID)(&pDev.device.Description[0]),
		&pDev.ftHandle,
	))
}
