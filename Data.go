package main

import (
	"fmt"
	"strconv"
)

// preparado para 0 até 255
type Data struct {
	//hw, hh, r, ang, crx, cry, cgx, cgy, cbx, cby, a int
	hw, hh int
}

func (d *Data) toString() string {
	return fmt.Sprintf("%x%x", d.hw, d.hh)
}

func (d *Data) hexToInt(hex string) int {
	n, _ := strconv.ParseUint(hex, 16, 32)
	return int(n)
}

func (d *Data) fromString(data string) {
	d.hw = d.hexToInt(data[0:2])
	d.hh = d.hexToInt(data[2:4])
}
