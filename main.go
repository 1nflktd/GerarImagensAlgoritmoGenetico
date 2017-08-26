package main

import (
	"bytes"
	"log"
	"net/http"
	"strconv"
	"encoding/base64"
	"html/template"

	"image"
	"image/jpeg"
)

const X, Y = 280, 240

func imageHandler(w http.ResponseWriter, r *http.Request) {
	app := &App{}
	app.Run(w, r, CreateData())
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
