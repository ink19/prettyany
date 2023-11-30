package prettyany

type prettyAnyFieldType struct {
	name string
	fieldType string
	val string
	fields []*prettyAnyFieldType
}

func NewTextFmtField(fieldType string, val string) *prettyAnyFieldType {
	return &prettyAnyFieldType{
		fieldType: fieldType,
		val: val,
		fields: []*prettyAnyFieldType{},
	}
}

func (f *prettyAnyFieldType) AddElem(name string, item *prettyAnyFieldType) {
	item.name = name
	f.fields = append(f.fields, item)
}

func (f *prettyAnyFieldType) String(fmtHandler NewPrettyAnyFmtHandler) string {
	return fmtHandler().String(f)
}
