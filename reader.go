// https://tour.golang.org/methods/22

package tour

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	l := len(b)
	for i := 0; i < l; i++ {
		b[i] = 'A'
	}
	return l, nil
}
