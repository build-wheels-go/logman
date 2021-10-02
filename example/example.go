package main

import (
	"logman"
	"os"
)

func main() {
	logman.Info("std log")
	//修改日志级别
	logman.SetOptions(logman.WithLevel(logman.DebugLevel))
	logman.Debug("change std log level to debug")
	//修改日志格式
	logman.SetOptions(logman.WithFormatter(&logman.JsonFormatter{}))
	logman.Debug("log in json format")
	logman.Info("another log in json format")
	logman.Infof("file name is %s, size is %d Kb", "test.log", 1024)
	//输出到文件
	fd, err := os.OpenFile("test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logman.Fatal("create file test.log failed")
	}
	defer fd.Close()

	logger := logman.New(
		logman.WithLevel(logman.InfoLevel),
		logman.WithOutput(fd),
		logman.WithFormatter(&logman.JsonFormatter{}),
	)
	logger.Info("custom log with json formatter")
}
