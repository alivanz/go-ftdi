package ftdi

import (
	"testing"
	"time"
)

func TestUnit(t *testing.T) {
	CreateDeviceInfoList()
	GetDeviceInfoList()
}

func TestValue(t *testing.T) {
	x := int(inMilis(time.Second))
	if x != 1000 {
		t.Log(x)
		t.Fail()
	}
}
