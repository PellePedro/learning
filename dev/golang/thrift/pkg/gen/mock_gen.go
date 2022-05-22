package gen

import (
	"bytes"
	"fmt"
	"strings"

	"go.uber.org/thriftrw/compile"
)

type FuncSpec struct {
	Srvice       string
	Name         string
	Arguments    string
	ReturnType   string
	MockedReturn interface{}
}

type ServiceSpec struct {
	ServiceName string
	FuncSpec    []FuncSpec
}

func NewServiceSpec(name string) *ServiceSpec {
	return &ServiceSpec{
		ServiceName: name,
		FuncSpec:    make([]FuncSpec, 0),
	}
}

func GenerateMockService(serviceName string, thriftIdl string) (*ServiceSpec, error) {
	serviceSpec := NewServiceSpec(serviceName)
	compiler, err := compile.Compile(thriftIdl, compile.NonStrict())
	if err != nil {
		return nil, err
	}
	sp, err := compiler.LookupService(serviceName)
	if err != nil {
		return nil, err
	}

	for _, spec := range sp.Functions {
		fSpec := FuncSpec{Name: spec.MethodName()}
		var args bytes.Buffer
		args.WriteString(" ctx context.Context ")

		for _, arg := range spec.ArgsSpec {
			argType := typeReference(compiler.Name, arg.Type)
			args.WriteString(fmt.Sprintf(" %s %s ", arg.Name, argType))
		}
		fSpec.Arguments = strings.Replace(args.String(), "  ", ", ", -1)

		rspec := spec.ResultSpec.ReturnType
		if rspec == nil {
			fSpec.ReturnType = "error"
		} else {
			retType := typeReference(compiler.Name, spec.ResultSpec.ReturnType)
			fSpec.ReturnType = fmt.Sprintf("( %s ,error )", retType)
		}
		serviceSpec.FuncSpec = append(serviceSpec.FuncSpec, fSpec)
	}

	return serviceSpec, nil
}

var tplServiceMock = `{{ define "tplServiceMock" }}
package main

type Handler struct {}

{{- range .FuncSpec }}
func (h *Handler) {{ .Name }}({{.Arguments}}) {{.ReturnType}}{
	fmt.Println("Function {{.Name}} called")
}

{{- end }}
func GetProcessor() *{{ .ServiceName }}Processor {
	return New{{ .ServiceName }}Processor(&Handler{})
}
{{ end }}
`
