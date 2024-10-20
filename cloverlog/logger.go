package cloverlog

import (
	"io"
	"log"
	"os"
)

const (
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	blue   = "\033[34m"
)

type Logger struct {
	debug  *log.Logger
	info   *log.Logger
	warn   *log.Logger
	err    *log.Logger
	writer io.Writer
}
type LoggerOptions struct {
	Prefix     string
	HideTime   bool
	HidePrefix bool
}

func getLog(writer io.Writer, logType string, prefix string, hideTime bool, hidePrefix bool) *log.Logger {
	var flag int
	if !hideTime {
		flag = log.Ltime | log.Ldate
	}
	var prefixText string

	if !hidePrefix {
		switch logType {
		case "debug":
			prefixText = blue + prefix + " [DEBUG] " + reset
		case "info":
			prefixText = green + prefix + " [INFO] " + reset
		case "warn":
			prefixText = yellow + prefix + " [WARN] " + reset
		case "error":
			prefixText = red + prefix + " [ERROR] " + reset
		default:
			prefixText = green + prefix + " [INFO] " + reset
		}
	}
	return log.New(writer, prefixText, flag)
}
func NewLogger(opts LoggerOptions) *Logger {
	writer := io.Writer(os.Stdout)

	var prefix string

	if !opts.HidePrefix {
		prefix = opts.Prefix
	}

	return &Logger{
		writer: writer,
		debug:  getLog(writer, "debug", prefix, opts.HideTime, opts.HidePrefix),
		info:   getLog(writer, "info", prefix, opts.HideTime, opts.HidePrefix),
		warn:   getLog(writer, "warn", prefix, opts.HideTime, opts.HidePrefix),
		err:    getLog(writer, "error", prefix, opts.HideTime, opts.HidePrefix),
	}
}

func (l *Logger) Debug(v ...interface{}) {
	l.debug.Println(v...)
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.debug.Printf(format, v...)
}

func (l *Logger) Info(v ...interface{}) {
	l.info.Println(v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.info.Printf(format, v...)
}

func (l *Logger) Warn(v ...interface{}) {
	l.warn.Println(v...)
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.warn.Printf(format, v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.err.Println(v...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.err.Printf(format, v...)
}
