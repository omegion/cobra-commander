package cmd

import (
	"strings"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

const (
	JsonFormat = "json"
	TextFormat = "text"
)

// InitLogger inits logrus logger.
func InitLogger(logLevel, logFormat string) {
	log.SetFormatter(CreateFormatter(logFormat))
	log.SetLevel(createLogLevel(logLevel))
}

// CreateFormatter create logrus formatter by string
func CreateFormatter(logFormat string) logrus.Formatter {
	var formatType logrus.Formatter

	switch strings.ToLower(logFormat) {
	case JsonFormat:
		formatType = &logrus.JSONFormatter{}
	case TextFormat:
		formatType = &logrus.TextFormatter{
			TimestampFormat: "02-01-2006 15:04:05",
			FullTimestamp:   true,
			ForceColors:     true,
		}
	default:
		formatType = &logrus.TextFormatter{}
	}

	return formatType
}

func createLogLevel(logLevel string) log.Level {
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		level = logrus.InfoLevel
	}
	return level
}
