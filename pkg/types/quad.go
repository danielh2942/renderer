package types

import (
	"math"

	"github.com/danielh2942/renderer/pkg/primitives"
)

// Quad is a 4 point Polygon
type Quad struct {
	/*
		Quad line mappings

		1->2
		1->3
		2->4
		3->4
	*/
	Point1 primitives.Vector2d `json:"Point1"` // Top Left
	Point2 primitives.Vector2d `json:"Point2"` // Top Right
	Point3 primitives.Vector2d `json:"Point3"` // Bottom Left
	Point4 primitives.Vector2d `json:"Point4"` // Bottom Right
}

func (q *Quad) Translate(vec primitives.Vector2d) {
	q.Point1.Translate(vec)
	q.Point2.Translate(vec)
	q.Point3.Translate(vec)
	q.Point4.Translate(vec)
}

func (q *Quad) Scale(scaleFactor float64) {
	q.Point1.Scale(scaleFactor)
	q.Point2.Scale(scaleFactor)
	q.Point3.Scale(scaleFactor)
	q.Point4.Scale(scaleFactor)
}

func (q *Quad) ScaleX(scaleFactor float64) {
	q.Point1.ScaleX(scaleFactor)
	q.Point2.ScaleX(scaleFactor)
	q.Point3.ScaleX(scaleFactor)
	q.Point4.ScaleX(scaleFactor)
}

func (q *Quad) ScaleY(scaleFactor float64) {
	q.Point1.ScaleY(scaleFactor)
	q.Point2.ScaleY(scaleFactor)
	q.Point3.ScaleY(scaleFactor)
	q.Point4.ScaleY(scaleFactor)
}

func (q *Quad) ScaleXY(scaleFactorX float64, scaleFactorY float64) {
	q.Point1.ScaleXY(scaleFactorX, scaleFactorY)
	q.Point2.ScaleXY(scaleFactorX, scaleFactorY)
	q.Point3.ScaleXY(scaleFactorX, scaleFactorY)
	q.Point4.ScaleXY(scaleFactorX, scaleFactorY)
}

func (q *Quad) RotateAbout(vec primitives.Vector2d, angleRads float64) {
	q.Point1.RotateAbout(vec, angleRads)
	q.Point2.RotateAbout(vec, angleRads)
	q.Point3.RotateAbout(vec, angleRads)
	q.Point4.RotateAbout(vec, angleRads)
}

func (q *Quad) Render() ([]primitives.Vector2d, error) {

	// Get longest line
	l1 := q.Point1.GetRelativeCoords(q.Point2)
	maxLineLength := l1.GetAbs()
	l1 = q.Point1.GetRelativeCoords(q.Point3)
	maxLineLength = math.Max(maxLineLength, l1.GetAbs())
	l1 = q.Point2.GetRelativeCoords(q.Point4)
	maxLineLength = math.Max(maxLineLength, l1.GetAbs())
	l1 = q.Point3.GetRelativeCoords(q.Point4)
	maxLineLength = math.Max(maxLineLength, l1.GetAbs())

	// Allocate 4x maximum line length for every line on the quad
	lineLen := int(math.Ceil(maxLineLength))
	pts := 4 * lineLen
	newCoords := make([]primitives.Vector2d, pts)
	tChange := 1 / maxLineLength
	i := 0
	// Draw it clockwise
	for t := 0.0; t <= 1.0; t += tChange {
		minT := 1 - t
		//1->2
		newCoords[i] = primitives.Vector2d{
			X: (minT * q.Point1.X) + (t * q.Point2.X),
			Y: (minT * q.Point1.Y) + (t * q.Point2.Y),
		}
		//2->4
		newCoords[lineLen+i] = primitives.Vector2d{
			X: (minT * q.Point2.X) + (t * q.Point4.X),
			Y: (minT * q.Point2.Y) + (t * q.Point4.Y),
		}
		//3->4
		newCoords[(2*lineLen)+i] = primitives.Vector2d{
			X: (minT * q.Point4.X) + (t * q.Point3.X),
			Y: (minT * q.Point4.Y) + (t * q.Point3.Y),
		}
		//1->3
		newCoords[(3*lineLen)+i] = primitives.Vector2d{
			X: (minT * q.Point3.X) + (t * q.Point1.X),
			Y: (minT * q.Point3.Y) + (t * q.Point1.Y),
		}
		i++
	}
	return newCoords, nil
}
