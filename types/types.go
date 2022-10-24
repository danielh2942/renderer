package types

// types is a package containing all the base types known to the renderer library

import (
	"errors"
	"math"
)

// Obj2D is an interface defining the API for any 2D object that can be handled by the renderer
type Obj2D interface {
	Translate(Vector2d)            // Move Obj2D by magnitude provided by vector passed
	Scale(float64)                 // Scale X,Y of all Points by Z magnitude
	ScaleX(float64)                // Scale X of all Points by Z magnitude
	ScaleY(float64)                // Scale Y of all Points by Z magnitude
	ScaleXY(float64, float64)      // Scale X and Y of all points by Z and A respectively
	RotateAbout(Vector2d, float64) // Rotate all points around given point by X coords
	RotateAboutCenter(float64)     // Rotate a 2D Object about its center
}

type Renderable interface {
	Render() ([]Vector2d, error)
}

// RenderableComposite constructs an Obj2D image as a more complicated shape
type RenderableComposite interface {
	RenderComposite() ([][]Vector2d, error)
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

// GetCompositeVector2dMinMax does the same but for 2d arrays
func GetCompositeVector2dMinMax(inpVec [][]Vector2d) ([2]Vector2d, error) {
	outp := [2]Vector2d{}

	if len(inpVec) < 1 {
		return outp, errors.New("need at least one slice of vectors in slice")
	}

	for _, p := range inpVec {
		tMinMax, err := GetVector2dMinMax(p)
		if err != nil {
			return outp, err
		}

		outp[0].X = math.Min(tMinMax[0].X, outp[0].X)
		outp[0].Y = math.Min(tMinMax[0].Y, outp[0].Y)
		outp[1].X = math.Max(tMinMax[1].X, outp[1].X)
		outp[1].Y = math.Max(tMinMax[1].Y, outp[1].Y)
	}

	return outp, nil
}

// GetCosineSimilarity gets the cosine similarity between two vectors
// P1 is end 1
// P2 is the pivot point
// P3 is end 2
func GetCosineSimilarity(p1, p2, p3 Vector2d) float64 {
	tP1 := p1.GetRelativeCoords(p2)
	tP3 := p3.GetRelativeCoords(p2)

	return (tP1.GetDotProduct(tP3) / (tP1.GetAbs() * tP3.GetAbs()))
}
