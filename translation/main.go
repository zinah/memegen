package translation

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"
)

func Translate(text string, dictionary map[string]string) (string) {
	// TODO handle punctuation, especially at the end
	// TODO handle empty text?
	tokens := strings.Split(text, " ")
	// TODO Should I use make with len and cap of tokens here instead and fill it, not append to it?
	var translatedText []string
	for _, word := range tokens {
		translatedWord := dictionary[strings.ToLower(word)]
		if translatedWord != "" {
			translatedText = append(translatedText, translatedWord)
		} else {
			translatedText = append(translatedText, word)
		}
	}
	return strings.Join(translatedText, " ")
}
	
func GetDictionaryFromJson(path string) (map[string]string) {
	jsonDictionary, err := ioutil.ReadFile(path)
    if err != nil {
        log.Fatal("Error when opening file: ", err)
	}

	var dictionary map[string]string
	// TODO handle errors when JSON malformed
	json.Unmarshal([]byte(jsonDictionary), &dictionary)

	return dictionary
}