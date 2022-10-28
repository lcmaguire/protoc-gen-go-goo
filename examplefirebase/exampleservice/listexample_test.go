package exampleservice

import (
	"fmt"
	testing "testing"

	sample "github.com/lcmaguire/protoc-gen-go-goo/examplefirebase/sample"
)

func TestListExample(t *testing.T) {
	t.Parallel()
	res := &sample.ListExampleResponse{}

	fmt.Println(res.ProtoReflect().Descriptor())
	// Fields().Get(0).Name()
}
