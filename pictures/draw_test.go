package pictures

import (
	"testing"
	"github.com/corona10/goimagehash"
	"github.com/fogleman/gg"
)

func TestApplyCaption(t *testing.T) {
	t1, t2 := "y doez python programmers wear glasses?", "cuz they can't C"
	testImg, _ := gg.LoadImage("./fixtures/default_test_cat.png")
	expected, _ := gg.LoadImage("./fixtures/meme_cat.jpeg")
	meme, err := ApplyCaption(testImg, t1, t2, "./fixtures/fonts/impact.ttf")

	imgHash, _ := goimagehash.AverageHash(meme)
	expectedHash, _ := goimagehash.AverageHash(expected)
	d, _ := imgHash.Distance(expectedHash)
	if d != 0 {
		t.Errorf("Image different than expected")
	}
	if err != nil {
		t.Fatalf("Got error: %s", err)
	}
}