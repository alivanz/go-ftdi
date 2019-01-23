package ftdi

// #include "ftd2xx.h"
import "C"

type Handle struct{ handle C.FT_HANDLE }

func Open(index int) (*Handle, Status) {
	var h C.FT_HANDLE
	status := Status(C.FT_Open(C.int(index), &h))
	if status != FT_OK {
		return nil, status
	}
	return &Handle{h}, status
}
func (h *Handle) SetDataCharacteristics(wordLength, stopBits, parity C.uchar) error {
	status := Status(C.FT_SetDataCharacteristics(h.handle, wordLength, stopBits, parity))
	if status != FT_OK {
		return status
	}
	return nil
}
func (h *Handle) SetBaudRate(baud int) Status {
	return Status(C.FT_SetBaudRate(h.handle, C.uint(baud)))
}
func (h *Handle) Close() error {
	status := Status(C.FT_Close(h.handle))
	if status == FT_OK {
		return nil
	}
	return status
}

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
