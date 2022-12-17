package exampleservice

import (
	"testing"
)

// examplefirebase/exampleservice/createexample.go
// go test ./examplefirebase/exampleservice/... -run="TestMong" count 1
func TestMong(t *testing.T) {
	t.Parallel()
	mongito()
}
