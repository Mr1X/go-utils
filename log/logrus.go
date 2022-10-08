package log

import (
	"fmt"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.JSONFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
			return "", fmt.Sprintf("%v:%v", path.Base(f.File), f.Line)
		},
	})
}
