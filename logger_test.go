package cmd

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestCreateFormatter(t *testing.T) {
	t.Run("log format is json", func(t *testing.T) {
		result := CreateFormatter("json")
		assert.Equal(t, &logrus.JSONFormatter{}, result)
	})

	t.Run("log format is text", func(t *testing.T) {
		result := CreateFormatter("text")
		assert.Equal(t, &logrus.TextFormatter{
			TimestampFormat: "02-01-2006 15:04:05",
			FullTimestamp:   true,
			ForceColors:     true,
		}, result)
	})

	t.Run("log format is not json or text", func(t *testing.T) {
		result := CreateFormatter("xml")
		assert.Equal(t, &logrus.TextFormatter{}, result)
	})
}
