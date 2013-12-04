package hu

import "testing"
import "fmt"

func TestPanity(t *testing.T) {

        if false {
                t.Errorf("Pomething wacky is wrong.")
        }


}

var P1 = PartOfSpeech{"P1", nil, nil}
var P2 = PartOfSpeech{"P2", []PartOfSpeech{P1, P1}, nil}
var P3 = PartOfSpeech{"P3", []PartOfSpeech{P1, P2}, []PartOfSpeech{P2, P2}}
var D1 = Definition{"D1", P1}
var D2 = Definition{"D2", P3}

var L1 = Lexicon{D1, D2}

func TestDefinitionString(t *testing.T) {
	if D2.String() != "D2 (P3)." {
		t.Errorf("Expected 'D2 (P3).' Got ", fmt.Sprintf("%s", D2))
	}
}

func TestLexiconString(t *testing.T) {
	if L1.String() != "D1 (P1).\nD2 (P3).\n" {
		t.Errorf(L1.String())
	}
}
