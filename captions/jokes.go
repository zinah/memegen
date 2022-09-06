package captions

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func GetJoke(CaptionSourceURL string) (string, string) {
	textAboveDefault, textBelowDefault := "How do you know God is a shitty programmer?", "He wrote the OS for an entire universe, but didn't leave a single useful comment."
	client := http.Client{
		Timeout: 1 * time.Second,
	}
	resp, err := client.Get(CaptionSourceURL)
    if err != nil {
        log.Println(err)
		return textAboveDefault, textBelowDefault
    }
    defer resp.Body.Close()
	// TODO handle errors
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return textAboveDefault, textBelowDefault
	}

	var jokeRes map[string]string
	// TODO handle errors when JSON malformed
	json.Unmarshal([]byte(body), &jokeRes)
	// TODO use jokesRes["type"] ("twopart"/"single") instead
	setup, ok := jokeRes["setup"]
	delivery, ok := jokeRes["delivery"]
	if ok {
		return setup, delivery
	} else {
		if joke, ok := jokeRes["joke"]; ok {
			return "", joke
		}
	}

	return textAboveDefault, textBelowDefault
}