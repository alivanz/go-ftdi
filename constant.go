package ftdi

// #include "ftd2xx.h"
import "C"

const (
	// Word length
	FT_BITS_8 = C.FT_BITS_8
	FT_BITS_7 = C.FT_BITS_7

	// Stop bits
	FT_STOP_BITS_1 = C.FT_STOP_BITS_1
	FT_STOP_BITS_2 = C.FT_STOP_BITS_2

	// Parity
	FT_PARITY_NONE  = C.FT_PARITY_NONE
	FT_PARITY_ODD   = C.FT_PARITY_ODD
	FT_PARITY_EVEN  = C.FT_PARITY_EVEN
	FT_PARITY_MARK  = C.FT_PARITY_MARK
	FT_PARITY_SPACE = C.FT_PARITY_SPACE
)
