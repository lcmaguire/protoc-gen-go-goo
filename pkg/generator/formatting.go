package generator

import (
	"fmt"
	"strings"

	"github.com/lcmaguire/protoc-gen-go-goo/pkg/templates"
)

func genMethodCaller(in string) string {
	return fmt.Sprintf(templates.MethodCallerTemplate, strings.ToLower(in[0:1]), in)
}

// getParamPKG protoc import will get the last section of an import path to alias a pkg. This func helps you get it.
func getParamPKG(in string) string {
	arr := strings.Split(in, "/")
	return strings.Trim(arr[len(arr)-1], `"`)
}
