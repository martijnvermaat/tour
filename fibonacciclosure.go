// https://tour.golang.org/moretypes/23

package tour

func Fibonacci() func() int {
	f0, f1 := -1, 1
	return func() int {
		f0, f1 = f1, f0+f1
		return f1
	}
}
