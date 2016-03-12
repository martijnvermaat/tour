// https://tour.golang.org/methods/23

package tour

import "io"

type Rot13Reader struct {
	R io.Reader
}

func (r13 Rot13Reader) Read(b []byte) (int, error) {
	n, err := r13.R.Read(b)
	if err == nil {
		for i := 0; i < n; i++ {
			if 'a' <= b[i] && b[i] <= 'z' {
				b[i] = (b[i]+13-'a')%26 + 'a'
			}
			if 'A' <= b[i] && b[i] <= 'Z' {
				b[i] = (b[i]+13-'A')%26 + 'A'
			}
		}
	}
	return n, err
}
