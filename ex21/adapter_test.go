package main

import "testing"

func TestAdapter(t *testing.T) {
	var w WantToUse
	adapter := NewAdapter(&w)

	got := adapter.SomeFuncFromFamousLib()
	want := "AlreadyExistedFunc"
	if got != want {
		t.Fatalf(`want %s, got %s`, want, got)
	}
}
