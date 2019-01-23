package ftdi

// #include "ftd2xx.h"
import "C"

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
