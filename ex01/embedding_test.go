package ex01

import "testing"

func TestSomeParentFunc(t *testing.T) {
	child := &Action{
		Human: Human{
			fieldOne: 456,
			fieldSec: "GOOD BYE",
		},
	}

	got := child.someParentFunc()
	want := "456 GOOD BYE"

	if got != want {
		t.Fatalf("got %s want %s", got, want)
	}
}
