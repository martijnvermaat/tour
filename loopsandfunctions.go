// https://tour.golang.org/flowcontrol/8

package tour

import "math"

// Max diff of the squared result (result has a smaller diff)
const D float64 = 1e-5

func Sqrt(x float64) (z float64) {
	z = 1.0
	for {
		d := z*z - x
		if math.Abs(d) <= D {
			return
		}
		z = z - d/(2*x)
	}
}
