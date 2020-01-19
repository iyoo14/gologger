package gologger

import (
	"log"
	"os"
	"testing"
)

func TestLogger(t *testing.T) {
	gopath := os.Getenv("GOPATH")
	logger := NewLogger(gopath+"log", "gologger")
	log.Println("test")
	//logger.Close()
	logger.Println("start test.")
	logger.Println("end test")
}
