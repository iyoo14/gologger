package gologger

import (
	"os"
	"testing"
)

func TestLogger(t *testing.T) {
	gopath := os.Getenv("GOPATH")
	logger := NewLogger(gopath+"/log", "gologger")
	//logger.Close()
	//log.Println("test")
	//logger.Close()
	logger.Println("start test.")
	logger.Printf("%d, %s", 1, "test")
	logger.Println("end test")
}
