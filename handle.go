package ftdi

// #include "ftd2xx.h"
// #include <stdlib.h>
import "C"
import "unsafe"

// Handle D2XX handle
type Handle struct{ handle C.FT_HANDLE }

// Open Open D2XX by index from GetDeviceInfoList
func Open(index int) (*Handle, Status) {
	var h C.FT_HANDLE
	status := Status(C.FT_Open(C.int(index), &h))
	if status != FT_OK {
		return nil, status
	}
	return &Handle{h}, status
}

// OpenEx Open the specified device and return a handle that will be used for subsequent accesses. The device can be specified by its serial number, device description or location.
//
// This function can also be used to open multiple devices simultaneously. Multiple devices can be specified by serial number, device description or location ID (location information derived from the physical location of a device on USB). Location IDs for specific USB ports can be obtained using the utility USBView and are given in hexadecimal format. Location IDs for devices connected to a system can be obtained by calling FT_GetDeviceInfoList or FT_ListDevices with the appropriate flags.
func OpenEx(s string, flags C.uint) (*Handle, Status) {
	pvArg1 := C.CString(s)
	defer C.free(unsafe.Pointer(pvArg1))
	var h C.FT_HANDLE
	status := Status(C.FT_OpenEx((C.PVOID)(pvArg1), flags, &h))
	if status != FT_OK {
		return nil, status
	}
	return &Handle{h}, status
}

// SetDataCharacteristics Set data characteristics. Tips: call immediately after Open
func (h *Handle) SetDataCharacteristics(wordLength, stopBits, parity C.uchar) error {
	status := Status(C.FT_SetDataCharacteristics(h.handle, wordLength, stopBits, parity))
	if status != FT_OK {
		return status
	}
	return nil
}

// SetBaudRate Set BAUD rate
func (h *Handle) SetBaudRate(baud int) Status {
	return Status(C.FT_SetBaudRate(h.handle, C.uint(baud)))
}

// Close Close handle
func (h *Handle) Close() error {
	status := Status(C.FT_Close(h.handle))
	if status == FT_OK {
		return nil
	}
	return status
}

// Write Write to handle
func (h *Handle) Write(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}
	var n C.uint
	status := Status(C.FT_Write(h.handle, (C.PVOID)(&p[0]), C.uint(len(p)), &n))
	if status != FT_OK {
		return int(n), status
	}
	return int(n), nil
}

// Read Read from handle
func (h *Handle) Read(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}
	var n C.uint
	status := Status(C.FT_Read(h.handle, (C.PVOID)(&p[0]), C.uint(len(p)), &n))
	if status != FT_OK {
		return int(n), status
	}
	return int(n), nil
}
