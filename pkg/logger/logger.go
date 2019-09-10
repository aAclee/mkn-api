package logger

import (
	"log"
	"os"
)

// Logger reprents an object for logging ot os.Stdout
type Logger struct {
	log *log.Logger
}

// CreateLogger returns a new Logger that writes to standard out
func CreateLogger() Logger {
	return Logger{
		log: log.New(os.Stdout, "[MKN-SERVER] ", log.Ldate+log.Ltime+log.LUTC),
	}
}

// Info will log arguments in the manner of fmt.Print
func (l Logger) Info(v ...interface{}) {
	l.log.Print("INFO: ", v)
}

// Infof will log arguments are handled in the manner of fmt.Printf.
func (l Logger) Infof(format string, v ...interface{}) {
	l.log.Printf("INFO: "+format, v)
}

// Error will log arguments in the manner of fmt.Print.
func (l Logger) Error(v ...interface{}) {
	l.log.Print("ERROR: ", v)
}

// Errorf will log arguments are handled in the manner of fmt.Printf.
func (l Logger) Errorf(format string, v ...interface{}) {
	l.log.Printf("ERROR: "+format, v)
}
