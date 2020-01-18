package gologger

import (
	"fmt"
	"log"
	"os"
	"time"
)

func Logger(filepath string, prefix string) *log.Logger {
	n := time.Now()
	ff := fmt.Sprintf("%04d%02d%02d", n.Year(), n.Month(), n.Day())
	file, _ := os.OpenFile(filepath+"/"+ff+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer func() {
		file.Close()
	}()
	logger := log.New(file, "["+prefix+"]", log.LstdFlags)
	return logger
}
