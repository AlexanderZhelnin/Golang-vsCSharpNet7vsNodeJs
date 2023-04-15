package main

import (
	"math"
)

func translate(cs []float64, pr DrawPr) {

	for i := 0; i < len(cs); i += 2 {

		cs[i] = (cs[i] - pr.Left) * pr.Scale
		cs[i+1] = (pr.Top - cs[i+1]) * pr.Scale
	}
}

// func translateRect(r Rect, pr DrawPr) {
// 	r.Left = (r.Left - pr.Left) * pr.Scale
// 	r.Right = (r.Right - pr.Left) * pr.Scale

// 	r.Bottom = (pr.Top - r.Bottom) * pr.Scale
// 	r.Top = (pr.Top - r.Top) * pr.Scale
// }

// func calcRect(coords []float64) Rect {

// 	left := coords[0]
// 	bottom := coords[1]

// 	right := coords[0]
// 	top := coords[1]

// 	for i := 2; i < len(coords); i += 2 {
// 		x := coords[i]
// 		y := coords[i+1]

// 		if left > x {
// 			left = x
// 		}
// 		if bottom > y {
// 			bottom = y
// 		}

// 		if right < x {
// 			right = x
// 		}
// 		if top < y {
// 			top = y
// 		}
// 	}

// 	return Rect{Left: left, Bottom: bottom, Right: right, Top: top}
// }

// func calcMaxLen(coords []float64) int {
// 	result := 0
// 	x1 := coords[0]
// 	y1 := coords[1]
// 	max := 0.0

// 	for i := 2; i < len(coords); i += 2 {

// 		x2 := coords[i]
// 		y2 := coords[i+1]

// 		dx := x2 - x1
// 		dy := y2 - y1

// 		l := math.Sqrt(dx*dx + dy*dy)
// 		if max < l {
// 			max = l
// 			result = i - 2
// 		}
// 	}

// 	return result
// }

func optimize(mas []float64, l float64) []float64 {

	count := len(mas)
	var coords []float64

	if count < 5 {
		coords = make([]float64, count)
		copy(coords, mas)
		return coords
	}

	coords = make([]float64, 0)

	lastCoordX1 := mas[0]
	lastCoordY1 := mas[1]
	lastCoordX2 := mas[2]
	lastCoordY2 := mas[3]

	coords = append(coords, lastCoordX1, lastCoordY1)

	for i := 4; i < count; i += 2 {
		if !IsPointOnLine(lastCoordX1, lastCoordY1, lastCoordX2, lastCoordY2, mas[i], mas[i+1], l) {
			lastCoordX1 = mas[i-2]
			lastCoordY1 = mas[i-1]

			lastCoordX2 = mas[i]
			lastCoordY2 = mas[i+1]

			coords = append(coords, lastCoordX1, lastCoordY1)
		}
	}

	coords = append(coords, mas[count-2], mas[count-1])

	return coords
}

func IsPointOnLine(pX1, pY1, pX2, pY2, pX3, pY3, l float64) bool {

	if pX3 == pX1 && pX3 == pY1 {
		return true
	}

	aX := pX2 - pX1
	aY := pY2 - pY1
	// вектор повёрнутый на 90
	pX4 := -aY + pX3
	pY4 := aX + pY3

	retX := 0.0
	retY := 0.0
	if pX2 == pX1 {
		retX = pX1
		if pY4 == pY3 {
			retY = pY3
		} else if pX4 != pX3 {
			retY = (pX1-pX3)*(pY4-pY3)/(pX4-pX3) + pY3
		}
	} else if pY2 == pY1 {
		retY = pY1
		if pX4 == pX3 {
			retX = pX3
		} else if pY4 != pY3 {
			retX = (pY1-pY3)*(pX4-pX3)/(pY4-pY3) + pX3
		}
	} else if pX4 == pX3 {
		retX = pX3
		retY = (pX3-pX1)*(pY2-pY1)/(pX2-pX1) + pY1
	} else if pY4 == pY3 {
		retY = pY3
		retX = (pY3-pY1)*(pX2-pX1)/(pY2-pY1) + pX1
	} else {
		k1 := (pY2 - pY1) / (pX2 - pX1)
		k2 := (pY4 - pY3) / (pX4 - pX3)

		retX = (k1*pX1 - k2*pX3 + pY3 - pY1) / (k1 - k2)
		retY = (retX-pX1)*k1 + pY1
	}

	retX -= pX3
	retY -= pY3

	return math.Sqrt(retX*retX+retY*retY) < l

}
