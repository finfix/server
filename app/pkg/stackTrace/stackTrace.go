package stackTrace

import (
	"fmt"
	"runtime"
	"strings"
)

const (
	ThisCall = iota + 1
	SkipThisCall
	SkipPreviousCaller
	Skip2PreviousCallers
	Skip3PreviousCallers
)

type stackTracer struct {
	serviceName string
}

var stackTracerInstance = &stackTracer{
	serviceName: "",
}

func Init(serviceName string) {
	stackTracerInstance = &stackTracer{
		serviceName: serviceName,
	}
}

func GetStackTrace(skip int) []string {
	var pcs [32]uintptr
	n := runtime.Callers(0, pcs[:])
	var path []string
	for i := skip; i < n; i++ {
		_, file, line, _ := runtime.Caller(i)
		if strings.Contains(file, stackTracerInstance.serviceName) {
			path = append(path, fmt.Sprintf("%s:%d", file, line))
		}
	}
	return path
}
