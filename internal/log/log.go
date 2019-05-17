package log

import (
	"path"
	"sync"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
)

var defaultLogger *logrus.Logger
var once sync.Once

func Init(logLevel logrus.Level) {
	defaultLogger = logrus.New()
	// defaultLogger.SetReportCaller(true)

	// defaultLogger.SetLevel(logrus.DebugLevel)
	defaultLogger.SetLevel(logLevel)

	baseLogPath := path.Join("./logs", "log.out")
	writer, err := rotatelogs.New(
		baseLogPath+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(baseLogPath),   // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(time.Hour),       // 文件最大保存时间
		rotatelogs.WithRotationTime(time.Hour), // 日志切割时间间隔
	)
	if err != nil {
		panic(err)
	}
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer, // 为不同级别设置不同的输出目的
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, defaultLogger.Formatter)
	defaultLogger.AddHook(lfHook)
}

func Default() *logrus.Logger {
	if defaultLogger == nil {
		once.Do(func() {
			Init(logrus.WarnLevel)
		})
	}
	return defaultLogger
}
