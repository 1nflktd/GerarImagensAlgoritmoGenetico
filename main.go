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
	"encoding/base64"
	"html/template"
	//"math/rand"
	//"time"
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

func printImage(respWr http.ResponseWriter, req *http.Request, d *Data) {
	//rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	var w, h int = 280, 240
	//var hw, hh float64 = float64(rd.Intn(w) / 2), float64(rd.Intn(h) / 2)
	var hw, hh float64 = float64(d.hw), float64(d.hh)
	//r := 40.0
	r := float64(d.r)
	//ang := 2 * math.Pi / 3
	ang := float64(d.ang)
	//ang := float64(rd.Intn(90))
	cr := &Circle{hw - r*math.Sin(0), hh - r*math.Cos(0), 60}
	cg := &Circle{hw - r*math.Sin(ang), hh - r*math.Cos(ang), 60}
	cb := &Circle{hw - r*math.Sin(-ang), hh - r*math.Cos(-ang), 60}
	m := image.NewRGBA(image.Rect(0, 0, w, h))
	for x := 1; x < w; x++ {
		for y := 1; y < h; y++ {
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
	//writeImage(respWr, &img)
	writeImageWithTemplate(respWr, &img)
}

func imageHandler(w http.ResponseWriter, r *http.Request) {
	app := &App{}
	app.Run(w, r, &Data{hw: 140, hh:120, r: 40, ang: 60, a: 255})
}

var ImageTemplate string = `<!DOCTYPE html>
	<html lang="en"><head></head>
	<body><img src="data:image/jpg;base64,{{.Image}}"></body>`

func writeImageWithTemplate(w http.ResponseWriter, img *image.Image) {
	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, *img, nil); err != nil {
		log.Println("unable to encode image.")
	}

	str := base64.StdEncoding.EncodeToString(buffer.Bytes())

	if tmpl, err := template.New("image").Parse(ImageTemplate); err != nil {
		log.Println("unable to parse image template.")
	} else {
		data := map[string]interface{}{"Image": str}
		if err = tmpl.Execute(w, data); err != nil {
			log.Println("unable to execute template.")
		}
	}
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
	http.HandleFunc("/image/", imageHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
