package main

import (
	"flag"
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
	CaptionSourceURL string
}

func handler(w http.ResponseWriter, r *http.Request, configuration Configuration, dictionary map[string]string) {
	text1, text2 := captions.GetJoke(configuration.CaptionSourceURL)
	textAbove, textBelow := translation.Translate(text1, dictionary), translation.Translate(text2, dictionary)
	fontPath := "./assets/fonts/impact.ttf"
	if bgImage, err := pictures.GetImage(configuration.ImageSourceURL, configuration.DefaultImageLocalPath); err != nil {
		web.GetErrorResponse(w, err)
	} else {
		if img, err := pictures.ApplyCaption(bgImage, textAbove, textBelow, fontPath); err != nil {
			// TODO at this point the background image is loaded successfully, maybe at least show that?
			web.GetErrorResponse(w, err)
		} else {
			web.WriteTemplateWithImage(w, &img)
		}
	}
}

func main() {
	var environment string
	flag.StringVar(&environment, "env", "dev", "Environment type, default is dev")
	flag.Parse()

	configuration := Configuration{}
	err := gonfig.GetConf(fmt.Sprintf("./%s_config.json", environment), &configuration)
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