package prettyany

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	"golang.org/x/exp/constraints"
)

type PrettyAny struct{}

func NewPrettyAny() *PrettyAny {
	return &PrettyAny{}
}

func (p *PrettyAny) Pretty(val any) string {
	fmtValue := p.dispatch("", reflect.ValueOf(val))
	return fmtValue.String(NewPrettyAnyTextFmt)
}

func (p *PrettyAny) dispatch(typePrefix string, val reflect.Value) *prettyAnyFieldType {
	tval := val.Type()
	switch tval.Kind() {
	case
		reflect.Bool,
		reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Uintptr,
		reflect.Float32,
		reflect.Float64,
		reflect.Complex64,
		reflect.Complex128,
		reflect.UnsafePointer,
		reflect.String:
		return p.printDirect(typePrefix, val)
	case reflect.Pointer:
		return p.printPointer(typePrefix, val)
	case reflect.Interface:
		return p.dispatch(typePrefix, val.Elem())
	case reflect.Map:
		return p.printMap(typePrefix, val)
	case reflect.Array, reflect.Slice:
		return p.printSlice(typePrefix, val)
	case reflect.Struct:
		return p.printStruct(typePrefix, val)
	case reflect.Chan, reflect.Func:
		return p.printType(typePrefix, val)
	}
	return nil
}

func (p *PrettyAny) printDirect(typePrefix string, val reflect.Value) *prettyAnyFieldType {
	field := NewTextFmtField(typePrefix+val.Type().Kind().String(), fmt.Sprintf("%v", val))
	return field
}

func (p *PrettyAny) printPointer(typePrefix string, val reflect.Value) *prettyAnyFieldType {
	if val.IsNil() {
		field := NewTextFmtField(typePrefix+val.Type().Kind().String(), "nil")
		return field
	}
	return p.dispatch("*"+typePrefix, val.Elem())
}

func (p *PrettyAny) printMap(typePrefix string, val reflect.Value) *prettyAnyFieldType {
	keys := val.MapKeys()
	sort.Slice(keys, func(i, j int) bool {
		return keys[i].String() < keys[j].String()
	})
	field := NewTextFmtField(typePrefix+val.Type().String(), "")
	for _, item := range keys {
		iVal := val.MapIndex(item)
		iField := p.dispatch("", iVal)
		field.AddElem(item.String(), iField)
	}
	return field
}

func printIntergerArray[T constraints.Integer, S ~[]T](s S) string {
	sb := &strings.Builder{}
	for i, item := range s {
		sb.WriteString(fmt.Sprintf("%d", item))
		if i != len(s) - 1 {
			sb.WriteString(",")
		}
	}
	return sb.String()
}

func printFloatArray[T constraints.Float, S ~[]T](s S) string {
	sb := &strings.Builder{}
	for i, item := range s {
		sb.WriteString(fmt.Sprintf("%f", item))
		if i != len(s) - 1 {
			sb.WriteString(",")
		}
	}
	return sb.String()
}

func (p *PrettyAny) printSimpleSlice(s any) string {
	switch sv := s.(type) {
	case []int:
		return printIntergerArray(sv)
	case []uint:
		return printIntergerArray(sv)
	case []int16:
		return printIntergerArray(sv)
	case []uint16:
		return printIntergerArray(sv)
	case []int32:
		return printIntergerArray(sv)
	case []uint32:
		return printIntergerArray(sv)
	case []int64:
		return printIntergerArray(sv)
	case []uint64:
		return printIntergerArray(sv)

	case []float32:
		return printFloatArray(sv)
	case []float64:
		return printFloatArray(sv)
	default:
		return "not support type"
	}
}

func (p *PrettyAny) printSlice(typePrefix string, val reflect.Value) *prettyAnyFieldType {
	switch val.Type().Elem().Kind() {
	// 针对[]byte、[]int8、[]uint8等类型，直接输出hex
	case reflect.Uint8:
		field := NewTextFmtField(typePrefix+val.Type().String(), fmt.Sprintf("%x", val.Bytes()))
		return field
	case reflect.Int8:
		cval := make([]byte, 0, val.Len())
		for _, item := range val.Interface().([]int8) {
			cval = append(cval, byte(item))
		}
		field := NewTextFmtField(typePrefix+val.Type().String(), fmt.Sprintf("%x", cval))
		return field
	// 针对[]int、[]uint等类型，直接平铺输出
	case reflect.Int, reflect.Uint, reflect.Int16, reflect.Uint16, reflect.Int32, reflect.Uint32, reflect.Int64, reflect.Uint64, reflect.Float32, reflect.Float64:
		field := NewTextFmtField(typePrefix+val.Type().String(), p.printSimpleSlice(val.Interface()))
		return field
	}

	iNum := val.Len()
	field := NewTextFmtField(typePrefix+val.Type().String(), "")
	for i := 0; i < iNum; i++ {
		iVal := val.Index(i)
		iField := p.dispatch("", iVal)
		field.AddElem(fmt.Sprintf("%d", i), iField)
	}
	return field
}

func (p *PrettyAny) printStruct(typePrefix string, val reflect.Value) *prettyAnyFieldType {
	fieldNum := val.NumField()
	field := NewTextFmtField(typePrefix+val.Type().String(), "")
	fieldNames := []string{}
	for i := 0; i < fieldNum; i++ {
		fieldNames = append(fieldNames, val.Type().Field(i).Name)
	}
	sort.Slice(fieldNames, func(i, j int) bool {
		return fieldNames[i] < fieldNames[j]
	})
	for _, fieldName := range fieldNames {
		iField := p.dispatch("", val.FieldByName(fieldName))
		field.AddElem(fieldName, iField)
	}
	return field
}

func (p *PrettyAny) printType(typePrefix string, val reflect.Value) *prettyAnyFieldType {
	field := NewTextFmtField(typePrefix+val.Type().String(), "")
	return field
}
