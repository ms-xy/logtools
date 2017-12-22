package logtools

import (
	"fmt"
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"os"
	"path/filepath"
	"runtime"
)

var (
	initialized = false
)

const (
	DebugLevel = logrus.DebugLevel
	InfoLevel  = logrus.InfoLevel
	WarnLevel  = logrus.WarnLevel
	ErrorLevel = logrus.ErrorLevel
	FatalLevel = logrus.FatalLevel
	PanicLevel = logrus.PanicLevel
)

type Fields logrus.Fields

// initialization and configuration

func Initialize() {
	formatter := new(prefixed.TextFormatter)
	formatter.FullTimestamp = true
	logrus.SetFormatter(formatter)
	logrus.SetOutput(os.Stdout)
	SetLevel(DebugLevel)
	initialized = true
}

func SetLevel(level logrus.Level) {
	logrus.SetLevel(level)
}

// helper function to determine the caller

func getCaller() string {
	if _, path, line, ok := runtime.Caller(3); ok {
		_, file := filepath.Split(path)
		return fmt.Sprintf("%s:%d", file, line)
	}
	return "<???>"
}

// internal wrapper

func Prefix() *logrus.Entry {
	if !initialized {
		Initialize()
	}
	return logrus.WithFields(logrus.Fields{
		"prefix": getCaller()})
}

// prepare function basically

func WithFields(fields Fields) *logrus.Entry {
	return Prefix().WithFields(map[string]interface{}(fields))
}

// convenience functions

func Debug(i ...interface{}) {
	Prefix().Debug(i...)
}
func Debugf(format string, i ...interface{}) {
	Prefix().Debugf(format, i...)
}

func Log(i ...interface{}) {
	Prefix().Info(i...)
}
func Logf(format string, i ...interface{}) {
	Prefix().Infof(format, i...)
}

func Info(i ...interface{}) {
	Prefix().Info(i...)
}
func Infof(format string, i ...interface{}) {
	Prefix().Infof(format, i...)
}

func Warn(i ...interface{}) {
	Prefix().Warn(i...)
}
func Warnf(format string, i ...interface{}) {
	Prefix().Warnf(format, i...)
}

func Error(i ...interface{}) {
	Prefix().Error(i...)
}
func Errorf(format string, i ...interface{}) {
	Prefix().Errorf(format, i...)
}

func Fatal(i ...interface{}) {
	Prefix().Fatal(i...)
}
func Fatalf(format string, i ...interface{}) {
	Prefix().Fatalf(format, i...)
}

func Panic(i ...interface{}) {
	Prefix().Panic(i...)
}
func Panicf(format string, i ...interface{}) {
	Prefix().Panicf(format, i...)
}
