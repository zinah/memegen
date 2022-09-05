package pictures

import (
	"image"
	"github.com/fogleman/gg"
)

func ApplyText(ctx *gg.Context, textStartPosX float64, textStartPosY float64, text string, maxWidth float64, strokeSize int) (*gg.Context){
	return ctx
}

func ApplyCaption(bgImage image.Image, textAbove string, textBelow string) (image.Image, error) {
	return bgImage, nil
}