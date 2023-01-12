package logging

import (
	"errors"

	"github.com/byyjoww/league-mentor/config"
	"github.com/byyjoww/league-mentor/constants"
	"github.com/sirupsen/logrus"
)

func NewLogrusLogger(configs config.Logging) logrus.FieldLogger {
	logger := logrus.New()

	level, err := getLogLevel(configs.Level)
	if err != nil {
		logrus.Panic("failed to get log level")
	}

	logrus.SetLevel(level)
	logrus.SetFormatter(&logrus.JSONFormatter{})

	logger.SetLevel(level)
	logger.SetFormatter(&logrus.JSONFormatter{})

	return logger
}

func getLogLevel(level string) (logrus.Level, error) {
	switch level {
	case constants.LogLevelPanic:
		return logrus.PanicLevel, nil
	case constants.LogLevelFatal:
		return logrus.FatalLevel, nil
	case constants.LogLevelError:
		return logrus.ErrorLevel, nil
	case constants.LogLevelWarn:
		return logrus.WarnLevel, nil
	case constants.LogLevelInfo:
		return logrus.InfoLevel, nil
	case constants.LogLevelDebug:
		return logrus.DebugLevel, nil
	case constants.LogLevelTrace:
		return logrus.TraceLevel, nil
	default:
		return 0, errors.New("unable to determine logging level")
	}
}
