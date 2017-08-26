package main

import (
	"fmt"
	"strconv"
	"math/rand"
	"time"
)

type Circle struct {
	x, y, r, red, green, blue uint8
}

type Rectangle struct {
	x, y, w, h, red, green, blue uint8
}

type Triangle struct {
	p1, p2, p3, red, green, blue uint8
}

// preparado para 0 até 255
type Data struct {
	circles []Circle
	rectangles []Rectangle
	triangles []Triangle
	nCircles, nRectangles, nTriangles int
}

func NewData(circles []Circle, rectangles []Rectangle, triangles []Triangle, nCircles, nRectangles, nTriangles int) *Data {
	return &Data{circles, rectangles, triangles, nCircles, nRectangles, nTriangles}
}

func (d *Data) toString() string {
	hex := ""

	for _, c := range d.circles {
		hex += fmt.Sprintf("%02x%02x%02x%02x%02x%02x", c.x, c.y, c.r, c.red, c.green, c.blue)
	}

	for _, r := range d.rectangles {
		hex += fmt.Sprintf("%02x%02x%02x%02x%02x%02x%02x", r.x, r.y, r.w, r.h, r.red, r.green, r.blue)
	}

	for _, t := range d.triangles {
		hex += fmt.Sprintf("%02x%02x%02x%02x%02x%02x", t.p1, t.p2, t.p3, t.red, t.green, t.blue)
	}

	return hex
}

func (d *Data) hexToUint(hex string) uint8 {
	n, _ := strconv.ParseUint(hex, 16, 32)
	return uint8(n)
}

func (d *Data) fromString(data string, nCircles, nRectangles, nTriangles int) {
	d.nCircles, d.nRectangles, d.nTriangles = nCircles, nRectangles, nTriangles
	d.circles = make([]Circle, nCircles)
	for i := 0; i < nCircles; i++ {
		base := i * 6 // 6 fields
		circle := Circle{
			d.hexToUint(data[base:base+2]),
			d.hexToUint(data[base+2:base+4]),
			d.hexToUint(data[base+4:base+6]),
			d.hexToUint(data[base+6:base+8]),
			d.hexToUint(data[base+8:base+10]),
			d.hexToUint(data[base+10:base+12]),
		}
		d.circles[i] = circle
	}

	d.rectangles = make([]Rectangle, nRectangles)
	iniR := (nCircles * 6) + 12
	for i := 0; i < nRectangles; i++ {
		base := iniR + (i * 7) // 7 fields
		rectangle := Rectangle{
			d.hexToUint(data[base:base+2]),
			d.hexToUint(data[base+2:base+4]),
			d.hexToUint(data[base+4:base+6]),
			d.hexToUint(data[base+6:base+8]),
			d.hexToUint(data[base+8:base+10]),
			d.hexToUint(data[base+10:base+12]),
			d.hexToUint(data[base+12:base+14]),
		}
		d.rectangles[i] = rectangle
	}

	d.triangles = make([]Triangle, nTriangles)
	iniT := iniR + (nRectangles * 7) + 14
	for i := 0; i < nTriangles; i++ {
		base := iniT + (i * 6) // 6 fields
		triangle := Triangle{
			d.hexToUint(data[base:base+2]),
			d.hexToUint(data[base+2:base+4]),
			d.hexToUint(data[base+4:base+6]),
			d.hexToUint(data[base+6:base+8]),
			d.hexToUint(data[base+8:base+10]),
			d.hexToUint(data[base+10:base+12]),
		}
		d.triangles[i] = triangle
	}
}

func CreateData() *Data {
	now := time.Now()
	nCircles := 3 // now.Hour()
	nRectangles := now.Minute()
	nTriangles := now.Second()
	rd := rand.New(rand.NewSource(now.UnixNano()))

    // numero de retangulos, baseado na hora
	circles := make([]Circle, nCircles)
    for i := 0; i < nCircles; i++ {
		circle := Circle{uint8(rd.Intn(X)), uint8(rd.Intn(Y)), uint8(rd.Intn(50)), uint8(rd.Intn(255)), uint8(rd.Intn(255)), uint8(rd.Intn(255))}
		circles[i] = circle
	}

    // numero de circulos, baseado nos minutos
	rectangles := make([]Rectangle, nRectangles)
    for i := 0; i < nRectangles; i++ {
		rectangle := Rectangle{uint8(rd.Intn(X)), uint8(rd.Intn(Y)), uint8(rd.Intn(X/5)), uint8(rd.Intn(Y/5)), uint8(rd.Intn(255)), uint8(rd.Intn(255)), uint8(rd.Intn(255))}
		rectangles[i] = rectangle
	}

    // numero de triangulos, baseado nos segundos
	triangles := make([]Triangle, nTriangles)
    for i := 0; i < nTriangles; i++ {
		triangle := Triangle{uint8(rd.Intn(X/10)), uint8(rd.Intn(Y/10)), uint8(rd.Intn(Y/10)), uint8(rd.Intn(255)), uint8(rd.Intn(255)), uint8(rd.Intn(255))}
		triangles[i] = triangle
	}

	return NewData(circles, rectangles, triangles, nCircles, nRectangles, nTriangles)
}

