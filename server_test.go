package panda

import (
	"fmt"
	"testing"
)

func ControllerTest(c *Controller) {

	c.Write([]byte("hello"))
}

// go test -run=TestServer
func TestServer(t *testing.T) {
	HandlerRouter(CASUAL, "/", ControllerTest)
	Run()
}
