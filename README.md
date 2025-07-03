# 🖥️ GoQRTerm

**Generate and display QR codes directly in your terminal with ease!**

---

## 📦 Installation
Get the module:
```bash
$ go get github.com/pedroborgesdev/goqrterm
```

Import the module into your Go project:

```go
package goqrterm

import (
	"github.com/pedroborgesdev/goqrterm/qr"
)
```

---

## 🚀 Usage

Simply call the Print function with the content you want to encode:

```go
func main() {
	goqrterm.Print("this is a simple test!")
}
```

### New features added:

* `goqrterm.Image("file_path")`

  Generates a QR code from the given string and saves the image as `./assets/qrcode.png`.

* `goqrterm.Decode("image_path")`

  Decodes the QR code from the given image path and returns a string with the decoded content.

Example:

```go
func main() {
	err := goqrterm.Image("Hello from GoQRTerm!")
	if err != nil {
		panic(err)
	}
	println("QR code image saved to ./assets/qrcode.png")

	decoded, err := goqrterm.Decode("./assets/qrcode.png")
	if err != nil {
		panic(err)
	}
	println("Decoded from QR Code:", decoded)
}
```

---

## 🖼️ Example Output

When executed, your terminal will display the following QR code:

<p align="center">
  <img src="./assets/example.png" alt="QR Code Example" width="800">
</p>

---

## 📱 Scan the QR Code

You can easily scan the displayed QR code with your smartphone to view the content:

<p align="center">
  <img src="./assets/camera.jpg" alt="Scanning QR Code" width="800">
</p>

---

## 💡 Features

✅ Simple to use
✅ Pure Go
✅ Terminal-friendly QR code display
✅ Generate QR code images (`Image`)
✅ Decode QR codes from images (`Decode`)

---

## 📝 License

Distributed under the MIT License. See LICENSE for more information.
