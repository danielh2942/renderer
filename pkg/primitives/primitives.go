package primitives

import (
	"errors"
	"math"
)

type Obj2D interface {
	Translate(Vector2d)            // Move Obj2D by magnitude provided by vector passed
	Scale(float64)                 // Scale X,Y of all Points by Z magnitude
	ScaleX(float64)                // Scale X of all Points by Z magnitude
	ScaleY(float64)                // Scale Y of all Points by Z magnitude
	ScaleXY(float64, float64)      // Scale X and Y of all points by Z and A respectively
	RotateAbout(Vector2d, float64) // Rotate all points around given point by X coords
}

type Renderable interface {
	Render() ([]Vector2d, error)
}

// Helper functions
// GetVector2dMinMax returns the minimum and maximum X,Y values shared across a group of vectors
func GetVector2dMinMax(inpVec []Vector2d) ([2]Vector2d, error) {
	outp := [2]Vector2d{{}, {}}
	if len(inpVec) < 1 {
		return outp, errors.New("need at least one vector in list")
	}
	outp[0].X = inpVec[0].X
	outp[0].Y = inpVec[0].X
	outp[1].X = outp[0].X
	outp[1].X = outp[0].Y

	for _, vec := range inpVec {
		outp[0].X = math.Min(outp[0].X, vec.X)
		outp[1].X = math.Max(outp[1].X, vec.X)
		outp[0].Y = math.Min(outp[0].Y, vec.Y)
		outp[1].Y = math.Max(outp[1].Y, vec.Y)
	}

	return outp, nil
}
