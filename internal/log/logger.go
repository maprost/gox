package log

import "fmt"

type logger struct {
	level Level
}

var log logger

func InitLogger(level Level) {
	log = logger{
		level: level,
	}
}

func Debug(args ...interface{}) {
	Print(LevelDebug, args...)
}

func Info(args ...interface{}) {
	Print(LevelInfo, args...)
}

func Warn(args ...interface{}) {
	Print(LevelWarn, args...)
}

func Print(logLevel Level, args ...interface{}) {
	if log.level <= logLevel {
		fmt.Println(args...)
	}
}
