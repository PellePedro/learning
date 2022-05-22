package gen

import (
	"fmt"

	"go.uber.org/thriftrw/compile"
)

func typeReference(pkgName string, spec compile.TypeSpec) string {
	switch s := spec.(type) {
	case *compile.BoolSpec:
		return "bool"
	case *compile.I8Spec:
		return "int8"
	case *compile.I16Spec:
		return "int16"
	case *compile.I32Spec:
		return "int32"
	case *compile.I64Spec:
		return "int64"
	case *compile.DoubleSpec:
		return "float64"
	case *compile.StringSpec:
		return "string"
	case *compile.BinarySpec:
		return "[]byte"
	case *compile.StructSpec:
		return fmt.Sprintf("*%s.%s", pkgName, s.Name)
	case *compile.ListSpec:
		// Add packagename to no native types
		name := s.ThriftName()
		tName := s.ValueSpec.ThriftName()
		_ = name
		_ = tName
		if baseType, ok := baseTypes[s.ValueSpec.ThriftName()]; ok {
			return baseType
		}
		return fmt.Sprintf("[]*%s.%s", pkgName, s.ValueSpec.ThriftName())
	case *compile.MapSpec:
		k := typeReference(pkgName, s.KeySpec)
		v := typeReference(pkgName, s.ValueSpec)
		return fmt.Sprintf("map[%s]%s", k, v)
	default:
		a := s
		_ = a
	}

	return ""
}

func getReturnSignature(spec compile.TypeSpec) string {
	switch s := spec.(type) {
	case *compile.BoolSpec:
		return "(bool, error)"
	case *compile.I8Spec:
		return "(uint8, error)"
	case *compile.I16Spec:
		return "(uint16, error)"
	case *compile.I32Spec:
		return "(uint32, error)"
	case *compile.I64Spec:
		return "(uint64, error)"
	case *compile.DoubleSpec:
		return "(double, error)"
	case *compile.StringSpec:
		return "(string, error)"
	case *compile.StructSpec:
		return fmt.Sprintf("(%s, error)", s.Name)
	}
	return "(error)"
}

var baseTypes = map[string]string{
	"bool":   "bool",
	"byte":   "int8",
	"i8":     "int8",
	"i16":    "int16",
	"i32":    "int32",
	"i64":    "int64",
	"double": "float64",
	"string": "string",
	"binary": "[]byte",
}
