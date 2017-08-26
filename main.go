package main

import (
	"bytes"
	"log"
	"net/http"
	"strconv"
	"encoding/base64"
	"html/template"
	"math/rand"
	"time"

	"image"
	"image/jpeg"

	"github.com/fogleman/gg"
)

const X, Y = 280, 240

func printImage(respWr http.ResponseWriter, req *http.Request, d *Data) {
    dc := gg.NewContext(X, Y)
    dc.SetRGB(1, 1, 1)
    dc.Clear()

    for _, c := range d.circles {
	    dc.DrawCircle(float64(c.x), float64(c.y), float64(c.r))
	    dc.SetRGB(float64(c.red), float64(c.green), float64(c.blue))
	    dc.FillPreserve()
	    dc.Stroke()
	}

    for _, r := range d.rectangles {
	    dc.DrawRectangle(float64(r.x), float64(r.y), float64(r.w), float64(r.h))
	    dc.SetRGB(float64(r.red), float64(r.green), float64(r.blue))
	    dc.FillPreserve()
	    dc.Stroke()
	}

    for _, t := range d.triangles {
	    dc.LineTo(float64(t.p1), float64(t.p2))
	    dc.LineTo(float64(t.p2), float64(t.p3))
	    dc.LineTo(float64(t.p3), float64(t.p1))
	    dc.SetRGB(float64(t.red), float64(t.green), float64(t.blue))
	    dc.FillPreserve()
	    dc.Stroke()
	}

	var img image.Image = dc.Image()
	writeImageWithTemplate(respWr, &img)
}

func createData() *Data {
	now := time.Now()
	nCircles := 3 // now.Hour()
	nRectangles := now.Minute()
	nTriangles := now.Second()
	rd := rand.New(rand.NewSource(now.UnixNano()))

    // numero de retangulos, baseado na hora
	circles := make([]Circle, nCircles)
    for i := 0; i < nCircles; i++ {
		circle := Circle{x: uint8(rd.Intn(X)), y: uint8(rd.Intn(Y)), r: uint8(rd.Intn(50)), red: uint8(rd.Intn(255)), green: uint8(rd.Intn(255)), blue: uint8(rd.Intn(255))}
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

func imageHandler(w http.ResponseWriter, r *http.Request) {
	app := &App{}
	app.Run(w, r, createData())
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
