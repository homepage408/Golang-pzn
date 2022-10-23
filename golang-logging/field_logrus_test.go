package belajar_golang_logger

import (
	"testing"

	"github.com/sirupsen/logrus"
)

type FormatField struct {
	Name string
	User string
}

func TestFields(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"name": "Mackbook Pro M1",
	}).Info("HELLLOO")
}
