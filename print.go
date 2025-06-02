package goqrterm

import (
	"fmt"
	"os"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
)

func Print(text string) error {
	w := os.Stdout

	writer := qrcode.NewQRCodeWriter()

	hints := map[gozxing.EncodeHintType]interface{}{
		gozxing.EncodeHintType_MARGIN: 0,
	}

	bitMatrix, err := writer.Encode(text, gozxing.BarcodeFormat_QR_CODE, 0, 0, hints)
	if err != nil {
		return err
	}

	black := "█"
	white := " "
	borderLeft := "█ "
	borderRight := " █"
	borderTop := "▀"

	height := bitMatrix.GetHeight()
	width := bitMatrix.GetWidth()

	fmt.Fprint(w, black)
	for x := 0; x < width; x++ {
		fmt.Fprint(w, borderTop)
	}
	fmt.Fprintln(w, (borderTop + borderTop + black))

	for y := 0; y < height; y += 2 {
		fmt.Fprint(w, borderLeft)
		for x := 0; x < width; x++ {
			upper := bitMatrix.Get(x, y)
			lower := false
			if y+1 < height {
				lower = bitMatrix.Get(x, y+1)
			}

			var char string
			switch {
			case upper && lower:
				char = black
			case upper && !lower:
				char = "▀"
			case !upper && lower:
				char = "▄"
			default:
				char = white
			}
			fmt.Fprint(w, char)
		}
		fmt.Fprintln(w, borderRight)
	}

	fmt.Fprint(w, borderTop)
	for x := 0; x < width; x++ {
		fmt.Fprint(w, borderTop)
	}
	fmt.Fprintln(w, (borderTop + borderTop + borderTop))

	return nil
}
