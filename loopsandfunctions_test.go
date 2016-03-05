package tour

import (
	"math"
	"testing"
)

func TestSqrt(t *testing.T) {
	cases := []float64{1, 2, 3, 4, 5, 4.5, 78}

	for i := 0; i < len(cases); i += 1 {
		sqrt := Sqrt(cases[i])
		mathSqrt := math.Sqrt(cases[i])
		d := math.Abs(sqrt - mathSqrt)
		if d > D {
			t.Errorf("Sqrt(%f) == %f, want %f, diff %f",
				cases[i], sqrt, mathSqrt, d)
		}
	}
}
