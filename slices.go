package tour

func Pic(dx, dy int) [][]uint8 {
	r := make([][]uint8, dy)
	for i := range r {
		r[i] = make([]uint8, dx)
	}

	for x := 0; x < dx; x++ {
		for y := 0; y < dy; y++ {
			r[y][x] = uint8(x ^ y)
		}
	}

	return r
}
