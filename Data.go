package main

import (
	"fmt"
	"strconv"
)

// preparado para 0 até 255
type Data struct {
	//crx, cry, cgx, cgy, cbx, cby, 
	hw, hh, r, ang, a int
}

func (d *Data) toString() string {
	return fmt.Sprintf("%x%x%x%x%x", d.hw, d.hh, d.r, d.ang, d.a)
}

func (d *Data) hexToInt(hex string) int {
	n, _ := strconv.ParseUint(hex, 16, 32)
	return int(n)
}

func (d *Data) fromString(data string) {
	d.hw = d.hexToInt(data[0:2])
	d.hh = d.hexToInt(data[2:4])
	d.r = d.hexToInt(data[4:6])
	d.ang = d.hexToInt(data[6:8])
	d.a = d.hexToInt(data[8:10])
}
