package main

import (
	"bytes"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"net/http"
	"strconv"
	"math"
)

type Circle struct {
	X, Y, R float64
}

func (c *Circle) Brightness(x, y float64) uint8 {
	var dx, dy float64 = c.X - x, c.Y - y
	d := math.Sqrt(dx * dx + dy * dy) / c.R
	if d > 1 {
		return 0
	} else {
		return 255
	}
}

func imageHandler(respWr http.ResponseWriter, req *http.Request) {
	var w, h int = 280, 240
	var hw, hh float64 = float64(w / 2), float64(h / 2)
	r := 40.0
	ang := 2 * math.Pi / 3
	cr := &Circle{hw - r*math.Sin(0), hh - r*math.Cos(0), 60}
	cg := &Circle{hw - r*math.Sin(ang), hh - r*math.Cos(ang), 60}
	cb := &Circle{hw - r*math.Sin(-ang), hh - r*math.Cos(-ang), 60}

	m := image.NewRGBA(image.Rect(0, 0, w, h))
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			c := color.RGBA{
				cr.Brightness(float64(x), float64(y)),
				cg.Brightness(float64(x), float64(y)),
				cb.Brightness(float64(x), float64(y)),
				255,
			}
			m.Set(x, y, c)
		}
	}

	var img image.Image = m
	writeImage(respWr, &img)
}

// writeImage escreve uma imagem no response writer
func writeImage(w http.ResponseWriter, img *image.Image) {
	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, *img, nil); err != nil {
		log.Println("Nao foi possivel codificar a imagem.")
	}

	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Println("Nao foi possivel escrever a imagem.")
	}
}

func main() {
	//http.HandleFunc("/image/", imageHandler)
	//log.Fatal(http.ListenAndServe(":8000", nil))
	app := &App{}
	app.Run()
}
