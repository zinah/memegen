package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/tkanos/gonfig"
	"github.com/zinah/memegen/captions"
	"github.com/zinah/memegen/pictures"
	"github.com/zinah/memegen/translation"
	"github.com/zinah/memegen/web"
)

type Configuration struct {
	ImageSourceURL string
	DefaultImageLocalPath string
	DictionaryLocalPath string
}

func handler(w http.ResponseWriter, r *http.Request, configuration Configuration, dictionary map[string]string) {
	text1, text2 := captions.GetJoke()
	textAbove, textBelow := translation.Translate(text1, dictionary), translation.Translate(text2, dictionary)
	if bgImage, err := pictures.GetImage(configuration.ImageSourceURL, configuration.DefaultImageLocalPath); err != nil {
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
	configuration := Configuration{}
	err := gonfig.GetConf("dev_config.json", &configuration)
	if err != nil {
		panic(err)
	}
	// TODO dictionary should be in some sort of DB instead
	dictionary := translation.GetDictionaryFromJson(configuration.DictionaryLocalPath)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, configuration, dictionary)
	}).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}