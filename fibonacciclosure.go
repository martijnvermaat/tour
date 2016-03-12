// https://tour.golang.org/moretypes/23

package tour

func Fibonacci() func() int {
	x, y := 0, 1
	return func() int {
		defer func() { x, y = y, x+y }()
		return x
	}
}
