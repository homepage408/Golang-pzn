package belajar_golang_logger

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLevels(t *testing.T) {
	logger := logrus.New()

	logger.Trace("This is Trace")
	logger.Debug("This is Debug")
	logger.Info("This is Info")
	logger.Warn("This is Warn")
	logger.Error("This is Error")

}

func TestLoggingLevels(t *testing.T) {
	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel)

	logger.Trace("This is Trace")
	logger.Debug("This is Debug")
	logger.Info("This is Info")
	logger.Warn("This is Warn")
	logger.Error("This is Error")

}
