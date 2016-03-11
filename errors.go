// https://tour.golang.org/methods/20

package tour

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

// Max diff of the squared result (result has a smaller diff)
const D2 float64 = 1e-5

func Sqrt2(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := 1.0
	for {
		d := z*z - x
		if math.Abs(d) <= D {
			return z, nil
		}
		z = z - d/(2*x)
	}
}
