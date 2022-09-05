package pictures

import (
	"image"
	"github.com/fogleman/gg"
	"log"
	"net/http"
	"time"
)

func GetImageStatic(path string) (image.Image, error) {
	bgImage, err := gg.LoadImage(path)
	if err != nil {
		return nil, err
	}
	return bgImage, nil
}

func GetImageHttp(url string) (image.Image, error) {
	// TODO timeout should be a config option
	client := http.Client{
		Timeout: 1 * time.Second,
	}
	resp, err := client.Get(url)
    if err != nil {
        log.Println(err)
		return nil, err
    }
    defer resp.Body.Close()
	bgImage, _, _ := image.Decode(resp.Body)
	if err != nil {
		return nil, err
	}
	return bgImage, nil
}

func GetImage(imageSourceURL string, defaultImageLocalPath string) (image.Image, error) {
	img, err := GetImageHttp(imageSourceURL)
	// TODO load the default picture only once and pass it into this function
	// instead of loading every time something goes wrong?
	// Maybe no point in loading it anticipating failure
	// At least it should be cached somehow
	if err != nil {
		log.Println("Getting default image")
		img, err := GetImageStatic(defaultImageLocalPath)
		if err != nil {
			return nil ,err
		}

		return img, nil
	}

	return img, nil
}