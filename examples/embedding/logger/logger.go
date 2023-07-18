package logger

import (
	"fmt"
	"practice-go-examples/examples/embedding/guards"
)

type Logger struct {
	*guards.Guard
}

func NewLogger(guard *guards.Guard) *Logger {
	return &Logger{Guard: guard}
}

func (mylogger *Logger) LogInfo(message string) {
	if mylogger.Guard.HasPermissions() {
		fmt.Println(message)
	} else {
		fmt.Println("You don't have permission to log")
	}
}
