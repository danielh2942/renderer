package types

import (
	"math"

	"github.com/danielh2942/renderer/pkg/primitives"
)

type NPoly struct {
	// each point is connected to the next and the last is connected to the 0th one
	Points []primitives.Vector2d `json:"Points"`
}

// Translate moves all points by some vector
func (np *NPoly) Translate(vec primitives.Vector2d) {
	for i, j := range np.Points {
		j.Translate(vec)
		np.Points[i] = j
	}
}

// Scale scales all points by some value
func (np *NPoly) Scale(scaleFactor float64) {
	for i, j := range np.Points {
		j.Scale(scaleFactor)
		np.Points[i] = j
	}
}

// ScaleX scales the X value of all points by some value
func (np *NPoly) ScaleX(scaleFactor float64) {
	for i, j := range np.Points {
		j.ScaleX(scaleFactor)
		np.Points[i] = j
	}
}

// ScaleY scales all Y values by some value
func (np *NPoly) ScaleY(scaleFactor float64) {
	for i, j := range np.Points {
		j.ScaleY(scaleFactor)
		np.Points[i] = j
	}
}

// ScaleXY scales X and Y values by different values
func (np *NPoly) ScaleXY(scaleFactorX float64, scaleFactorY float64) {
	for i, j := range np.Points {
		j.ScaleXY(scaleFactorX, scaleFactorY)
		np.Points[i] = j
	}
}

// RotateAbout rotates all points around some arbitrary point by some magnitude
func (np *NPoly) RotateAbout(vec primitives.Vector2d, angleRads float64) {
	for i, j := range np.Points {
		j.RotateAbout(vec, angleRads)
		np.Points[i] = j
	}
}

// getCenter returns the Center Vector
func (np *NPoly) getCenter() primitives.Vector2d {
	// Get center of the polygon
	var mX float64 = 0
	var mY float64 = 0

	for _, pts := range np.Points {
		mX += pts.X
		mY += pts.Y
	}
	mX /= float64(len(np.Points))
	mY /= float64(len(np.Points))

	return primitives.Vector2d{X: mX, Y: mY}
}

// RotateAboutCenter rotates about the center of a polygon
func (np *NPoly) RotateAboutCenter(angleRads float64) {

	center := np.getCenter()

	// Rotate about it
	for i, pt := range np.Points {
		pt.RotateAbout(center, angleRads)
		np.Points[i] = pt
	}
}

// GetTriangles decomposes a shape into it's composite triangles
func (np *NPoly) GetTriangles() []Triangle {
	if len(np.Points) < 3 {
		return []Triangle{}
	} else if len(np.Points) == 3 {
		return []Triangle{MakeTriangle(np.Points[0], np.Points[1], np.Points[2])}
	}

	// Get the center point

	center := np.getCenter()

	// Figure out triangles

	triangles := make([]Triangle, len(np.Points))
	for i := range np.Points {
		if i == len(np.Points)-1 {
			triangles[i] = MakeTriangle(center, np.Points[i], np.Points[0])
		} else {
			triangles[i] = MakeTriangle(center, np.Points[i], np.Points[i+1])
		}
	}
	return triangles
}

// Render draws the poly as points
func (np *NPoly) Render() ([]primitives.Vector2d, error) {
	maxLineLength := 0.0
	for i, v := range np.Points {
		// always check next val (or if it's the last one, check the first)
		if i == len(np.Points)-1 {
			tLine := v.GetRelativeCoords(np.Points[0])
			maxLineLength = math.Max(maxLineLength, tLine.GetAbs())
		} else {
			tLine := v.GetRelativeCoords(np.Points[i+1])
			maxLineLength = math.Max(maxLineLength, tLine.GetAbs())
		}
	}

	lineLen := int(math.Ceil(maxLineLength))
	pts := len(np.Points) * lineLen
	newCoords := make([]primitives.Vector2d, pts+1)
	tChange := 1 / maxLineLength
	i := 0

	for t := 0.0; t < 1.0; t += tChange {
		minT := 1 - t
		for j, v := range np.Points {
			if j == len(np.Points)-1 {
				newCoords[(j*lineLen)+i] = primitives.Vector2d{
					X: (minT * v.X) + (t * np.Points[0].X),
					Y: (minT * v.Y) + (t * np.Points[0].Y),
				}
			} else {
				newCoords[(j*lineLen)+i] = primitives.Vector2d{
					X: (minT * v.X) + (t * np.Points[j+1].X),
					Y: (minT * v.Y) + (t * np.Points[j+1].Y),
				}
			}
		}
		i++
	}
	return newCoords, nil
}

func (np *NPoly) RenderComposite() ([][]primitives.Vector2d, error) {

	triangles := np.GetTriangles()

	outp := make([][]primitives.Vector2d, len(np.Points))
	for i, t := range triangles {
		tPts, _ := t.Render()
		outp[i] = tPts
	}
	return outp, nil
}
