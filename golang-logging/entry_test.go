package belajar_golang_logger

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestEntry(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	entry := logrus.NewEntry(logger)
	entry.WithFields(logrus.Fields{"Username": "Teguh"})
	entry.Info("HEHEHEHE")
}
