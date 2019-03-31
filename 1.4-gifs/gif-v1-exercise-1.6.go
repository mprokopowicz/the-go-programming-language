package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{
	color.RGBA{0, 0, 0, 1},
	color.RGBA{0, 0, 100, 1},
	color.RGBA{0, 0, 120, 1},
	color.RGBA{0, 0, 140, 1},
	color.RGBA{0, 0, 160, 1},
	color.RGBA{0, 0, 180, 1},
}

const (
	primaryColorIndex   = 0
	secondaryColorIndex = 1
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)

	var nextColorIndex uint8 = 1
	freq := rand.Float64() * 1.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i += 1 {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		nextColorIndex++
		nextColorIndex = uint8(nextColorIndex%4 + 1)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)

			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), nextColorIndex)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim)
}
