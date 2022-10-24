package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"

	"github.com/danielh2942/renderer/rend2img"
	"github.com/danielh2942/renderer/types"
)

func main() {
	star := types.NPoly{
		Points: []types.Vector2d{
			{X: 300, Y: 0},
			{X: 388, Y: 179},
			{X: 585, Y: 207},
			{X: 442, Y: 346},
			{X: 476, Y: 507},
			{X: 300, Y: 450},
			{X: 124, Y: 507},
			{X: 158, Y: 346},
			{X: 15, Y: 207},
			{X: 222, Y: 179},
		},
	}

	// Unfilled composite
	compPts, _ := star.RenderComposite()

	minMax, _ := types.GetCompositeVector2dMinMax(compPts)

	minMax[1] = minMax[1].GetRelativeCoords(minMax[0])

	img := image.NewRGBA(image.Rect(0, 0, int(minMax[1].X)+1, int(minMax[1].Y)+1))

	for _, arr := range compPts {
		for _, pt := range arr {
			img.SetRGBA(int(math.Round(pt.X)), int(math.Round(pt.Y)), color.RGBA{R: 0xFF, G: 0x00, B: 0x00, A: 0xFF})
		}
	}

	f, err := os.Create("composite_shape.png")
	if err != nil {
		fmt.Println("ERROR", err)
	}

	if err = png.Encode(f, img); err != nil {
		fmt.Println("ERROR", err)
	}
	f.Close()

	// Render out the boxes of the afforementioned shape

	// img = image.NewRGBA(image.Rect(0, 0, int(minMax[1].X)+1, int(minMax[1].Y)+1))

	triangles := star.GetTriangles()

	for _, t := range triangles {
		tMinMaxPts, _ := types.GetVector2dMinMax([]types.Vector2d{t.Point1, t.Point2, t.Point3})

		mQd := types.Quad{
			Point1: tMinMaxPts[0],
			Point4: tMinMaxPts[1],
			Point2: types.Vector2d{
				X: tMinMaxPts[1].X,
				Y: tMinMaxPts[0].Y,
			},
			Point3: types.Vector2d{
				X: tMinMaxPts[0].X,
				Y: tMinMaxPts[1].Y,
			},
		}

		for y := mQd.Point1.Y; y <= mQd.Point3.Y; y++ {
			intersections := 0
			var prev int
			for x := mQd.Point1.X; x <= mQd.Point2.X; x++ {
				if _, _, _, a := img.At(int(math.Round(x)), int(math.Round(y))).RGBA(); a != 0 {
					intersections++
					for _, _, _, a := img.At(int(math.Round(x)), int(math.Round(y))).RGBA(); a != 0; _, _, _, a = img.At(int(math.Round(x)), int(math.Round(y))).RGBA() {
						x++
					}
					x--

					if intersections%2 == 1 {
						prev = int(math.Round(x))
					} else {
						rend2img.ScanLineFill(img, prev+1, int(math.Round(x)), int(math.Round(y)), [4]byte{0xFF, 0x00, 0x00, 0xFF})
					}
				}
			}
		}

		/*
			mQdPts, _ := mQd.Render()

			for _, pt := range mQdPts {
				img.Set(int(math.Round(pt.X)), int(math.Round(pt.Y)), color.RGBA{
					R: 0xFF,
					G: 0xFF,
					B: 0xFF,
					A: 0xFF,
				})
			}
		*/
	}

	f, err = os.Create("composite_shape_filled.png")
	if err != nil {
		fmt.Println("ERROR", err)
	}

	if err = png.Encode(f, img); err != nil {
		fmt.Println("ERROR", err)
	}
	f.Close()
}
