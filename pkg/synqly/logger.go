package synqly

import (
	"fmt"
	"os"
	"time"
)

type Logger struct{}

func (l Logger) Printf(format string, v ...interface{}) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	msg := fmt.Sprintf(format, v...)
	os.Stdout.WriteString("[SYNQLY] " + timestamp + " " + msg + "\n")
}

func (l Logger) Fatal(v ...interface{}) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	os.Stderr.WriteString("[SYNQLY] " + timestamp + " ")
	os.Stderr.WriteString(fmt.Sprintln(v...))
	os.Exit(1)
}
