package ftdi

// #include "ftd2xx.h"
import "C"

func CreateDeviceInfoList() (int, Status) {
	var n C.uint
	status := Status(C.FT_CreateDeviceInfoList(&n))
	return int(n), status
}
