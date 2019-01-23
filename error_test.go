package ftdi

import "testing"

func TestError(t *testing.T) {
	if FT_OK.Error() != "FT_OK" {
		t.Log(FT_OK.Error())
		t.Fail()
	}
}
