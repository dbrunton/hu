package hu

import "testing"

func TestLex(t *testing.T) {
	p, _ := Lex("Now is the time.")
	if p[2] != "the" && len(p) != 3 {
		t.Errorf("Lexing gone wild: ", p)
	}
}
