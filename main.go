package main

import (
	"bytes"
	"log"
	"net/http"
	"strconv"
	"math"
	"encoding/base64"
	"html/template"
	"math/rand"
	"time"

	"image"
	"image/jpeg"

	"github.com/fogleman/gg"
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
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))

	const X, Y = 280, 240
    dc := gg.NewContext(280, 240)
    dc.SetRGB(1, 1, 1)
    dc.Clear()

    // numero de retangulos, baseado na hora
    for i := uint8(0); i < d.hours; i++ {
	    red := rd.Float64() * float64(d.red)
	    green := rd.Float64() * float64(d.green)
	    blue := rd.Float64() * float64(d.blue)

	    dc.DrawCircle(float64(rd.Intn(X)), float64(rd.Intn(Y)), float64(rd.Intn(5)))
	    dc.SetRGB(red/255, green/255, blue/255)
	    dc.FillPreserve()
	    dc.Stroke()
	}

    // numero de circulos, baseado nos minutos
    for i := uint8(0); i < d.minutes; i++ {
	    red := rd.Float64() * float64(d.red)
	    green := rd.Float64() * float64(d.green)
	    blue := rd.Float64() * float64(d.blue)

	    dc.DrawRectangle(float64(rd.Intn(X)), float64(rd.Intn(Y)), float64(rd.Intn(X/5)), float64(rd.Intn(Y/5)))
	    dc.SetRGB(red/255, green/255, blue/255)
	    dc.FillPreserve()
	    dc.Stroke()
	}

    // numero de triangulos, baseado nos segundos
    for i := uint8(0); i < d.seconds; i++ {
	    red := rd.Float64() * float64(d.red)
	    green := rd.Float64() * float64(d.green)
	    blue := rd.Float64() * float64(d.blue)

    	x1, y1, y2 := float64(rd.Intn(X/10)), float64(rd.Intn(Y/10)), float64(rd.Intn(Y/10))
	    dc.LineTo(x1, y1)
	    dc.LineTo(y1, y2)
	    dc.LineTo(y2, x1)
	    dc.SetRGB(red/255, green/255, blue/255)
	    dc.FillPreserve()
	    dc.Stroke()
	}

	var img image.Image = dc.Image()
	writeImageWithTemplate(respWr, &img)
}

func imageHandler(w http.ResponseWriter, r *http.Request) {
	app := &App{}

	red := uint8(152)
	green := uint8(132)
	blue := uint8(255)

	now := time.Now()
	app.Run(w, r, NewData(uint8(now.Hour()), uint8(now.Minute()), uint8(now.Second()), red, green, blue))
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
