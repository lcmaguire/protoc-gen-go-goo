package generator

import (
	"fmt"
	"strings"

	"github.com/lcmaguire/protoc-gen-go-goo/pkg/templates"
)

// move to go template and use gen.Write
func formatMethod(methodCaller string, methodName string, requestType string, responseType string) string {
	return fmt.Sprintf(
		templates.MethodTemplate,
		methodName,
		methodCaller,
		methodName,
		requestType,
		responseType,
	)
}

func formatService(serviceName string, pkg string) string {
	return fmt.Sprintf(templates.ServiceTemplate, serviceName, serviceName, pkg, serviceName)
}

func formatTestFile(method string, service string) string {
	return fmt.Sprintf(templates.TestFileTemplate, method, service, method)
}

// // move to helper
func genMethodCaller(in string) string {
	return fmt.Sprintf(templates.MethodCallerTemplate, strings.ToLower(in[0:1]), in)
}

// move to helper
func getParamPKG(in string) string {
	arr := strings.Split(in, "/")
	return strings.Trim(arr[len(arr)-1], `"`)
}
