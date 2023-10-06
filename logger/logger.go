package logger

import (
	"github.com/Ruoyiran/exif-reader/formatter"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"runtime"
	"strconv"
)

func initLogger() {
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&formatter.EasyFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "[%lvl%]: %time% - %src% %msg%\n",
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			fileName := path.Base(frame.File) + ":" + strconv.Itoa(frame.Line)
			return "", fileName
		},
	})
	logrus.SetOutput(io.MultiWriter(os.Stdout))
	logrus.SetLevel(logrus.DebugLevel)
}

func init() {
	initLogger()
}
