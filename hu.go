package hu

import "fmt"
import "bytes"
import "strings"

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

func (p PartOfSpeech) String() string { return  p.name }

func (d Definition) String() string { return fmt.Sprintf("%s (%s).", d.term, d.PartOfSpeech.name) }

func (l Lexicon) String() string {
	var b bytes.Buffer
	for _, d := range l {
		fmt.Fprintf(&b, "%s", strings.Title(d.term))
	}
	return string(b.Bytes())
}
