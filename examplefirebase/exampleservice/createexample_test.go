package exampleservice

import (
	"testing"
)

// examplefirebase/exampleservice/createexample.go
// go test ./examplefirebase/exampleservice/... -run="TestMong" count 1
// docker run --name mongodb -d -p 27017:27017 mongo to run locally . https://www.mongodb.com/compatibility/docker
func TestMong(t *testing.T) {
	t.Parallel()
	mongito()
}
