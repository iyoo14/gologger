package gologger

import (
	"fmt"
	"log"
	"os"
	"time"
)

type Logger struct {
	fd     *os.File
	logger *log.Logger
}

func NewLogger(filepath string, prefix string) *Logger {
	n := time.Now()
	ff := fmt.Sprintf("%s-%04d%02d%02d.log", prefix, n.Year(), n.Month(), n.Day())
	file, err := os.OpenFile(filepath+"/"+ff, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	logger := log.New(file, "["+prefix+"]", log.LstdFlags)
	lg := new(Logger)
	lg.fd = file
	lg.logger = logger

	return lg
}

func (lg *Logger) Println(v ...interface{}) {
	lg.logger.Println(v...)
}

func (lg *Logger) Printf(format string, v ...interface{}) {
	lg.logger.Printf(format, v...)
}

func (lg *Logger) Fatal(v ...interface{}) {
	lg.logger.Fatal(v...)
}

func (lg *Logger) Fatalf(format string, v ...interface{}) {
	lg.logger.Fatalf(format, v...)
}

func (lg *Logger) Close() {
	defer lg.fd.Close()
}
