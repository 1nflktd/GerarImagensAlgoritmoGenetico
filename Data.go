package main

import (
	"fmt"
	"strconv"
)

// preparado para 0 até 255
type Data struct {
	hours, minutes, seconds, red, green, blue uint8
}

func NewData(hours, minutes, seconds, red, green, blue uint8) *Data {
	return &Data{hours: hours, minutes: minutes, seconds: seconds, red: red, green: green, blue: blue}
}

func (d *Data) toString() string {
	return fmt.Sprintf("%02x%02x%02x%02x%02x%02x", d.hours, d.minutes, d.seconds, d.red, d.green, d.blue)
}

func (d *Data) hexToUint(hex string) uint8 {
	n, _ := strconv.ParseUint(hex, 16, 32)
	return uint8(n)
}

func (d *Data) fromString(data string) {
	d.hours = d.hexToUint(data[0:2])
	d.minutes = d.hexToUint(data[2:4])
	d.seconds = d.hexToUint(data[4:6])
	d.red = d.hexToUint(data[6:8])
	d.green = d.hexToUint(data[8:10])
	d.blue = d.hexToUint(data[10:12])
}
