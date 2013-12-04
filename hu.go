package hu

import "fmt"
import "bytes"

type PartOfSpeech struct {
	name string
	in []PartOfSpeech
	out []PartOfSpeech
}

type Definition struct {
	term string // headword?
	PartOfSpeech
//	[]*term
}

type Lexicon []Definition

type Sentence []string

func (p PartOfSpeech) String() string { return  p.name }

func (d Definition) String() string { return fmt.Sprintf("%s (%s).", d.term, d.PartOfSpeech.name) }

func (l Lexicon) String() string {
	var b bytes.Buffer
	for _, d := range l {
		fmt.Fprintf(&b, "%s\n", d)
	}
	return string(b.Bytes())
}
