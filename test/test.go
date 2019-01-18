package main

import (
	"log"

	ftdi "github.com/alivanz/go-ftdi"
)

func main() {
	var numDev uint
	ftdi.CreateDeviceInfoList(&numDev)
	log.Print(numDev)
}
