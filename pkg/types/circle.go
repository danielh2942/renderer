package types

import (
	"github.com/danielh2942/renderer/pkg/bezier"
	"github.com/danielh2942/renderer/pkg/primitives"
)

const (
	a float64 = 1.00005519
	b float64 = 0.55342686
	c float64 = 0.99873585
)

type Circle struct {
	Center primitives.Vector2d
	Radius float64
}

func (cir *Circle) Render() ([]primitives.Vector2d,error) {
	// semi arc approx p0 = (0,a), p1 = (b,c), p2 = (c,b), p3(a,0)
	// a = 1.00005519
	// b = 0.55342686
	// c = 0.99873585
	radA := cir.Radius * a
	radB := cir.Radius * b
	radC := cir.Radius * c

	tTop := primitives.Vector2d{
		X: cir.Center.X,
		Y: cir.Center.Y - radA,
	}

	tTopRightP1 := primitives.Vector2d{
		X: cir.Center.X + radB,
		Y: cir.Center.Y - radC,
	}

	tTopRightP2 := primitives.Vector2d{
		X: cir.Center.X + radC,
		Y: cir.Center.Y - radB,
	}

	tRight := primitives.Vector2d{
		X: cir.Center.X + radA,
		Y: cir.Center.Y,
	}

	tBottomRightP1 := primitives.Vector2d{
		X: cir.Center.X + radB,
		Y: cir.Center.Y + radC,
	}

	tBottomRightP2 := primitives.Vector2d{
		X: cir.Center.X + radC,
		Y: cir.Center.Y + radB,
	}

	tBottom := primitives.Vector2d{
		X: cir.Center.X,
		Y: cir.Center.Y + radA,
	}

	tBottomLeftP1 := primitives.Vector2d {
		X: cir.Center.X - radB,
		Y: cir.Center.Y + radC,
	}

	tBottomLeftP2 := primitives.Vector2d {
		X: cir.Center.X - radC,
		Y: cir.Center.Y + radB,
	}

	tLeft := primitives.Vector2d{
		X: cir.Center.X - radA,
		Y: cir.Center.Y,
	}

	tTopLeftP1 := primitives.Vector2d{
		X: cir.Center.X - radB,
		Y: cir.Center.Y - radC,
	}

	tTopLeftP2 := primitives.Vector2d{
		X: cir.Center.X - radC,
		Y: cir.Center.Y - radB,
	}

	mArr, _ := bezier.DrawCurve(2000,tTop,tTopRightP1,tTopRightP2,tRight)
	x, _ := bezier.DrawCurve(2000,tRight,tBottomRightP2,tBottomRightP1,tBottom)
	mArr = append(mArr, x...)
	x, _ = bezier.DrawCurve(2000,tBottom,tBottomLeftP1,tBottomLeftP2,tLeft)
	mArr = append(mArr, x...)
	x, _ = bezier.DrawCurve(2000,tLeft,tTopLeftP2,tTopLeftP1,tTop)
	mArr = append(mArr,x...)
	return mArr, nil
}
