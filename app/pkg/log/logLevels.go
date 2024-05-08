package log

type logLevel string

// Уровни логов
const (
	warningLevel = logLevel("WARN ")
	errorLevel   = logLevel("ERROR")
	infoLevel    = logLevel("INFO ")
	fatalLevel   = logLevel("FATAL")
	debugLevel   = logLevel("DEBUG")
)
