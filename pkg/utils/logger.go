package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
)

type Logger struct {
	Level string
}

func NewLogger(level string) *Logger {
	return &Logger{Level: level}
}

func (l *Logger) Info(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	log.Printf("%s %s", color.CyanString("[INFO]"), msg)
}

func (l *Logger) Error(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	log.Printf("%s %s", color.RedString("[ERROR]"), msg)
}

func (l *Logger) Warn(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	log.Printf("%s %s", color.YellowString("[WARN]"), msg)
}

func (l *Logger) Success(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	log.Printf("%s %s", color.GreenString("[SUCCESS]"), msg)
}

func (l *Logger) Fatal(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	fmt.Fprintf(os.Stderr, "%s %s\n", color.RedString("[FATAL]"), msg)
	os.Exit(1)
}
