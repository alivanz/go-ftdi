package ftdi

import (
	"testing"
)

func TestUnit(t *testing.T) {
	var numDev uint
	CreateDeviceInfoList(&numDev)
}
