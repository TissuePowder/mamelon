package logger

import (
	"fmt"
	"log/slog"
	"os"
	"runtime"
	"strings"
	"sync"
)

type Logger struct {
	SLogger *slog.Logger
	mu      sync.Mutex
}

func New() *Logger {
	return &Logger{
		SLogger: slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})),
	}
}

func (l *Logger) Error(msg string, args ...any) {
	l.SLogger.Error(msg)

	stack := make([]uintptr, 10)
	// skip 2 frames
	length := runtime.Callers(2, stack[:])
	frames := runtime.CallersFrames(stack[:length])

	for {
		frame, more := frames.Next()
		if strings.Contains(frame.File, "mamelon") {
			fmt.Printf("File: %s, Function: %s, Line: %d\n", frame.File, frame.Function, frame.Line)
		}
		if !more {
			break
		}
	}
}

func (l *Logger) Info(msg string, args ...any) {
	l.SLogger.Info(msg)
}
