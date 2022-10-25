package types

import (
	//"encoding/json"
	"math"
)

// Vector3d is the primitive type for all 3D objects
type Vector3d struct {
	Vector2d
	Z float64
}

// ScaleXZ scales X and Z values by some factors
func (v3d *Vector3d) ScaleXZ(scaleX, scaleZ float64) {
	v3d.X *= scaleX
	v3d.Z *= scaleZ
}

// ScaleYZ scales Y and Z values by some factors
func (v3d *Vector3d) ScaleYZ(scaleY, scaleZ float64) {
	v3d.Y *= scaleY
	v3d.Z *= scaleZ
}

// ScaleXYZ scales X, Y, and Z values by some factor
func (v3d *Vector3d) ScaleXYZ(scaleX, scaleY, scaleZ float64) {
	v3d.X *= scaleX
	v3d.Y *= scaleY
	v3d.Z *= scaleZ
}

// Scale3D scales X,Y,and Z by some factor
func (v3d *Vector3d) Scale3D(scaleFactor float64) {
	v3d.Scale(scaleFactor)
	v3d.Z *= scaleFactor
}

// Translate3D moves the vector by some point
func (v3d *Vector3d) Translate3D(vector Vector3d) {
	v3d.Translate(vector.Vector2d)
	v3d.Z += vector.Z
}

// GetRelativeCoords3D gets the relative coordinates to some point in 3d space
func (v3d *Vector3d) GetRelativeCoords3D(point Vector3d) Vector3d {
	return Vector3d{
		Vector2d: v3d.GetRelativeCoords(point.Vector2d),
		Z:        v3d.Z - point.Z,
	}
}

// RotateAboutX rotates the vector about X by some rads
func (v3d *Vector3d) RotateAboutX(degRads float64) {
	tCos := math.Cos(degRads)
	tSin := math.Sin(degRads)

	tempY := v3d.Y
	tempZ := v3d.Z

	v3d.Y = (tCos * tempY) - (tSin * tempZ)
	v3d.Z = (tSin * tempY) - (tCos * tempZ)
}

// RotateAboutY rotates about Y by some rads
func (v3d *Vector3d) RotateAboutY(degRads float64) {
	tCos := math.Cos(degRads)
	tSin := math.Sin(degRads)

	tempX := v3d.X
	tempZ := v3d.Z

	v3d.X = (tCos * tempX) + (tSin * tempZ)
	v3d.Z = (-tSin * tempX) + (tCos * tempZ)
}

// RotateAboutZ rotates about the Z axis by some amount
func (v3d *Vector3d) RotateAboutZ(degRads float64) {
	tCos := math.Cos(degRads)
	tSin := math.Sin(degRads)

	tempX := v3d.X
	tempY := v3d.Y

	v3d.X = (tCos * tempX) - (tSin * tempY)
	v3d.Y = (tSin * tempX) + (tCos * tempY)
}

// RotateAbout3D rotates the vector about some point on the
// X, Y, and Z axes (in that order)
// values in the amounts vector should be angles in radians
func (v3d *Vector3d) RotateAbout3D(point, amounts Vector3d) {
	temp := v3d.GetRelativeCoords3D(point)
	temp.RotateAboutX(amounts.X)
	temp.RotateAboutY(amounts.Y)
	temp.RotateAboutZ(amounts.Z)
	temp.Translate3D(point)
	v3d = &temp
}

// RotateAboutCenter3D does nothing on this as it's the center of itself
func (v3d *Vector3d) RotateAboutCenter3D(ret Vector3d) {}
/*
// MarshallJSON has to be manually implemented for this as inheritance is funky
func (v3d *Vector3d) MarshallJSON() ([]byte, error) {
	return json.Marshal(&struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
		Z float64 `json:"z"`
	}{
		X: v3d.X,
		Y: v3d.Y,
		Z: v3d.Z,
	})
}

// UnmarshallJSON has to be manually implemented for this as inheritance is funky
func (v3d *Vector3d) UnmarshallJSON(data []byte) error {
	temp := &struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
		Z float64 `json:"z"`
	}{}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	v3d.X = temp.X
	v3d.Y = temp.Y
	v3d.Z = temp.Z

	return nil
}
*/
