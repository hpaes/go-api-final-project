package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Fields map[string]interface{}

type LoggerService interface {
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(format string, args ...interface{})
	Fatal(format string, args ...interface{})
	WithFields(fields Fields) LoggerService
}

type loggerService struct {
	logger *zap.SugaredLogger
}

func NewLoggerService() LoggerService {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")

	l, err := config.Build()
	if err != nil {
		panic(err)
	}
	sugar := l.Sugar()
	defer l.Sync()

	return &loggerService{
		logger: sugar,
	}
}

func (l *loggerService) Info(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

func (l *loggerService) Warn(format string, args ...interface{}) {
	l.logger.Warnf(format, args...)
}

func (l *loggerService) Error(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

func (l *loggerService) Fatal(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}

func (l *loggerService) WithFields(fields Fields) LoggerService {
	var f = make([]interface{}, 0)
	for k, v := range fields {
		f = append(f, k)
		f = append(f, v)
	}
	newLogger := l.logger.With(f...)
	return &loggerService{
		logger: newLogger,
	}
}
