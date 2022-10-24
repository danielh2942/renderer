package rend2img

import (
	"image"
	"image/color"
	"math"
	"unsafe"

	"github.com/danielh2942/renderer/types"
)

/*
I might need this in the future, saving.
func blendColors(col color.Color, col2 color.Color) color.RGBA {
	r1, g1, b1, a1 := col.RGBA()
	r2, g2, b2, a2 := col2.RGBA()

	aNew := uint8(255 - (((255 - a1) * (255 - a2)) / 255))
	rNew := uint8(((65535 - r1) + (r2 * a2)) / 255)
	gNew := uint8(((65535 - g1) + (g2 * a2)) / 255)
	bNew := uint8(((65535 - b1) + (b2 * a2)) / 255)
	return color.RGBA{R: rNew, G: gNew, B: bNew, A: aNew}
}
*/

// DrawPoints draws the points provided as an image.
func DrawPoints(inp []types.Vector2d, col [4]byte) *image.RGBA {
	minMax, _ := types.GetVector2dMinMax(inp)
	minMax[1] = minMax[1].GetRelativeCoords(minMax[0])

	// Create minimum image required to store the shape
	img := image.NewRGBA(image.Rect(0, 0, int(minMax[1].X)+1, int(minMax[1].Y)+1))

	for _, pxl := range inp {
		// Adjust the points to be explicitly within the bounds of the image
		tmp := pxl.GetRelativeCoords(minMax[0])

		img.SetRGBA(int(math.Round(tmp.X)), int(math.Round(tmp.Y)), color.RGBA{
			R: col[0],
			G: col[1],
			B: col[2],
			A: col[3],
		})
	}
	return img
}

// ScanLineFill assumes X and Y are valid and does some kind of messy pointer manipulations
// in order to reduce the amount of work required to fill the image.
func ScanLineFill(img *image.RGBA, startX int, endX int, y int, col [4]byte) {
	imgPtr := (*[]uint32)(unsafe.Pointer(&img.Pix))
	colVal := (*uint32)(unsafe.Pointer(&col[0]))

	startIdx := img.Rect.Max.X*y + startX

	copyExpLimit := int(math.Pow(2, math.Floor(math.Log2(float64(endX-startX)))))
	(*imgPtr)[startIdx] = *colVal
	for x := 1; x < copyExpLimit; x *= 2 {
		copy((*imgPtr)[startIdx+x:], (*imgPtr)[startIdx:startIdx+x])
	}
	copy((*imgPtr)[startIdx+copyExpLimit:], (*imgPtr)[startIdx:(startIdx+(endX-startX))-copyExpLimit])
}

// FillShape uses even-odd rule to fill a space within colored points.
// It screws up on shapes with insets, but technically they're no longer
// One shape so that is expected behavior.
func FillShape(img *image.RGBA, col [4]byte) {
	bounds := img.Bounds()
	imgMaxX := bounds.Max.X
	imgMaxY := bounds.Max.Y
	for y := 1; y < imgMaxY-1; y++ {
		intersections := 0
		prev := [2]int{0, 0}
		for x := 0; x < imgMaxX; x++ {
			// Check for ANY color whatsoever
			if _, _, _, a := img.At(x, y).RGBA(); a != 0 {
				intersections++
				// Look until there is no color
				for _, _, _, a := img.At(x, y).RGBA(); a != 0 && x < imgMaxX; _, _, _, a = img.At(x, y).RGBA() {
					x++
				}
				x--

				// If Odd: Entering shape, else: Leaving
				if intersections%2 == 1 {
					prev[0] = x
					prev[1] = y
				} else {
					// Draw a scanline
					ScanLineFill(img, prev[0]+1, x, y, col)
				}
			}
		}
	}
}

// DrawCompositePoints takes a composite shape and draws it.
func DrawCompositePoints(inpVec [][]types.Vector2d, col [4]byte) (*image.RGBA, error) {
	minMax, err := types.GetCompositeVector2dMinMax(inpVec)
	if err != nil {
		return nil, err
	}

	minMax[1] = minMax[1].GetRelativeCoords(minMax[0])

	img := image.NewRGBA(image.Rect(0, 0, int(minMax[1].X)+1, int(minMax[1].Y)+1))

	for _, arr := range inpVec {
		for _, pt := range arr {
			pt = pt.GetRelativeCoords(minMax[0])
			img.SetRGBA(int(math.Round(pt.X)), int(math.Round(pt.Y)), color.RGBA{
				R: col[0],
				G: col[1],
				B: col[2],
				A: col[3],
			})
		}
	}

	return img, nil
}
