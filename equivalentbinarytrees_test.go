package tour

import (
	"golang.org/x/tour/tree"
	"testing"
)

func TestWalk(t *testing.T) {
	ch := make(chan int)
	go Walk(tree.New(1), ch)

	for i := 0; i < 10; i++ {
		if n, ok := <-ch; !ok {
			t.Errorf("Walk did not send, want %d", i+1)
		} else if n != i+1 {
			t.Errorf("Walk sends %d, want %d", n, i+1)
		}
	}

	if n, ok := <-ch; ok {
		t.Errorf("Walk sends %d, want nothing", n)
	}
}

func TestSame(t *testing.T) {
	var a, b *tree.Tree

	a, b = tree.New(1), tree.New(1)
	if !Same(a, b) {
		t.Errorf("Same(\n  %v,\n  %v\n) == false, want true", a, b)
	}

	a, b = tree.New(1), tree.New(2)
	if Same(a, b) {
		t.Errorf("Same(\n  %v,\n  %v\n) == true, want false", a, b)
	}
}
