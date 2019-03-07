package avutil

// #include <libavutil/log.h>
import "C"

// Class ..
type Class C.AVClass

// ClassCategory ..
type ClassCategory C.AVClassCategory

// OptionRanges ..
type OptionRanges C.struct_AVOptionRanges

// LogLevel .
type LogLevel int

const (
	// LogLevelQuiet ..
	LogLevelQuiet = LogLevel(int(C.AV_LOG_QUIET))
	// LogLevelPanic ..
	LogLevelPanic = LogLevel(int(C.AV_LOG_PANIC))
	// LogLevelFatal ..
	LogLevelFatal = LogLevel(int(C.AV_LOG_FATAL))
	// LogLevelError ..
	LogLevelError = LogLevel(int(C.AV_LOG_ERROR))
	// LogLevelWarning ..
	LogLevelWarning = LogLevel(int(C.AV_LOG_WARNING))
	// LogLevelInfo ..
	LogLevelInfo = LogLevel(int(C.AV_LOG_INFO))
	// LogLevelVerbose ..
	LogLevelVerbose = LogLevel(int(C.AV_LOG_VERBOSE))
	// LogLevelDebug ..
	LogLevelDebug = LogLevel(int(C.AV_LOG_DEBUG))
	// LogLevelTrace ..
	LogLevelTrace = LogLevel(int(C.AV_LOG_TRACE))
)

// GetLogLevel ..
func GetLogLevel() int {
	return int(C.av_log_get_level())
}

// SetLogLevel ..
func SetLogLevel(level int) {
	C.av_log_set_level(C.int(level))
}
