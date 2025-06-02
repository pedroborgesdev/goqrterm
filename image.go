package main

import (
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
)

func Image(text string, filenames ...string) error {
	filename := "qrcode.png"
	if len(filenames) > 0 {
		filename = filenames[0]
	}

	writer := qrcode.NewQRCodeWriter()
	hints := map[gozxing.EncodeHintType]interface{}{
		gozxing.EncodeHintType_ERROR_CORRECTION: "M",
		gozxing.EncodeHintType_MARGIN:           1,
	}

	bitMatrix, err := writer.Encode(text, gozxing.BarcodeFormat_QR_CODE, 0, 0, hints)
	if err != nil {
		return err
	}

	scale := 20
	width := bitMatrix.GetWidth()
	height := bitMatrix.GetHeight()
	img := image.NewRGBA(image.Rect(0, 0, width*scale, height*scale))

	black := color.RGBA{0, 0, 0, 255}
	white := color.RGBA{255, 255, 255, 255}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			var c color.Color
			if bitMatrix.Get(x, y) {
				c = black
			} else {
				c = white
			}
			for dy := 0; dy < scale; dy++ {
				for dx := 0; dx < scale; dx++ {
					img.Set(x*scale+dx, y*scale+dy, c)
				}
			}
		}
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		return err
	}

	return nil
}
