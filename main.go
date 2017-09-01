package main

import (
	"bytes"
	"log"
	"net/http"
	"encoding/base64"
	"html/template"
	"flag"
	"strconv"

	"image"
	"image/jpeg"
)

const X, Y = 280, 240

func mainHandler(w http.ResponseWriter, r *http.Request) {
	if tmpl, err := template.New("Main.tpl").ParseFiles("Main.tpl"); err != nil {
		log.Println("unable to parse main template.", err)
	} else {
		if err = tmpl.Execute(w, nil); err != nil {
			log.Println("unable to execute template.", err)
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

	data := CreateData(r.PostForm["nome"][0])
	PrintImage(w, r, data, "Objetivo")

	taxaCrossover, errCrossover := strconv.ParseFloat(r.PostForm["taxaCrossover"][0], 64)
	if errCrossover != nil {
		log.Println("Erro ao converter valores do formulario.")
		return
	}

	taxaMutacao, errMutacao := strconv.ParseFloat(r.PostForm["taxaMutacao"][0], 64)
	if errMutacao != nil {
		log.Println("Erro ao converter valores do formulario.")
		return
	}

	elitismo := r.PostForm["elitismo"][0] == "S" 

	tamanhoPopulacao, errMPop := strconv.ParseInt(r.PostForm["tamanhoPopulacao"][0], 10, 32)
	if errMPop != nil {
		log.Println("Erro ao converter valores do formulario.")
		return
	}

	app := &App{taxaCrossover, taxaMutacao, elitismo, int(tamanhoPopulacao)}
	app.Run(w, r, data)
}

func writeImageWithTemplate(w http.ResponseWriter, img *image.Image, label string) {
	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, *img, nil); err != nil {
		log.Println("unable to encode image.")
	}

	imgBase64 := base64.StdEncoding.EncodeToString(buffer.Bytes())

	if tmpl, err := template.New("Image.tpl").ParseFiles("Image.tpl"); err != nil {
		log.Println("unable to parse image template.")
	} else {
		data := map[string]interface{}{"Image": imgBase64, "Label": label}
		if err = tmpl.Execute(w, data); err != nil {
			log.Println("unable to execute template.", err)
		}
	}
}

func main() {
	var porta = flag.String("Porta", "8000", "Digite a porta do servidor")
	flag.Parse()

	http.HandleFunc("/image", imageHandler)
	http.HandleFunc("/", mainHandler)
	log.Fatal(http.ListenAndServe(":" + *porta, nil))
}
