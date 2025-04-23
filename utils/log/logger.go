package log

import (
	"fmt"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gon.cl/algorithms/utils"
)

const __FALLBACK_LOG_LEVEL__ = "INFO"

type Log struct {
	log *zap.Logger
}

type Fields map[string]string

var instance *Log

func getLogLevel(logLevel string) zapcore.Level {
	var zapLogLevel zapcore.Level

	logLevel = strings.ToUpper(logLevel)
	logLevel = strings.TrimSpace(logLevel)

	fmt.Printf("Search for LOG_LEVEL availability: [%s] \n", logLevel)

	switch logLevel {
	case "INFO":
		zapLogLevel = zap.InfoLevel
	case "DEBUG":
		zapLogLevel = zap.DebugLevel
	case "WARN":
		zapLogLevel = zap.WarnLevel
	case "ERROR":
		zapLogLevel = zap.ErrorLevel
	default:
		zapLogLevel = zap.InfoLevel
	}
	return zapLogLevel
}

func init() {
	level := getLogLevel(utils.GetEnv("LOG_LEVEL", __FALLBACK_LOG_LEVEL__))

	log := initLoggerZap()

	fmt.Printf("Log loaded successfully with level: [%s] \n", level.CapitalString())

	instance = &Log{log: log}
}

func initLoggerZap() *zap.Logger {
	cfg := zap.Config{
		Encoding:         "json",
		DisableCaller:    true,
		Level:            zap.NewAtomicLevelAt(getLogLevel(utils.GetEnv("LOG_LEVEL", __FALLBACK_LOG_LEVEL__))),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stdout"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "msg",
			LevelKey:     "level",
			EncodeLevel:  zapcore.CapitalLevelEncoder,
			TimeKey:      "time",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	log, _ := cfg.Build()
	zap.AddCallerSkip(1)
	zap.ReplaceGlobals(log)
	return log
}

func Field(key string, value string) Fields { return instance.Field(key, value) }
func (logger *Log) Field(key string, value string) Fields {
	return Fields{
		key:   key,
		value: value,
	}
}
func WithFields(fields Fields) *Log { return instance.WithFields(fields) }
func (logger *Log) WithFields(fields Fields) *Log {
	logCustom := &Log{
		log: initLoggerZap(),
	}
	for k, v := range fields {
		logCustom.log = logCustom.log.With(zap.String(k, v))
	}
	return logCustom
}

func WithError(err error) *Log { return instance.WithError(err) }
func (logger *Log) WithError(err error) *Log {
	logCustom := &Log{
		log: initLoggerZap(),
	}
	logCustom.log = logCustom.log.With(zap.String("error", err.Error()))
	return logCustom
}

func WithField(key string, value interface{}) *Log { return instance.WithField(key, value) }
func (logger *Log) WithField(key string, value interface{}) *Log {
	logCustom := &Log{
		log: initLoggerZap(),
	}
	logCustom.log = logCustom.log.With(zap.Any(key, value))
	return logCustom
}

func Info(message string, args ...interface{}) { instance.Info(message, args...) }
func (logger *Log) Info(message string, args ...interface{}) {
	logger.log.Info(fmt.Sprintf(message, args...))
}

func Error(message string, args ...interface{}) { instance.Error(message, args...) }
func (logger *Log) Error(message string, args ...interface{}) {
	logger.log.Error(fmt.Sprintf(message, args...))
}

func Debug(message string, args ...interface{}) { instance.Debug(message, args...) }
func (logger *Log) Debug(message string, args ...interface{}) {
	logger.log.Debug(fmt.Sprintf(message, args...))
}

func Warn(message string, args ...interface{}) { instance.Warn(message, args...) }
func (logger *Log) Warn(message string, args ...interface{}) {
	logger.log.Warn(fmt.Sprintf(message, args...))
}

func Fatal(message string, args ...interface{}) { instance.Fatal(message, args...) }
func (logger *Log) Fatal(message string, args ...interface{}) {
	logger.log.Fatal(fmt.Sprintf(message, args...))
}

func Printf(message string, args ...interface{}) {
	instance.Printf(message, args...)
}
func (logger *Log) Printf(message string, args ...interface{}) {
	logger.log.Info(fmt.Sprintf(message, args...))
}
