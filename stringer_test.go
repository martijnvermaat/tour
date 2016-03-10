package tour

import (
	"fmt"
	"testing"
)

func TestIPAddr(t *testing.T) {
	for i, c := range ipAddrTests {
		if got := fmt.Sprint(c.in); got != c.want {
			t.Errorf("Sprint(cases[%d]) == %v, want %v",
				i, got, c.want)
		}
	}
}

var ipAddrTests = []struct {
	in   IPAddr
	want string
}{
	{IPAddr{1, 2, 3, 4}, "1.2.3.4"},
	{IPAddr{127, 0, 0, 1}, "127.0.0.1"},
	{IPAddr{8, 8, 8, 8}, "8.8.8.8"},
}
