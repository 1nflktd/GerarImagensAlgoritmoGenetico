package main

import (
	"bytes"
	"log"
	"net/http"
	"encoding/base64"
	"html/template"

	"image"
	"image/jpeg"
)

const X, Y = 280, 240

var MainTemplate string = `<!DOCTYPE html>
	<html lang="en"><head></head>
	<body>
		<form action="/image" method="post">
			<label>Digite seu nome</label>
			<input type="text" name="nome" id="nome">
			<input type="submit" value="Enviar">
		</form>
	</body>
	</html>`

func mainHandler(w http.ResponseWriter, r *http.Request) {
	if tmpl, err := template.New("main").Parse(MainTemplate); err != nil {
		log.Println("unable to parse main template.")
	} else {
		if err = tmpl.Execute(w, nil); err != nil {
			log.Println("unable to execute template.")
		}
	}
}

func imageHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println("Erro ao obter valores do formulario")
		return
	}

	if r.PostForm["nome"] == nil {
		http.Redirect(w, r, "/", 200)
		return
	}

	app := &App{}
	app.Run(w, r, CreateData(r.PostForm["nome"][0]))
}

var ImageTemplate string = `<!DOCTYPE html>
	<html lang="en"><head></head>
	<body><img src="data:image/jpg;base64,{{.Image}}"></body>
	</html>`

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

func main() {
	http.HandleFunc("/image", imageHandler)
	http.HandleFunc("/", mainHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
