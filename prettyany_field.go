package prettyany

type prettyAnyFieldType struct {
	fieldType string
	val string
	fields map[string]*prettyAnyFieldType
}

func NewTextFmtField(fieldType string, val string) *prettyAnyFieldType {
	return &prettyAnyFieldType{
		fieldType: fieldType,
		val: val,
		fields: map[string]*prettyAnyFieldType{},
	}
}

func (f *prettyAnyFieldType) AddElem(fieldName string, item *prettyAnyFieldType) {
	f.fields[fieldName] = item
}

func (f *prettyAnyFieldType) String(fmtHandler NewPrettyAnyFmtHandler) string {
	return fmtHandler().String(f)
}
