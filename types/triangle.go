package types

import (
	"math"
)

// MakeTriangle creates a triangle struct populated with all the appropriate stuff
func MakeTriangle(p1, p2, p3 Vector2d) Triangle {
	return Triangle{
		Point1: p1,
		Point2: p2,
		Point3: p3,
		Centroid: Vector2d{
			X: (p1.X + p2.X + p3.X) / 3,
			Y: (p1.Y + p2.Y + p3.Y) / 3,
		},
	}
}

// Triangle is a type describing a Triangle on a 2D plane
type Triangle struct {
	Point1   Vector2d `json:"Point1"`   // Point 1 of the Triangle
	Point2   Vector2d `json:"Point2"`   // Point 2 of the Triangle
	Point3   Vector2d `json:"Point3"`   // Point 3 of the Triangle
	Centroid Vector2d `json:"Centroid"` // The Centroid of the Triangle
}

// calculateCentorid recalculates the Centroid of a triangle, this is especially important with scaling operations
func (t *Triangle) calculateCentroid() {
	t.Centroid = Vector2d{
		X: (t.Point1.X + t.Point2.X + t.Point3.X) / 3,
		Y: (t.Point1.Y + t.Point2.Y + t.Point3.Y) / 3,
	}
}

// Translate translates a triangle by some vector X,Y
func (t *Triangle) Translate(vec Vector2d) {
	t.Point1.Translate(vec)
	t.Point2.Translate(vec)
	t.Point3.Translate(vec)
	t.Centroid.Translate(vec)
}

// Scale scales all the points of a triangle by a given amount
func (t *Triangle) Scale(scaleFactor float64) {
	t.Point1.Scale(scaleFactor)
	t.Point2.Scale(scaleFactor)
	t.Point3.Scale(scaleFactor)
	t.calculateCentroid()
}

// ScaleX scales the X values of a triangle by a given amount
func (t *Triangle) ScaleX(scaleFactorX float64) {
	t.Point1.ScaleX(scaleFactorX)
	t.Point2.ScaleX(scaleFactorX)
	t.Point3.ScaleX(scaleFactorX)
	t.calculateCentroid()
}

// ScaleY scales the Y values of a triangle by a given amount
func (t *Triangle) ScaleY(scaleFactorY float64) {
	t.Point1.ScaleY(scaleFactorY)
	t.Point2.ScaleY(scaleFactorY)
	t.Point3.ScaleY(scaleFactorY)
	t.calculateCentroid()
}

// ScaleXY scales the X and Y values of a triangle by a given amount
func (t *Triangle) ScaleXY(scaleFactorX float64, scaleFactorY float64) {
	t.Point1.ScaleXY(scaleFactorX, scaleFactorY)
	t.Point2.ScaleXY(scaleFactorX, scaleFactorY)
	t.Point3.ScaleXY(scaleFactorX, scaleFactorY)
	t.calculateCentroid()
}

// RotateAbout rotates a triangle about a given point
func (t *Triangle) RotateAbout(vec Vector2d, angleRads float64) {
	t.Point1.RotateAbout(vec, angleRads)
	t.Point2.RotateAbout(vec, angleRads)
	t.Point3.RotateAbout(vec, angleRads)
	t.Centroid.RotateAbout(vec, angleRads)
}

// RotateAboutCenter rotates a triangle about the center of itself
func (t *Triangle) RotateAboutCenter(angleRads float64) {
	t.Point1.RotateAbout(t.Centroid, angleRads)
	t.Point2.RotateAbout(t.Centroid, angleRads)
	t.Point3.RotateAbout(t.Centroid, angleRads)
}

// Render renders the triangle
func (t *Triangle) Render() ([]Vector2d, error) {
	// 1 -> 2
	// 2 -> 3
	// 1 -> 3
	l := t.Point1.GetRelativeCoords(t.Point2)
	maxLineLength := l.GetAbs()
	l = t.Point2.GetRelativeCoords(t.Point3)
	maxLineLength = math.Max(l.GetAbs(), maxLineLength)
	l = t.Point1.GetRelativeCoords(t.Point3)
	maxLineLength = math.Max(l.GetAbs(), maxLineLength)

	lineLen := int(math.Ceil(maxLineLength))

	pts := 3 * lineLen

	newCoords := make([]Vector2d, pts)

	tChange := 1 / maxLineLength

	i := 0

	for tv := 0.0; tv < 1.0; tv += tChange {
		minT := 1 - tv

		newCoords[i] = Vector2d{
			X: (minT * t.Point1.X) + (tv * t.Point2.X),
			Y: (minT * t.Point1.Y) + (tv * t.Point2.Y),
		}

		newCoords[lineLen+i] = Vector2d{
			X: (minT * t.Point2.X) + (tv * t.Point3.X),
			Y: (minT * t.Point2.Y) + (tv * t.Point3.Y),
		}

		newCoords[(2*lineLen)+i] = Vector2d{
			X: (minT * t.Point3.X) + (tv * t.Point1.X),
			Y: (minT * t.Point3.Y) + (tv * t.Point1.Y),
		}

		i++
	}

	return newCoords, nil
}
