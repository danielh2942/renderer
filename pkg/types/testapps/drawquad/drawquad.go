package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"unsafe"

	"github.com/danielh2942/renderer/pkg/primitives"
	"github.com/danielh2942/renderer/pkg/types"
)

func imgset(img *image.RGBA, myCol []byte) {
	// This is a memcpy like function for setting an entire image to some color
	imPtr := (*[]uint32)(unsafe.Pointer(&img.Pix))
	imCol := (*uint32)(unsafe.Pointer(&myCol[0]))

	// As the pointer has no upper limit I have to write provisions to prevent
	// Segfaults, looks kinda yucky but it works so idk
	imSize := img.Rect.Max.X * img.Rect.Max.Y
	copyExpLimit := int(math.Pow(2, math.Floor(math.Log2(float64(imSize)))))

	(*imPtr)[0] = *imCol
	for x := 1; x < copyExpLimit; x *= 2 {
		copy((*imPtr)[x:], (*imPtr)[:x])
	}
	copy((*imPtr)[copyExpLimit:], (*imPtr)[:(imSize-copyExpLimit)])
}

func main() {
	img := image.NewRGBA(image.Rect(0, 0, 2000, 2000))
	imgset(img, []byte{0, 0, 0, 255})

	f, err := os.Create("test.png")
	if err != nil {
		return
	}
	defer f.Close()

	mQuad := types.Quad{
		Point1: primitives.Vector2d{X: 500, Y: 500},
		Point2: primitives.Vector2d{X: 1000, Y: 500},
		Point3: primitives.Vector2d{X: 500, Y: 1000},
		Point4: primitives.Vector2d{X: 1000, Y: 1000},
	}

	pts, _ := mQuad.Render()

	for _, pt := range pts {
		img.SetRGBA(int(pt.X), int(pt.Y), color.RGBA{255, 255, 255, 255})
	}

	mQuad2 := types.Quad{
		Point1: primitives.Vector2d{X: 750, Y: 250},
		Point2: primitives.Vector2d{X: 1250, Y: 750},
		Point3: primitives.Vector2d{X: 750, Y: 1250},
		Point4: primitives.Vector2d{X: 250, Y: 750},
	}

	pts, _ = mQuad2.Render()
	for _, pt := range pts {
		img.SetRGBA(int(pt.X), int(pt.Y), color.RGBA{255, 255, 255, 255})
	}

	if err := png.Encode(f, img); err != nil {
		fmt.Println(err)
	}
}
