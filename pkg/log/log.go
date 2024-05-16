package log

import (
	"errors"
	"fmt"
	"log"
	"os"
)

type Level int

const (
	DEBUG Level = iota
	INFO
	WARNING
	CRITICAL
	FATAL
)

var levels = []string{"DBG", "INFO", "WARN", "CRIT"}
var Logger *log.Logger
var GameLogger *log.Logger

var MinLevel = INFO

func init() {
	Logger = log.New(os.Stdout, "goker > ", log.Ldate|log.Ltime|log.Lmsgprefix)
	GameLogger = log.New(os.Stdout, "", 0)

	if MinLevel == DEBUG {
		Logger.SetFlags(Logger.Flags() | log.Lshortfile)
		GameLogger.SetFlags(Logger.Flags() | log.Lshortfile)
	}
}

func Debug(msg string) {
	if MinLevel <= DEBUG {
		Logger.Print("[" + levels[DEBUG] + "]  " + msg)
	}
}
func Debugf(format string, values ...interface{}) {
	if MinLevel <= DEBUG {
		Logger.Printf("["+levels[DEBUG]+"]  "+format, values...)
	}
}
func Info(msg string) {
	if MinLevel <= INFO {
		Logger.Print("[" + levels[INFO] + "]  " + msg)
	}
}
func Infof(format string, values ...interface{}) {
	if MinLevel <= INFO {
		Logger.Printf("["+levels[INFO]+"]  "+format, values...)
	}
}
func Warn(msg string) {
	if MinLevel <= WARNING {
		Logger.Print("[" + levels[WARNING] + "]  " + msg)
	}
}
func Warnf(format string, values ...interface{}) {
	if MinLevel <= WARNING {
		Logger.Printf("["+levels[WARNING]+"]  "+format, values...)
	}
}
func Critical(msg string) {
	Logger.Print("[" + levels[CRITICAL] + "]  " + msg)
}
func Criticalf(format string, values ...interface{}) {
	Logger.Printf("["+levels[CRITICAL]+"]  "+format, values...)
}
func Fatal(msg string) {
	Logger.Fatal("[" + levels[CRITICAL] + "]  " + msg)
}
func Fatalf(format string, values ...interface{}) {
	Logger.Fatalf("["+levels[CRITICAL]+"]  "+format, values...)
}
func Error(msg string) error {
	return errors.New(msg)
}
func Errorf(format string, values ...interface{}) error {
	//Is there an easy way not to use fmt here ? only use ...
	return fmt.Errorf(format, values...)
}
func GLog(msg string) {
	GameLogger.Print(msg)
}
func GLogf(format string, values ...interface{}) {
	GameLogger.Printf(format, values...)
}
