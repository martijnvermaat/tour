package tour

func Fibonacci() func() int {
	f0, f1 := -1, 1
	return func() (f2 int) {
		f2 = f0 + f1
		f0, f1 = f1, f2
		return
	}
}
