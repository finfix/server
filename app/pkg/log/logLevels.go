package log

type logLevel string

// Уровни логов
const (
	warningLevel = logLevel("warning")
	errorLevel   = logLevel("error")
	infoLevel    = logLevel("info")
	fatalLevel   = logLevel("fatal")
	debugLevel   = logLevel("debug")
	panicLevel   = logLevel("panic")
)
