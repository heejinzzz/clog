package clog

import (
	"fmt"
	"io"
	"log"
)

// store all loggers created
var loggers []*Logger

type Logger struct {
	writer        io.Writer
	prefix        string
	prefixPattern string
	flags         int
	level         Level
	enableColor   bool

	debugLogger *log.Logger
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
	fatalLogger *log.Logger
}

type Level uint8

// log levels
const (
	DebugLevel Level = 0
	InfoLevel  Level = 1
	WarnLevel  Level = 2
	ErrorLevel Level = 3
	FatalLevel Level = 4
	Off        Level = 5
)

// default patterns
var (
	defaultPattern             = fmt.Sprintf("%c[%d;%%dm%%-8s%c[0m", 0x1B, 1, 0x1B)
	defaultPatternDisableColor = "%-8s"
	defaultPrefixPattern       = fmt.Sprintf("%c[%d;%dm%%s%c[0m", 0x1B, 1, prefixColor, 0x1B)
)

// default flags
var defaultFlags = log.LstdFlags

// default colors
const (
	debugColor = 32 // green
	infoColor  = 34 // blue
	warnColor  = 33 // yellow
	errorColor = 35 // purple
	fatalColor = 31 // red

	prefixColor = 36 // green blue
)

// New create a new logger with output writer and prefix
func New(writer io.Writer, prefix string) *Logger {
	prefixPattern := fmt.Sprintf(defaultPrefixPattern, prefix)
	debugLogger := log.New(writer, prefixPattern+fmt.Sprintf(defaultPattern, debugColor, "[DEBUG]"), defaultFlags)
	infoLogger := log.New(writer, prefixPattern+fmt.Sprintf(defaultPattern, infoColor, "[INFO]"), defaultFlags)
	warnLogger := log.New(writer, prefixPattern+fmt.Sprintf(defaultPattern, warnColor, "[WARN]"), defaultFlags)
	errorLogger := log.New(writer, prefixPattern+fmt.Sprintf(defaultPattern, errorColor, "[ERROR]"), defaultFlags)
	fatalLogger := log.New(writer, prefixPattern+fmt.Sprintf(defaultPattern, fatalColor, "[FATAL]"), defaultFlags)

	logger := &Logger{
		writer:        writer,
		prefix:        prefix,
		prefixPattern: prefixPattern,
		flags:         defaultFlags,
		level:         DebugLevel,
		enableColor:   true,

		debugLogger: debugLogger,
		infoLogger:  infoLogger,
		warnLogger:  warnLogger,
		errorLogger: errorLogger,
		fatalLogger: fatalLogger,
	}
	loggers = append(loggers, logger)
	return logger
}

// Writer return logger's output writer
func (logger *Logger) Writer() io.Writer {
	return logger.writer
}

// Prefix return logger's prefix
func (logger *Logger) Prefix() string {
	return logger.prefix
}

// Flags return logger's flags
func (logger *Logger) Flags() int {
	return logger.flags
}

// Level return logger's level
func (logger *Logger) Level() Level {
	return logger.level
}

// SetWriter set output writer of the logger
func (logger *Logger) SetWriter(writer io.Writer) {
	logger.writer = writer

	logger.debugLogger.SetOutput(writer)
	logger.infoLogger.SetOutput(writer)
	logger.warnLogger.SetOutput(writer)
	logger.errorLogger.SetOutput(writer)
	logger.fatalLogger.SetOutput(writer)
}

// SetWriterAll set all loggers' output writer
func SetWriterAll(writer io.Writer) {
	for _, logger := range loggers {
		logger.SetWriter(writer)
	}
}

// SetPrefix set prefix of the logger
func (logger *Logger) SetPrefix(prefix string) {
	logger.prefix = prefix
	logger.prefixPattern = fmt.Sprintf(defaultPrefixPattern, prefix)

	if logger.enableColor {
		logger.debugLogger.SetPrefix(logger.prefixPattern + fmt.Sprintf(defaultPattern, debugColor, "[DEBUG]"))
		logger.infoLogger.SetPrefix(logger.prefixPattern + fmt.Sprintf(defaultPattern, infoColor, "[INFO]"))
		logger.warnLogger.SetPrefix(logger.prefixPattern + fmt.Sprintf(defaultPattern, warnColor, "[WARN]"))
		logger.errorLogger.SetPrefix(logger.prefixPattern + fmt.Sprintf(defaultPattern, errorColor, "[ERROR]"))
		logger.fatalLogger.SetPrefix(logger.prefixPattern + fmt.Sprintf(defaultPattern, fatalColor, "[FATAL]"))
	} else {
		logger.debugLogger.SetPrefix(logger.prefix + fmt.Sprintf(defaultPatternDisableColor, "[DEBUG]"))
		logger.infoLogger.SetPrefix(logger.prefix + fmt.Sprintf(defaultPatternDisableColor, "[INFO]"))
		logger.warnLogger.SetPrefix(logger.prefix + fmt.Sprintf(defaultPatternDisableColor, "[WARN]"))
		logger.errorLogger.SetPrefix(logger.prefix + fmt.Sprintf(defaultPatternDisableColor, "[ERROR]"))
		logger.fatalLogger.SetPrefix(logger.prefix + fmt.Sprintf(defaultPatternDisableColor, "[FATAL]"))
	}
}

// SetPrefixAll set all loggers' prefix
func SetPrefixAll(prefix string) {
	for _, logger := range loggers {
		logger.SetPrefix(prefix)
	}
}

// SetFlags set flags of the logger
func (logger *Logger) SetFlags(flags int) {
	logger.flags = flags

	logger.debugLogger.SetFlags(flags)
	logger.infoLogger.SetFlags(flags)
	logger.warnLogger.SetFlags(flags)
	logger.errorLogger.SetFlags(flags)
	logger.fatalLogger.SetFlags(flags)
}

// SetFlagsAll set all loggers' flags
func SetFlagsAll(flags int) {
	for _, logger := range loggers {
		logger.SetFlags(flags)
	}
}

// SetLevel set log level of the logger
func (logger *Logger) SetLevel(level Level) {
	logger.level = level
}

// SetLevelAll set all loggers' level
func SetLevelAll(level Level) {
	for _, logger := range loggers {
		logger.SetLevel(level)
	}
}

// Debug print debug log
func (logger *Logger) Debug(v ...any) {
	if logger.level > DebugLevel {
		return
	}
	logger.debugLogger.Println(v...)
}

// Debugf print debug log with a format
func (logger *Logger) Debugf(format string, v ...any) {
	if logger.level > DebugLevel {
		return
	}
	logger.debugLogger.Printf(format, v...)
}

// Info print info log
func (logger *Logger) Info(v ...any) {
	if logger.level > InfoLevel {
		return
	}
	logger.infoLogger.Println(v...)
}

// Infof print info log with a format
func (logger *Logger) Infof(format string, v ...any) {
	if logger.level > InfoLevel {
		return
	}
	logger.infoLogger.Printf(format, v...)
}

// Warn print warn log
func (logger *Logger) Warn(v ...any) {
	if logger.level > WarnLevel {
		return
	}
	logger.warnLogger.Println(v...)
}

// Warnf print warn log with a format
func (logger *Logger) Warnf(format string, v ...any) {
	if logger.level > WarnLevel {
		return
	}
	logger.warnLogger.Printf(format, v...)
}

// Error print error log
func (logger *Logger) Error(v ...any) {
	if logger.level > ErrorLevel {
		return
	}
	logger.errorLogger.Println(v...)
}

// Errorf print error log with a format
func (logger *Logger) Errorf(format string, v ...any) {
	if logger.level > ErrorLevel {
		return
	}
	logger.errorLogger.Printf(format, v...)
}

// Fatal print fatal log
func (logger *Logger) Fatal(v ...any) {
	if logger.level > FatalLevel {
		return
	}
	logger.fatalLogger.Println(v...)
}

// Fatalf print fatal log with a format
func (logger *Logger) Fatalf(format string, v ...any) {
	if logger.level > FatalLevel {
		return
	}
	logger.fatalLogger.Printf(format, v...)
}

// IsColorEnabled return is the color output function of the logger enabled
func (logger *Logger) IsColorEnabled() bool {
	return logger.enableColor
}

// DisableColor disable the color output function of the logger
func (logger *Logger) DisableColor() {
	logger.enableColor = false
	logger.SetPrefix(logger.prefix)
}

// DisableColorAll disable the color output function of all loggers
func DisableColorAll() {
	for _, logger := range loggers {
		logger.DisableColor()
	}
}

// EnableColor enable the color output function of the logger
func (logger *Logger) EnableColor() {
	logger.enableColor = true
	logger.SetPrefix(logger.prefix)
}

// EnableColorAll enable the color output function of all loggers
func EnableColorAll() {
	for _, logger := range loggers {
		logger.EnableColor()
	}
}
