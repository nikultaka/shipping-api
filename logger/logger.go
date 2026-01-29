package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

type LogLevel string

const (
	INFO  LogLevel = "INFO"
	ERROR LogLevel = "ERROR"
	WARN  LogLevel = "WARN"
	DEBUG LogLevel = "DEBUG"
)

var logFile *os.File

func Init() error {
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		os.Mkdir("logs", 0755)
	}

	filename := fmt.Sprintf("logs/shipping-%s.log", time.Now().Format("2006-01-02"))
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	logFile = file
	log.SetOutput(file)
	return nil
}

func Info(message string, fields map[string]interface{}) {
	writeLog(INFO, message, fields)
}

func Error(message string, fields map[string]interface{}) {
	writeLog(ERROR, message, fields)
}

func Warn(message string, fields map[string]interface{}) {
	writeLog(WARN, message, fields)
}

func writeLog(level LogLevel, message string, fields map[string]interface{}) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	logEntry := fmt.Sprintf("[%s] %s - %s", timestamp, level, message)

	if len(fields) > 0 {
		for key, value := range fields {
			logEntry += fmt.Sprintf(" | %s=%v", key, value)
		}
	}

	fmt.Println(logEntry)

	if logFile != nil {
		fmt.Fprintln(logFile, logEntry)
	}
}

func Close() {
	if logFile != nil {
		logFile.Close()
	}
}
