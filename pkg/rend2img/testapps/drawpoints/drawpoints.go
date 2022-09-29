package main

import (
	"fmt"
	"image/png"
	"os"

	"github.com/danielh2942/renderer/pkg/primitives"
	"github.com/danielh2942/renderer/pkg/rend2img"
	"github.com/danielh2942/renderer/pkg/types"
)

func main() {
	q1 := types.Quad{
		Point1: primitives.Vector2d{X: 25, Y: 0},
		Point2: primitives.Vector2d{X: 75, Y: 0},
		Point3: primitives.Vector2d{X: 0, Y: 100},
		Point4: primitives.Vector2d{X: 100, Y: 100},
	}

	imgPts, _ := q1.Render()
	mImg := rend2img.DrawPoints(imgPts, [4]byte{0xFF, 0xFF, 0xFF, 0xFF})
	f, err := os.Create("output.png")
	if err != nil {
		fmt.Println("ERROR", err)
		return
	}

	if err = png.Encode(f, mImg); err != nil {
		fmt.Println("ERROR", err)
	}
	f.Close()

	// Filled Quad time
	f, err = os.Create("output_filled.png")
	if err != nil {
		fmt.Println("ERROR", err)
		return
	}

	rend2img.FillShape(mImg, [4]byte{0xFF, 0x00, 0x00, 0xFF})

	if err = png.Encode(f, mImg); err != nil {
		fmt.Println("ERROR", err)
	}

	f.Close()

	nPol := types.NPoly{
		Points: []primitives.Vector2d{
			{X: 500, Y: 0},
			{X: 0, Y: 200},
			{X: 150, Y: 500},
			{X: 850, Y: 500},
			{X: 1000, Y: 200},
		},
	}

	imgPts, _ = nPol.Render()
	mImg = rend2img.DrawPoints(imgPts, [4]byte{0xFF, 0x00, 0xFF, 0xFF})

	f, err = os.Create("output_npoly.png")
	if err != nil {
		fmt.Println("ERROR", err)
		return
	}

	if err = png.Encode(f, mImg); err != nil {
		fmt.Println("ERROR", err)
	}
	f.Close()

	rend2img.FillShape(mImg, [4]byte{0xFF, 0x00, 0xFF, 0xFF})

	f, err = os.Create("output_npoly_filled.png")
	if err != nil {
		fmt.Println("ERROR", err)
		return
	}

	if err = png.Encode(f, mImg); err != nil {
		fmt.Println("ERROR", err)
	}
	f.Close()
	
	star := types.NPoly{
		Points: []primitives.Vector2d{
			{X:300,Y:0},
			{X:388,Y:179},
			{X:585,Y:207},
			{X:442,Y:346},
			{X:476,Y:507},
			{X:300,Y:450},
			{X:124,Y:507},
			{X:158,Y:346},
			{X:15,Y:207},
			{X:222,Y:179},
		},
	}

	imgPts, _ = star.Render()
	mImg = rend2img.DrawPoints(imgPts, [4]byte{0x00,0xFF,0x00,0xFF})

	f, err = os.Create("star.png")
	if err != nil {
		fmt.Println("ERROR",err)
		return
	}

	if err = png.Encode(f,mImg); err != nil {
		fmt.Println("ERROR",err)
	}
	f.Close()

	rend2img.FillShape(mImg,[4]byte{0x00,0xFF,0x00,0xFF})

	f, err = os.Create("star_filled.png")
	if err != nil {
		fmt.Println("ERROR",err)
		return
	}

	if err = png.Encode(f,mImg); err != nil {
		fmt.Println("ERROR", err)
	}
	f.Close()

	circle := types.Circle{
		Center:primitives.Vector2d{X:300,Y:300},
		Radius:300,
	}
	
	imgPts, _ = circle.Render()
	mImg = rend2img.DrawPoints(imgPts,[4]byte{0x00,0x00,0xFF,0xFF})

	f, err = os.Create("Circle.png")
	if err != nil {
		fmt.Println("ERROR",err)
		return
	}

	if err = png.Encode(f, mImg);err != nil {
		fmt.Println("ERROR",err)
	}
	f.Close()

	rend2img.FillShape(mImg,[4]byte{0x00,0x00,0xFF,0xFF})

	f, err = os.Create("Circle_filled.png")
	if err!= nil {
		fmt.Println("ERROR",err)
		return
	}

	if err = png.Encode(f,mImg); err != nil {
		fmt.Println("ERROR",err)
	}
	f.Close()
}
