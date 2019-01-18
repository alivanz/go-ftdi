package ftdi

// #include "ftd2xx.h"
import "C"

type Status C.FT_STATUS

func CreateDeviceInfoList(pnumDev *uint) Status {
	var n C.uint
	status := Status(C.FT_CreateDeviceInfoList(&n))
	*pnumDev = uint(n)
	return status
}
