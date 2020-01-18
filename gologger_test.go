package gologger

import (
	"os"
	"testing"
)

func TestLogger(t *testing.T) {
	gopath := os.Getenv("GOPATH")
	logger := Logger(gopath+"log", "gologger")
	logger.Println("start test.")
	logger.Println("end test")
}
