package pictures

import (
	"image"
	"image/color"
	"github.com/fogleman/gg"
)

func ApplyText(ctx *gg.Context, textStartPosX float64, textStartPosY float64, text string, maxWidth float64, strokeSize int) (*gg.Context){
	ctx.SetColor(color.Black)
	for dy := -strokeSize; dy <= strokeSize; dy++ {
		for dx := -strokeSize; dx <= strokeSize; dx++ {
			if dx * dx + dy * dy >= strokeSize * strokeSize {
				// give it rounded corners
				continue
			}
			x := textStartPosX + float64(dx)
			y := textStartPosY + float64(dy)
			ctx.DrawStringWrapped(text, x, y, 0.5, 0.5, maxWidth, 1.5, gg.AlignCenter)
		}
	}

	ctx.SetColor(color.White)
	ctx.DrawStringWrapped(text, textStartPosX, textStartPosY, 0.5, 0.5, maxWidth, 1.5, gg.AlignCenter)

	return ctx
}

func ApplyCaption(bgImage image.Image, textAbove string, textBelow string) (image.Image, error) {
	imgWidth := bgImage.Bounds().Dx()
	imgHeight := bgImage.Bounds().Dy()

	ctx := gg.NewContext(imgWidth, imgHeight)
	ctx.DrawImage(bgImage, 0, 0)

	// TODO calculate appropriate font size based on the image size
	if err := ctx.LoadFontFace("./assets/fonts/impact.ttf", 30); err != nil {
		return nil, err
	}

	// TODO better way to find optimal placement of the text so it doesn't cover important parts of the image
	textAboveStartPosX := float64(imgWidth / 2)
	textAboveStartPosY := float64(100)
	textBelowStartPosX := float64(imgWidth / 2)
	textBelowStartPosY := float64(imgHeight - 120)

	maxWidth := float64(imgWidth) - 60.0
	strokeSize := 6

	ApplyText(ctx, textAboveStartPosX, textAboveStartPosY, textAbove, maxWidth, strokeSize)
	ApplyText(ctx, textBelowStartPosX, textBelowStartPosY, textBelow, maxWidth, strokeSize)

	return ctx.Image(), nil
}