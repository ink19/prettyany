package prettyany

import (
	"fmt"
	"strings"
)

type PrettyAnyFmtText struct {}

type PrettyAnyFmt interface {
	String(field *prettyAnyFieldType) string
}

type NewPrettyAnyFmtHandler func() PrettyAnyFmt

func NewPrettyAnyTextFmt() PrettyAnyFmt {
	return &PrettyAnyFmtText{}
}

func (p *PrettyAnyFmtText) String(field *prettyAnyFieldType) string {
	return p.string(0, field)
}

func (p *PrettyAnyFmtText) string(index int, field *prettyAnyFieldType) string {
	lineBuilder := &strings.Builder{}
	lineBuilder.WriteString(strings.Repeat("\t", index))
	lineBuilder.WriteString(fmt.Sprintf("%s(%s):%s\n", field.name, field.fieldType, field.val))

	for _, item := range field.fields {
		lineBuilder.WriteString(p.string(index + 1, item))
	}

	return lineBuilder.String()
}
