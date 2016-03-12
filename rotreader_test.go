package tour

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestRot13Reader(t *testing.T) {
	testString(t, "Lbh penpxrq gur pbqr!", "You cracked the code!")
	testString(t, "abc123xyz [_] ABC890XYZ", "nop123klm [_] NOP890KLM")
}

func testString(t *testing.T, in string, want string) {
	r := Rot13Reader{strings.NewReader(in)}
	b, _ := ioutil.ReadAll(r)
	rot13 := string(b)

	if rot13 != want {
		t.Errorf("rot13(%q) == %q, want %q", in, rot13, want)
	}
}
