package tour

import "testing"

func TestFibonacci(t *testing.T) {
	fib := []int{0, 1, 1, 2, 3, 5}

	f := Fibonacci()
	for i, want := range fib {
		if got := f(); got != want {
			t.Errorf("Fibonacci()()^%d == %d, want %d",
				i+1, got, want)
		}
	}
}
