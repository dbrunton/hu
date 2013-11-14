package hu

import "fmt"
import "bytes"
import "strings"

type signature struct {
	label string
	in signatures
	out signatures
}

type signatures []*signature

type definition struct {
	partOfSpeech *signature
	terms []*term
}

type term struct {
	label string
	definitions []*definition
}

func (s *signature) String() string {
	return fmt.Sprintf("(%s)", s.label)
}

func (s signatures) String() string {
	var b bytes.Buffer
	for i, s := range s {
		if i == 0 {
			fmt.Fprintf(&b, "%s", s.label)
		} else {
			fmt.Fprintf(&b, " %s", s.label)
		}
	}
	return string(b.Bytes())
}

func (d *definition) String() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "%s: ", d.partOfSpeech)
	for i, t := range d.terms {
		if i == 0 {
			fmt.Fprintf(&b, "%s", strings.Title(t.label))
		} else if i == len(d.terms) - 1 {
			fmt.Fprintf(&b, " %s.", t.label)
		} else {
			fmt.Fprintf(&b, " %s", t.label)
		}
	}
	return string(b.Bytes())
}

func (s *signature) Outputs() string {
	return fmt.Sprintf("%s", s.out)
}

func (s *signature) Inputs() string {
	return fmt.Sprintf("%s", s.in)
}

func (t *term) String() string {
	return t.label
}
