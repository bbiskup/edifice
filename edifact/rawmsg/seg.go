package rawmsg

import (
	"bytes"
	"fmt"
)

type RawSeg struct {
	id    string
	Elems []*RawDataElem
}

// From interface SegOrGroup
func (g *RawSeg) Id() string {
	return g.id
}

func (s *RawSeg) String() string {
	var buf bytes.Buffer
	for _, e := range s.Elems {
		buf.WriteString("\t\t" + e.String() + "\n")
	}
	return fmt.Sprintf("%s\n%s", s.id, buf.String())
}

func (s *RawSeg) AddElem(element *RawDataElem) {
	s.Elems = append(s.Elems, element)
}

func (s *RawSeg) AddElems(elements []*RawDataElem) {
	s.Elems = elements
}

func (s *RawSeg) Dump(indent int) string {
	indentStr := getIndentStr(indent)
	return fmt.Sprintf("%sRawSeg %s\n", indentStr, s.Id())
}

func NewSeg(id string) *RawSeg {
	return &RawSeg{id, []*RawDataElem{}}
}
