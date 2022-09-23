package primitives

import (
	"math"
)

// Vector2d is a 2d Vector storing doubles
type Vector2d struct {
	X float64 `json:"X"` // X Coordinate
	Y float64 `json:"Y"` // Y Coordinate
}

// Scale does what you think it does.
// it also deliberately allows for negative values so you can invert it this way :)
func (v2d *Vector2d) Scale(scaleFactor float64) {
	v2d.X *= scaleFactor
	v2d.Y *= scaleFactor
}

// ScaleX scales the X by a given value
func (v2d *Vector2d) ScaleX(scaleFactor float64) {
	v2d.X *= scaleFactor
}

// ScaleY scales the Y by a given value
func (v2d *Vector2d) ScaleY(scaleFactor float64) {
	v2d.Y *= scaleFactor
}

// ScaleXY allows you to scale X and Y by different vals in one function
func (v2d *Vector2d) ScaleXY(scaleFactorX float64, scaleFactorY float64) {
	v2d.X *= scaleFactorX
	v2d.Y *= scaleFactorY
}

// TranslateAndScale moves the Vector by X,Y
func (v2d *Vector2d) Translate(translate Vector2d) {
	v2d.X = v2d.X + translate.X
	v2d.Y = v2d.Y + translate.Y
}

func (v2d *Vector2d) GetDotProduct(vector Vector2d) float64 {
	return (v2d.X * vector.X) + (v2d.Y * vector.Y)
}

// GetRelativeCoords gets the relative coordinates of a point from another given point
func (v2d *Vector2d) GetRelativeCoords(point Vector2d) Vector2d {
	vect := Vector2d{v2d.X, v2d.Y}
	vect.X -= point.X
	vect.Y -= point.Y
	return vect
}

// RotateAbout rotates a point around an arbitrary point by some amount of radians
func (v2d *Vector2d) RotateAbout(point Vector2d, angleRads float64) {
	// Treat the point as the origin
	temp := v2d.GetRelativeCoords(point)
	// Rotate around Origin
	v2d.X = (math.Cos(angleRads) * temp.X) + (math.Sin(angleRads) * -temp.Y)
	v2d.Y = (math.Sin(angleRads) * temp.X) + (math.Cos(angleRads) * temp.Y)
	// Translate back to where it should be
	v2d.Translate(point)
}
