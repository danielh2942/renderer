package bezier

import (
	"errors"
	"math"

	"github.com/danielh2942/renderer/pkg/primitives"
)

// binomialCoefficients stores the first 26 rows of pascals triangle,
// Slightly unessecary but it saves some compute time :^)
// Don't think I'll ever even write something with 25 control points but sure be grand
// Also Pascals triangle repeats itself in reverse-order
var binomialCoefficients = [...][]uint64{
	{1},
	{1},
	{1, 2},
	{1, 3},
	{1, 4, 6},
	{1, 5, 10},
	{1, 6, 15, 20},
	{1, 7, 21, 35},
	{1, 8, 28, 56, 70},
	{1, 9, 36, 84, 126},
	{1, 10, 45, 120, 210, 252},
	{1, 11, 55, 165, 330, 462},
	{1, 12, 66, 220, 495, 792, 924},
	{1, 13, 78, 286, 715, 1287, 1716},
	{1, 14, 91, 364, 1001, 2002, 3003, 3432},
	{1, 15, 105, 455, 1365, 3003, 5005, 6435},
	{1, 16, 120, 560, 1820, 4368, 8008, 11440, 12870},
	{1, 17, 136, 680, 2380, 6188, 12376, 19448, 24310},
	{1, 18, 153, 816, 3060, 8568, 18564, 31824, 43758, 48620},
	{1, 19, 171, 969, 3876, 11628, 27132, 50388, 75582, 92378},
	{1, 20, 190, 1140, 4845, 15504, 38760, 77520, 125970, 167960, 184756},
	{1, 21, 210, 1330, 5985, 20349, 54264, 116280, 203490, 293930, 352716},
	{1, 22, 231, 1540, 7315, 26334, 74613, 170544, 319770, 497420, 646646, 705432},
	{1, 23, 253, 1771, 8855, 33649, 100947, 245157, 490314, 817190, 1144066, 1352078},
	{1, 24, 276, 2024, 10626, 42504, 134596, 346104, 735471, 1307504, 1961256, 2496144, 2704156},
	{1, 25, 300, 2300, 12650, 53130, 177100, 480700, 1081575, 2042975, 3268760, 4457400, 5200300},
}

var binomialArrSizes = [...]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10, 11, 11, 12, 12, 13, 13}

// getBinomialCoefficient does exactly what it says on the tin
// It assumes you never have a value greater than 25 BTW
func getBinomialCoefficient(x int, n int) uint64 {
	// some scratch work
	// size of slice is = math.Max(1,(x - (x % 2)))
	// math.min(n, size_of_slice - (n % size_of_slice))
	// 0 1 2
	if n >= binomialArrSizes[x] {
		n = binomialArrSizes[x] - (n % binomialArrSizes[x]) - 1
		if x%2 != 1 {
			n -= 1
		}
	}

	return binomialCoefficients[x][n]
}

// DrawCurve returns a slice containing points on a given bezier curve,
// Bezier curves take control points as arguments, the only time that
// All points given are guaranteed to exist on a line is when there's
// Only two points provided.
// this outputs it in whatever scale you provided the input in
// If the coords are in world space, it returns world space
// If they're screen space, it returns screen space, if it's pixels, it returns pixels etc. etc.
func DrawCurve(quantity uint, points ...primitives.Vector2d) ([]primitives.Vector2d, error) {
	// Add points and stuff
	NewPoints := []primitives.Vector2d{}
	OutputCoords := make([]primitives.Vector2d, quantity)
	var tally int = 0
	for _, x := range points {
		NewPoints = append(NewPoints, x)
		tally++
	}
	if tally < 2 || tally > 25 {
		return nil, errors.New("invalid quantity of points")
	}

	MinMaxVals, _ := primitives.GetVector2dMinMax(NewPoints)
	// Get the linear distance between points for getting the scale factor, this will be used again :)
	MaxPoint := MinMaxVals[1].GetRelativeCoords(MinMaxVals[0])
	scaleFactorY := 1 / MaxPoint.Y
	scaleFactorX := 1 / MaxPoint.X
	InverseMinPoint := primitives.Vector2d{X: MinMaxVals[0].X * -1, Y: MinMaxVals[0].Y * -1}

	for idx, val := range NewPoints {
		val.Translate(InverseMinPoint)
		val.ScaleXY(scaleFactorX, scaleFactorY)
		NewPoints[idx] = val
	}

	// Now the prep is out of the way, time to do the actual calculations
	var t float64 = 0.00
	var count uint = 0
	for count < quantity {
		var tX float64 = 0
		var tY float64 = 0
		for idx, point := range NewPoints {
			// 0^0 is assumed to be 1 for this btw, so I need to put some logic in
			// to prevent it from crashing when that happens
			// math.min(math.max(1-t,0),1)
			tmp := float64(getBinomialCoefficient(tally-1, idx)) * math.Pow(math.Min(math.Max(1-t, 0), 1), float64(tally-1-idx))
			tmp2 := math.Pow(t, float64(idx))
			tX += tmp * point.X * tmp2
			tY += tmp * point.Y * tmp2
		}
		OutputCoords[count] = primitives.Vector2d{X: tX, Y: tY}
		// Translate and scale back to what they should be
		OutputCoords[count].ScaleXY(MaxPoint.X, MaxPoint.Y)
		OutputCoords[count].Translate(MinMaxVals[0])
		t += 1 / float64(quantity)
		count++
	}
	return OutputCoords, nil
}
