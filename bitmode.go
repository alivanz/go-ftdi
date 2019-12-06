package ftdi

// #include "ftd2xx.h"
import "C"

const (
	// BitModeReset Reset
	BitModeReset byte = C.FT_BITMODE_RESET
	// BitModeAsyncBitBang Asynchronous Bit Bang
	BitModeAsyncBitBang byte = C.FT_BITMODE_ASYNC_BITBANG
	// BitModeMPSSE MPSSE (FT2232, FT2232H, FT4232H and FT232H devices only)
	BitModeMPSSE byte = C.FT_BITMODE_MPSSE
	// BitModeSyncBitBang Synchronous Bit Bang (FT232R, FT245R, FT2232, FT2232H, FT4232H and FT232H devices only)
	BitModeSyncBitBang byte = C.FT_BITMODE_SYNC_BITBANG
	// BitModeMCUHost MCU Host Bus Emulation Mode (FT2232, FT2232H, FT4232H and FT232H devices only)
	BitModeMCUHost byte = C.FT_BITMODE_MCU_HOST
	// BitModeFastSerial Fast Opto-Isolated Serial Mode (FT2232, FT2232H, FT4232H and FT232H devices only)
	BitModeFastSerial byte = C.FT_BITMODE_FAST_SERIAL
	// BitModeCBusBitBang CBUS Bit Bang Mode (FT232R and FT232H devices only)
	BitModeCBusBitBang byte = C.FT_BITMODE_CBUS_BITBANG
	// BitModeSyncFIFO Single Channel Synchronous 245 FIFO Mode (FT2232H and FT232H devices only)
	BitModeSyncFIFO byte = C.FT_BITMODE_SYNC_FIFO
)

// SetBitMode Enables different chip modes.
// mask:  Required value for bit mode mask. This sets up which bits are inputs and outputs. A bit value of 0 sets the corresponding pin to an input, a bit value of 1 sets the corresponding pin to an output.
//        In the case of CBUS Bit Bang, the upper nibble of this value controls which pins are inputs and outputs, while the lower nibble controls which of the outputs are high and low.
// mode:  Mode value.
func (h *Handle) SetBitMode(mask, mode byte) error {
	return errorStatus(C.FT_SetBitMode(h.handle, C.UCHAR(mask), C.UCHAR(mode)))
}

// GetBitMode Gets the instantaneous value of the data bus.
func (h *Handle) GetBitMode() (byte, error) {
	var mode C.UCHAR
	return byte(mode), errorStatus(C.FT_GetBitMode(h.handle, &mode))
}
