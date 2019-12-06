package ftdi

// #include "ftd2xx.h"
// #include <stdlib.h>
import "C"
import (
	"time"
	"unsafe"
)

// Handle D2XX handle
type Handle struct{ handle C.FT_HANDLE }

// Open Open D2XX by index from GetDeviceInfoList
func Open(index int) (*Handle, error) {
	var h C.FT_HANDLE
	err := errorStatus(C.FT_Open(C.int(index), &h))
	if err != nil {
		return nil, err
	}
	return &Handle{h}, nil
}

// OpenEx Open the specified device and return a handle that will be used for subsequent accesses. The device can be specified by its serial number, device description or location.
//
// This function can also be used to open multiple devices simultaneously. Multiple devices can be specified by serial number, device description or location ID (location information derived from the physical location of a device on USB). Location IDs for specific USB ports can be obtained using the utility USBView and are given in hexadecimal format. Location IDs for devices connected to a system can be obtained by calling FT_GetDeviceInfoList or FT_ListDevices with the appropriate flags.
func OpenEx(s string, flags C.DWORD) (*Handle, error) {
	pvArg1 := C.CString(s)
	defer C.free(unsafe.Pointer(pvArg1))
	var h C.FT_HANDLE
	err := errorStatus(C.FT_OpenEx((C.PVOID)(pvArg1), flags, &h))
	if err != nil {
		return nil, err
	}
	return &Handle{h}, nil
}

// OpenBySerialNumber shorthand for OpenEx(sn, FT_OPEN_BY_SERIAL_NUMBER)
func OpenBySerialNumber(sn string) (*Handle, error) { return OpenEx(sn, FT_OPEN_BY_SERIAL_NUMBER) }

// OpenByDescription shorthand for OpenEx(desc, FT_OPEN_BY_DESCRIPTION)
func OpenByDescription(desc string) (*Handle, error) { return OpenEx(desc, FT_OPEN_BY_DESCRIPTION) }

// OpenByLocation shorthand for OpenEx(loc, FT_OPEN_BY_LOCATION)
func OpenByLocation(loc string) (*Handle, error) { return OpenEx(loc, FT_OPEN_BY_LOCATION) }

// SetDataCharacteristics Set data characteristics. Tips: call immediately after Open
func (h *Handle) SetDataCharacteristics(wordLength, stopBits, parity C.uchar) error {
	err := errorStatus(C.FT_SetDataCharacteristics(h.handle, wordLength, stopBits, parity))
	if err != nil {
		return err
	}
	return nil
}

// SetBaudRate Set BAUD rate
func (h *Handle) SetBaudRate(baud int) error {
	return errorStatus(C.FT_SetBaudRate(h.handle, C.DWORD(baud)))
}

func inMilis(d time.Duration) C.DWORD {
	return C.DWORD(d / time.Millisecond)
}

// SetTimeouts This function sets the read and write timeouts for the device.
func (h *Handle) SetTimeouts(rTimeout, wTimeout time.Duration) error {
	return errorStatus(C.FT_SetTimeouts(h.handle, inMilis(rTimeout), inMilis(wTimeout)))
}

// Close Close handle
func (h *Handle) Close() error {
	err := errorStatus(C.FT_Close(h.handle))
	if err != nil {
		return err
	}
	return nil
}

// Write Write to handle
func (h *Handle) Write(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}
	var n C.DWORD
	err := errorStatus(C.FT_Write(h.handle, C.LPVOID(&p[0]), C.DWORD(len(p)), &n))
	if err != nil {
		return int(n), err
	}
	return int(n), nil
}

// Read Read from handle
func (h *Handle) Read(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}
	var n C.DWORD
	err := errorStatus(C.FT_Read(h.handle, C.LPVOID(&p[0]), C.DWORD(len(p)), &n))
	if err != nil {
		return int(n), err
	}
	return int(n), nil
}

// func (h *Handle) SetDeadline(t time.Time) error {
// 	d := t.Sub(time.Now())
// 	return h.SetTimeouts(d, d)
// }
// func (h *Handle) SetReadDeadline(t time.Time) error {
// 	return h.SetDeadline(t)
// }
// func (h *Handle) SetWriteDeadline(t time.Time) error {
// 	return h.SetDeadline(t)
// }
