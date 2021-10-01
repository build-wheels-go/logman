package logman

import (
	"bytes"
	"errors"
	"fmt"
)

const FmtEmptySeparate = ""

type Level uint8

const (
	//debug级别，用于调试代码，生产环境需要关闭
	DebugLevel Level = iota
	//info级别，默认级别，记录有用的信息，
	//辅助排查问题
	InfoLevel
	//warn级别，代码产生不符合预期的结果，
	//但是不会影响程序执行。需要去解决
	WarnLevel
	//Error级别，影响程序正常执行的错误，优先级高
	ErrorLevel
	//Panic级别，程序产生Panic，记录的日记信息
	PanicLevel
	//Fatal级别，程序发生致命错误，需要退出
	FatalLevel
)

var LevelNameMapping = map[Level]string{
	DebugLevel: "DEBUG",
	InfoLevel:  "INFO",
	WarnLevel:  "WARN",
	ErrorLevel: "ERROR",
	PanicLevel: "PANIC",
	FatalLevel: "FATAL",
}

var errUnmarshalNilLevel = errors.New("can`t unmarshal nil *Level")

func (l *Level) unmarshalText(text []byte) bool {
	switch string(text) {
	case "debug", "DEBUG":
		*l = DebugLevel
	case "info", "INFO", "":
		*l = InfoLevel
	case "warn", "WARN":
		*l = WarnLevel
	case "error", "ERROR":
		*l = ErrorLevel
	case "panic", "PANIC":
		*l = PanicLevel
	case "fatal", "FATAL":
		*l = FatalLevel
	default:
		return false
	}
	return true
}

func (l *Level) UnmarshalText(text []byte) error {
	if text == nil {
		return errUnmarshalNilLevel
	}
	if !l.unmarshalText(text) && !l.unmarshalText(bytes.ToLower(text)) {
		return fmt.Errorf("unrecognized level: %q", text)
	}
	return nil
}
