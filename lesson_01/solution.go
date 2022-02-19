package square

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

import "math"

type figureSides int

const SidesTriangle figureSides = 3
const SidesSquare figureSides = 4
const SidesCircle figureSides = 0

func CalcSquare(sideLen float64, sidesNum figureSides) (res float64) {
	switch sidesNum {
	case SidesTriangle:
		res = math.Pow(sideLen, 2) * math.Sqrt(3) / 4
	case SidesSquare:
		res = math.Pow(sideLen, 2)
	case SidesCircle:
		res = math.Pow(sideLen, 2) * math.Pi
	default:
		res = 0
	}
	return
}
