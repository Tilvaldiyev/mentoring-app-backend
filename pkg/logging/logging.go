package logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"runtime"
)

// writerHook - hook for writing in multiple outputs
type writerHook struct {
	Writer    []io.Writer
	LogLevels []logrus.Level
}

// Fire - fire appropriate hooks for a log entry
func (hook *writerHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return fmt.Errorf("entry to string err: %w", err)
	}

	for _, w := range hook.Writer {
		w.Write([]byte(line))
	}

	return nil
}

// Levels - return all log levels
func (hook *writerHook) Levels() []logrus.Level {
	return hook.LogLevels
}

// Init - initializing logger
func Init() (*logrus.Entry, error) {
	l := logrus.New()
	l.SetReportCaller(true)
	l.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			filename := path.Base(frame.File)
			return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", filename, frame.Line)
		},
		DisableColors: false,
		FullTimestamp: true,
	}

	err := os.MkdirAll("logs", 0777)
	if err != nil {
		return nil, fmt.Errorf("can't create a folder err: %w", err)
	}

	allFile, err := os.OpenFile("logs/all.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	if err != nil {
		return nil, fmt.Errorf("openning file err: %w", err)
	}

	l.SetOutput(io.Discard)

	l.AddHook(&writerHook{
		Writer: []io.Writer{os.Stdout, allFile},
		LogLevels: logrus.AllLevels,
	})

	l.SetLevel(logrus.TraceLevel)

	return logrus.NewEntry(l), nil
}
