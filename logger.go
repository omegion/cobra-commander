package cmd

import (
	"strings"

	"github.com/sirupsen/logrus"
)

const (
	jsonFormat = "json"
	textFormat = "text"
)

// CreateFormatter create logrus formatter by string.
func CreateFormatter(logFormat string) logrus.Formatter {
	var formatType logrus.Formatter

	switch strings.ToLower(logFormat) {
	case jsonFormat:
		formatType = &logrus.JSONFormatter{}
	case textFormat:
		formatType = &logrus.TextFormatter{
			TimestampFormat: "02-01-2006 15:04:05",
			FullTimestamp:   true,
			DisableColors:   true,
		}
	default:
		formatType = &logrus.TextFormatter{}
	}

	return formatType
}

func createLogLevel(logLevel string) logrus.Level {
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		level = logrus.InfoLevel
	}

	return level
}
