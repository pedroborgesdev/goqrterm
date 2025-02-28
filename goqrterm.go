package goqrterm

import (
	"fmt"
	"io"

	"github.com/qpliu/qrencode-go/qrencode"
)

func PrintCompactQR(w io.Writer, grid *qrencode.BitGrid) {
	black := "█"
	white := " "
	borderLeft := "█ "
	borderRight := " █"
	borderTop := "▀"

	height := grid.Height()
	width := grid.Width()

	fmt.Println("\n▌ GoQrShow is working now!")
	fmt.Println("▌ Follow GitHub Project: https://github.com/pedroborgesdev/GoQrShow.git")
	fmt.Println("▌ Created by pedroborgesdev\n ")

	fmt.Fprint(w, black)
	for x := 0; x < width; x++ {
		fmt.Fprint(w, borderTop)
	}
	fmt.Fprintln(w, (borderTop + borderTop + black))

	for y := 0; y < height; y += 2 {
		fmt.Fprint(w, borderLeft)
		for x := 0; x < width; x++ {
			upper := grid.Get(x, y)
			lower := false
			if y+1 < height {
				lower = grid.Get(x, y+1)
			}

			upperBlack := upper
			lowerBlack := lower

			var char string
			switch {
			case upperBlack && lowerBlack:
				char = black
			case upperBlack && !lowerBlack:
				char = "▀"
			case !upperBlack && lowerBlack:
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

	fmt.Println("▀ QRcode has been showed up!")
	fmt.Println("▀ You can scan now!")
}
