package main

import (
	"github.com/issue9/identicon"
	"image/color"
	"image/png"
	"os"
)

func main() {
	img, _ := identicon.Make(256, color.NRGBA{255, 255, 0, 255}, color.NRGBA{0, 0, 0, 255}, []byte("my name is lihaoquan"))
	fi, _ := os.Create("/Users/lihaoquan/Desktop/u1.png")
	png.Encode(fi, img)
	fi.Close()
}
