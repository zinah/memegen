package main

import (
	"log"
	"net/http"
	"github.com/zinah/memegen/captions"
	"github.com/zinah/memegen/pictures"
	"github.com/zinah/memegen/translation"
	"github.com/zinah/memegen/web"
)

func handler(w http.ResponseWriter, r *http.Request) {
	dictionary := translation.GetDictionaryFromJson("./assets/tranzlashun.json")
	text1, text2 := captions.GetJoke()
	textAbove, textBelow := translation.Translate(text1, dictionary), translation.Translate(text2, dictionary)
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