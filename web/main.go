package web

import (
	"fmt"
	"log"
	"html/template"
	"image"
	"net/http"
	"image/jpeg"
	"bytes"
	"encoding/base64"
)

// TODO add template?
var errorResponse string = `Oh noes, generic error page: `

var imageTemplate string = `<!DOCTYPE html>
<html lang="en"><head></head>
<body><img src="data:image/jpg;base64,{{.Image}}"></body>`

func GetErrorResponse(w http.ResponseWriter, err error) {
	// TODO error code?
	fmt.Fprintf(w, errorResponse, err)
}

func WriteTemplateWithImage(w http.ResponseWriter, img *image.Image) {
	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, *img, nil); err != nil {
		log.Fatalln("Unable to encode image.")
	}

	str := base64.StdEncoding.EncodeToString(buffer.Bytes())
	if tmpl, err := template.New("image").Parse(imageTemplate); err != nil {
		log.Println("Unable to parse image template.")
	} else {
		data := map[string]interface{}{"Image": str}
		if err = tmpl.Execute(w, data); err != nil {
			log.Println("Unable to execute template.")
		}
	}
}