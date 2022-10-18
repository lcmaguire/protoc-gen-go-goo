package generator

import (
	"fmt"
	"strings"

	"github.com/lcmaguire/protoc-gen-go-goo/pkg/templates"
)

func genMethodCaller(in string) string {
	return fmt.Sprintf(templates.MethodCallerTemplate, strings.ToLower(in[0:1]), in)
}

// move to helper
func getParamPKG(in string) string {
	arr := strings.Split(in, "/")
	return strings.Trim(arr[len(arr)-1], `"`)
}
