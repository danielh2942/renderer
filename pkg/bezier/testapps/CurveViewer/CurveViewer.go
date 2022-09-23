package main

// CurveViewer.go
// Version 1
// 2022 - Daniel Hannon (danielh2942)

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"

	"github.com/danielh2942/renderer/pkg/bezier"
	"github.com/danielh2942/renderer/pkg/primitives"
)

// CurveViewer, this is for checking the the bezier curve function works

func main() {
	// We'll check a few different curves on the same image I suppose
	// I'll use my native resolution for lack of a better value I can think of
	img := image.NewRGBA(image.Rect(0, 0, 1680, 1050))
	for y := 0; y < 1050; y++ {
		for x := 0; x < 1680; x++ {
			img.SetRGBA(x, y, color.RGBA{0, 0, 0, 255})
		}
	}
	// Test a straight line because why not
	straightLine := []primitives.Vector2d{{X: 0, Y: 0}, {X: 100, Y: 100}}
	// it's 45 deg so it should be fine to use 101 points
	outStraightLine, _ := bezier.DrawCurve(101, straightLine...)
	for _, pt := range outStraightLine {
		img.SetRGBA(int(pt.X), int(pt.Y), color.RGBA{255, 255, 255, 255})
	}

	// Make an Arc (Three points)
	arc := []primitives.Vector2d{{X: 0, Y: 849}, {X: 0, Y: 1049}, {X: 200, Y: 1049}}
	outArc, _ := bezier.DrawCurve(400, arc...)

	for _, pt := range outArc {
		img.SetRGBA(int(pt.X), int(pt.Y), color.RGBA{255, 255, 255, 255})
		img.SetRGBA(int(math.Round(pt.X)), int(math.Round(pt.Y)), color.RGBA{255, 255, 255, 255})
	}

	// Four point Arc
	fourArc := []primitives.Vector2d{{X: 600, Y: 700}, {X: 850, Y: 520}, {X: 600, Y: 20}, {X: 1280, Y: 720}}
	outFourArc, _ := bezier.DrawCurve(2000, fourArc...)

	for _, pt := range outFourArc {
		img.SetRGBA(int(pt.X), int(pt.Y), color.RGBA{255, 255, 255, 255})
		img.SetRGBA(int(math.Round(pt.X)), int(math.Round(pt.Y)), color.RGBA{255, 255, 255, 255})
	}

	// Output image
	f, err := os.Create("test.png")
	if err != nil {
		return
	}
	defer f.Close()

	if err := png.Encode(f, img); err != nil {
		fmt.Println(err)
	}
}
