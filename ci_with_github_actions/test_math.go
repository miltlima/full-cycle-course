package main

import "testing"

func TestSoma(t *testing.T) {

	total := Soma(1, 2)

	if total != 3 {
		t.Errorf("Soma deveria ser 3, mas foi %d", total)
	}
}


