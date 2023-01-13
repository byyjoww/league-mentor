package logging

import (
	app "github.com/byyjoww/league-mentor/services/http"
	"github.com/sirupsen/logrus"
)

type AppLogger = app.Logger

type AppLoggr struct {
	base logrus.FieldLogger
}

// New returns a new Logger implementation based on logrus
func NewAppLogger(logger logrus.FieldLogger) AppLogger {
	return NewAppLoggerWithLogger(logger)
}

// NewAppLoggerWithEntry returns a new Logger implementation based on a provided logrus entry instance
func NewAppLoggerWithEntry(logger *logrus.Entry) AppLogger {
	return &AppLoggr{base: logger}
}

// NewAppLoggerWithLogger returns a new Logger implementation based on a provided logrus instance
func NewAppLoggerWithLogger(logger logrus.FieldLogger) AppLogger {
	return &AppLoggr{base: logger}
}

func (l *AppLoggr) Fatal(format ...interface{}) {
	l.base.Fatal(format...)
}

func (l *AppLoggr) Fatalf(format string, args ...interface{}) {
	l.base.Fatalf(format, args...)
}

func (l *AppLoggr) Fatalln(args ...interface{}) {
	l.base.Fatalln(args...)
}

func (l *AppLoggr) Debug(args ...interface{}) {
	l.base.Debug(args...)
}

func (l *AppLoggr) Debugf(format string, args ...interface{}) {
	l.base.Debugf(format, args...)
}

func (l *AppLoggr) Debugln(args ...interface{}) {
	l.base.Debugln(args...)
}

func (l *AppLoggr) Error(args ...interface{}) {
	l.base.Error(args...)
}

func (l *AppLoggr) Errorf(format string, args ...interface{}) {
	l.base.Errorf(format, args...)
}

func (l *AppLoggr) Errorln(args ...interface{}) {
	l.base.Errorln(args...)
}

func (l *AppLoggr) Info(args ...interface{}) {
	l.base.Info(args...)
}

func (l *AppLoggr) Infof(format string, args ...interface{}) {
	l.base.Infof(format, args...)
}

func (l *AppLoggr) Infoln(args ...interface{}) {
	l.base.Infoln(args...)
}

func (l *AppLoggr) Warn(args ...interface{}) {
	l.base.Warn(args...)
}

func (l *AppLoggr) Warnf(format string, args ...interface{}) {
	l.base.Warnf(format, args...)
}

func (l *AppLoggr) Warnln(args ...interface{}) {
	l.base.Warnln(args...)
}

func (l *AppLoggr) Panic(args ...interface{}) {
	l.base.Panic(args...)
}

func (l *AppLoggr) Panicf(format string, args ...interface{}) {
	l.base.Panicf(format, args...)
}

func (l *AppLoggr) Panicln(args ...interface{}) {
	l.base.Panicln(args...)
}

func (l *AppLoggr) WithFields(fields map[string]interface{}) AppLogger {
	return &AppLoggr{base: l.base.WithFields(fields)}
}

func (l *AppLoggr) WithField(key string, value interface{}) AppLogger {
	return &AppLoggr{base: l.base.WithField(key, value)}
}

func (l *AppLoggr) WithError(err error) AppLogger {
	return &AppLoggr{base: l.base.WithError(err)}
}
