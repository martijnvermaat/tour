package tour

import (
	"math"
	"testing"
)

func TestSqrt2(t *testing.T) {
	cases := []float64{-6, -1, 1, 2, 3, 4, 5, 4.5, 78}

	for i := 0; i < len(cases); i += 1 {
		sqrt, err := Sqrt2(cases[i])
		if (cases[i] >= 0) != (err == nil) {
			t.Errorf("_, err Sqrt2(%f); err == %v", cases[i], err)
		}
		if err == nil {
			mathSqrt := math.Sqrt(cases[i])
			d := math.Abs(sqrt - mathSqrt)
			if d > D2 {
				t.Errorf("Sqrt(%f) == %f, want %f, diff %f",
					cases[i], sqrt, mathSqrt, d)
			}
		}
	}
}
