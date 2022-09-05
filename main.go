package main

import (
	"log"
	"net/http"
	"github.com/zinah/memegen/pictures"
	"github.com/zinah/memegen/captions"
	"github.com/zinah/memegen/web"
)

func handler(w http.ResponseWriter, r *http.Request) {
	textAbove, textBelow := captions.GetJoke()
	if bgImage, err := pictures.GetImage("https://cataas.com/cat", "./assets/default_cat.jpeg"); err != nil {
		web.GetErrorResponse(w, err)
	} else {
		if img, err := pictures.ApplyCaption(bgImage, textAbove, textBelow); err != nil {
			// TODO at this point the background image is loaded successfully, maybe at least show that?
			web.GetErrorResponse(w, err)
		} else {
			web.WriteTemplateWithImage(w, &img)
		}
	}
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))

}