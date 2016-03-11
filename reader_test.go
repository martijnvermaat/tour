package tour

import "testing"

func TestMyReader(t *testing.T) {
	var b []byte
	r := MyReader{}

	for i := 0; i < 10; i++ {
		b = make([]byte, 8)
		if n, err := r.Read(b); err != nil {
			t.Errorf("n, err := MyReader.Read(b); err == %v", err)
		} else {
			if n <= 0 {
				t.Errorf("n, err := MyReader.Read(b); n == %d", n)
			}
			if n > len(b) {
				t.Errorf("n, err := MyReader.Read(b); n == %d; len(b) == %d", n, len(b))
			}
			for j := 0; j < n; j++ {
				if b[j] != 'A' {
					t.Errorf("MyReader.Read(b); b[%d] == '%c', want 'A'", j, b[j])
				}
			}
		}
	}
}
