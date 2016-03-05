package tour

import "testing"

func TestPic(t *testing.T) {
	var p [][]uint8
	p = Pic(9, 7)

	if len(p) != 7 {
		t.Errorf("len(Pic(7, 9)) == %d, want 7", len(p))
	}

	for i, row := range p {
		if len(row) != 9 {
			t.Errorf("len(Pic(7, 9)[%d]) == %d, want 9", i, len(row))
		}
	}
}
