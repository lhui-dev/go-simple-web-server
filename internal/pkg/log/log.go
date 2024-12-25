package log

import (
	"github.com/leedev/go-simple-web-server/internal/configuration"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"time"
)

var logger *logrus.Logger
var logToFile *logrus.Logger

// 日志文件名
var loggerFile string

func setLogFile(file string) {
	loggerFile = file
}

// 初始化
func init() {
	filePath := filepath.Join(configuration.Config.Log.Path, configuration.Config.Log.FileName)
	setLogFile(filePath)
}

const (
	fileMode  = "file"
	mqMode    = "mq"
	esMode    = "es"
	mongoMode = "mongo"
)

// Log 日志方法调用
// Log 日志方法调用
func Log() *logrus.Logger {
	switch configuration.Config.Log.Mode {
	case fileMode:
		return doLogToFile()
	case mqMode:
		return doLogToMQ()
	case esMode:
		return doLogToElasticsearch()
	case mongoMode:
		return doLogToMongodb()
	default:
		return doLogToConsole()
	}
}

const timeFormat = "2006-01-02 15:04:05"

// 日志记录到文件方法
func doLogToFile() *logrus.Logger {
	if logToFile == nil {
		logToFile = logrus.New()

		// 显示行号
		logToFile.SetReportCaller(false)

		// 日志级别
		logToFile.SetLevel(logrus.DebugLevel)

		// 返回写日志对象logWriter
		logWriter, _ := rotatelogs.New(
			// 分割后的文件名称
			loggerFile+"_%Y%m%d.log",
			// 设置最大保存时间
			rotatelogs.WithMaxAge(time.Hour*24*30),
			// 设置日志切割时间间隔（1天）
			rotatelogs.WithRotationTime(time.Hour*24),
		)

		writeMap := lfshook.WriterMap{
			logrus.InfoLevel:  logWriter,
			logrus.FatalLevel: logWriter,
			logrus.DebugLevel: logWriter,
			logrus.WarnLevel:  logWriter,
			logrus.ErrorLevel: logWriter,
			logrus.PanicLevel: logWriter,
		}

		// 设置时间格式
		lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{TimestampFormat: timeFormat})

		// 新增 Hook
		logToFile.AddHook(lfHook)
	}
	return logToFile
}

// 日志记录到控制台方法
func doLogToConsole() *logrus.Logger {
	//  实现记录到控制台逻辑
	if logger == nil {
		logger = logrus.New()
		logger.Out = os.Stdout // 日志标准输出
		logger.Formatter = &logrus.JSONFormatter{TimestampFormat: timeFormat}
		logger.SetLevel(logrus.DebugLevel)
	}
	return logger
}

// 日志记录到消息队列MQ方法
func doLogToMQ() *logrus.Logger {
	// TODO 实现记录到消息队列逻辑
	return logrus.New()
}

// 日志记录到ES方法
func doLogToElasticsearch() *logrus.Logger {
	// TODO 实现记录到ES逻辑
	return logrus.New()
}

// 日志记录到MongoDB方法
func doLogToMongodb() *logrus.Logger {
	// TODO 实现记录到MongoDB逻辑
	return logrus.New()
}
