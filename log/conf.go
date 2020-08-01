package log

import "fmt"

import "time"

type LoggerOption func(l *Config)

func DefaultLogger(l *Config) {
	if l.LogName == "" {
		l.LogName = "log"
	}
	if l.ErrLogName == "" {
		l.ErrLogName = "err"
	}
	if l.PostFix == nil {
		l.PostFix = []PostFixOption{LoggerPostFixTime}
	}
	if l.MaxSize == 0 {
		// 默认大小16k
		l.MaxSize = 1024 * 16
	}
}

func LoggerOutput(postfixOption ...PostFixOption) LoggerOption {
	return func(l *Config) {
		l.PostFix = postfixOption
	}
}

type Config struct {
	LogDir     string
	LogName    string
	ErrLogName string
	MaxSize    uint32
	PostFix    []PostFixOption
}

type PostFixOption func(s string, l *Logger) string

func LoggerPostFixTime(s string, l *Logger) string {
	t := time.Now().Format("2006-01-02-15:04:05")
	return fmt.Sprintf("%v_%v", s, t)
}
