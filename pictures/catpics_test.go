package pictures

import (
	"bytes"
	"image"
	"net/http"
	"testing"
	"github.com/corona10/goimagehash"
	"github.com/jarcoal/httpmock"
)

func TestGetImageStaticCorrectPath(t *testing.T) {
	img, err := GetImageStatic("./fixtures/meme_cat.jpeg")
	if img != nil {
		t.Logf("Success !")
	}
	if err == nil {
		t.Logf("Success !")
	}
}

func TestGetImageStaticInorrectPath(t *testing.T) {
	img, err := GetImageStatic("./fixtures/non_existent.jpeg")
	if img == nil {
		t.Logf("Success !")
	}
	if err != nil {
		t.Logf("Success !")
	}
}

func TestGetImageHttp(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		"GET", "https://cataas.com/cat",
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewBytesResponse(200, httpmock.File("./fixtures/meme_cat.jpeg").Bytes()), nil
		})
	img, err := GetImageHttp("https://cataas.com/cat")

	expected, _, _ := image.Decode(bytes.NewReader(httpmock.File("./fixtures/meme_cat.jpeg").Bytes()))
	imgHash, _ := goimagehash.AverageHash(img)
	expectedHash, _ := goimagehash.AverageHash(expected)
	d, _ := imgHash.Distance(expectedHash)
	if d != 0 {
		t.Errorf("Image different than expected")
	}
	if err != nil {
		t.Fatalf("Got error: %s", err)
	}
}

func TestGetImageHttpNoImage(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://cataas.com/cat", httpmock.NewStringResponder(500, `error`))
	img, err := GetImageHttp("https://cataas.com/cat")
	if img != nil {
		t.Errorf("There should be no image returned")
	}
	if err == nil {
		t.Fatalf("No error raised")
	}
}

// TODO test timeout

func TestGetImage(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		"GET", "https://cataas.com/cat",
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewBytesResponse(200, httpmock.File("./fixtures/meme_cat.jpeg").Bytes()), nil
		})
	img, err := GetImage("https://cataas.com/cat", "./fixtures/default_cat.jpeg")

	expected, _, _ := image.Decode(bytes.NewReader(httpmock.File("./fixtures/meme_cat.jpeg").Bytes()))
	imgHash, _ := goimagehash.AverageHash(img)
	expectedHash, _ := goimagehash.AverageHash(expected)
	d, _ := imgHash.Distance(expectedHash)
	if d != 0 {
		t.Errorf("Image different than expected")
	}
	if err != nil {
		t.Fatalf("Got error: %s", err)
	}
}

func TestGetImageNoImage(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://cataas.com/cat", httpmock.NewStringResponder(500, `error`))
	img, err := GetImage("https://cataas.com/cat", "./fixtures/default_cat.jpeg")

	expected, _, _ := image.Decode(bytes.NewReader(httpmock.File("./fixtures/default_cat.jpeg").Bytes()))
	expectedHash, _ := goimagehash.AverageHash(expected)
	imgHash, _ := goimagehash.AverageHash(img)
	d, _ := imgHash.Distance(expectedHash)
	if d != 0 {
		t.Errorf("Image different than expected")
	}
	if err != nil {
		t.Fatalf("Got error: %s", err)
	}
}