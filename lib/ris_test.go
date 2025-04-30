package ris

import "testing"

func TestRISBasic(t *testing.T) {
	params := Params{A: 10, K: 100, Q: 1, B: 0.05, V: 75}
	score := RIS(150, 75, params)

	if score <= 0 {
		t.Errorf("Expected positive RIS, got %.2f", score)
	}
}
