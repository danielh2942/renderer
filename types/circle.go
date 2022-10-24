package types

import (
	"math"
)

const (
	// Values here were obtained from: https://spencermortensen.com/articles/types.circle/
	a float64 = 1.00005519
	b float64 = 0.55342686
	c float64 = 0.99873585
)

// Circle is a 2D Circle, it stores the Center and the Radius
type Circle struct {
	Center Vector2d
	Radius float64
}

// Translate moves the circle
func (cir *Circle) Translate(vec Vector2d) {
	cir.Center.Translate(vec)
}

// Scale scales the radius
func (cir *Circle) Scale(scaleFactor float64) {
	cir.Radius *= scaleFactor
}

// ScaleX does nothing
func (cir *Circle) ScaleX(float64) {}

// ScaleY does nothing
func (cir *Circle) ScaleY(float64) {}

// ScaleXY does nothing
func (cir *Circle) ScaleXY(float64, float64) {}

// RotateAbout rotates about an arbitrary point
func (cir *Circle) RotateAbout(vec Vector2d, angleRads float64) {
	cir.Center.RotateAbout(vec, angleRads)
}

// RotateAboutCenter does nothing
func (cir *Circle) RotateAboutCenter(float64) {}

// Render returns a slice containing the points on the perimeter of the circle
// It's mildly complex with a circle, check the link in the code for how it works
func (cir *Circle) Render() ([]Vector2d, error) {
	// semi arc approx p0 = (0,a), p1 = (b,c), p2 = (c,b), p3(a,0)
	radA := cir.Radius * a
	radB := cir.Radius * b
	radC := cir.Radius * c

	// Calculate the perimeter of the circle to get the minimum required points for this
	// 500 is a magic number that just works (so far anyway)
	// TODO: calculate some scaleable value to cover missing points
	perim := uint(math.Ceil(2*math.Pi*cir.Radius) + 500)

	// NB: Read this to understand how the rendering works
	// https://spencermortensen.com/articles/types.circle/

	tTop := Vector2d{
		X: cir.Center.X,
		Y: cir.Center.Y - radA,
	}

	tTopRightP1 := Vector2d{
		X: cir.Center.X + radB,
		Y: cir.Center.Y - radC,
	}

	tTopRightP2 := Vector2d{
		X: cir.Center.X + radC,
		Y: cir.Center.Y - radB,
	}

	tRight := Vector2d{
		X: cir.Center.X + radA,
		Y: cir.Center.Y,
	}

	tBottomRightP1 := Vector2d{
		X: cir.Center.X + radB,
		Y: cir.Center.Y + radC,
	}

	tBottomRightP2 := Vector2d{
		X: cir.Center.X + radC,
		Y: cir.Center.Y + radB,
	}

	tBottom := Vector2d{
		X: cir.Center.X,
		Y: cir.Center.Y + radA,
	}

	tBottomLeftP1 := Vector2d{
		X: cir.Center.X - radB,
		Y: cir.Center.Y + radC,
	}

	tBottomLeftP2 := Vector2d{
		X: cir.Center.X - radC,
		Y: cir.Center.Y + radB,
	}

	tLeft := Vector2d{
		X: cir.Center.X - radA,
		Y: cir.Center.Y,
	}

	tTopLeftP1 := Vector2d{
		X: cir.Center.X - radB,
		Y: cir.Center.Y - radC,
	}

	tTopLeftP2 := Vector2d{
		X: cir.Center.X - radC,
		Y: cir.Center.Y - radB,
	}

	mArr, _ := DrawCurve(perim/4, tTop, tTopRightP1, tTopRightP2, tRight)
	x, _ := DrawCurve(perim/4, tRight, tBottomRightP2, tBottomRightP1, tBottom)
	mArr = append(mArr, x...)
	x, _ = DrawCurve(perim/4, tBottom, tBottomLeftP1, tBottomLeftP2, tLeft)
	mArr = append(mArr, x...)
	x, _ = DrawCurve(perim/4, tLeft, tTopLeftP2, tTopLeftP1, tTop)
	mArr = append(mArr, x...)

	return mArr, nil
}
