package log

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

type Write struct {
	OutputPath string
	Size       int
	outputFile string
	sizeN      int
}

func NewLogWrite(output string, size int) *Write {
	if output == "" {
		output = "./"
	}
	if size == 0 {
		size = 2
	}
	return &Write{OutputPath: output, Size: size, outputFile: "", sizeN: 0}
}

// 1mb
const mb = 1048576

func (l *Write) Write(p []byte) (n int, err error) {
	if l.outputFile == "" || l.sizeN > l.Size*mb {
		l.sizeN = 0
		l.outputFile = l.generatePath()
	}
	file, err := os.OpenFile(l.outputFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		return 0, err
	}
	n, err = file.Write(p)
	l.sizeN += n
	return
}
func (l *Write) generatePath() string {
	format := time.Now().Format("2006-01-02-15_04_05") + ".log"
	la := strings.TrimRight(l.OutputPath, "/")
	return la + "/" + format
}

type StacktraceHook struct {
}

func (h *StacktraceHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.DebugLevel, logrus.ErrorLevel, logrus.TraceLevel}
}

func (h *StacktraceHook) Fire(e *logrus.Entry) error {
	if v, found := e.Data[logrus.ErrorKey]; found {
		if err, iserr := v.(error); iserr {
			type stackTracer interface {
				StackTrace() errors.StackTrace
			}
			if st, isst := err.(stackTracer); isst {
				stack := fmt.Sprintf("%+v", st.StackTrace())
				e.Data["stacktrace"] = stack
			} else {
				stack := debug.Stack()
				e.Data["stacktrace"] = string(stack)
			}
		}
	}
	return nil
}
