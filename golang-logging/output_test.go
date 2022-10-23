package belajar_golang_logger

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestOutput(t *testing.T) {
	logger := logrus.New()

	file, _ := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger.SetOutput(file)
	logger.WithFields(logrus.Fields{"animal": "HAHAH"}).Info("HELLOO PAKE FIELD")

	logger.Info("Hello world!")
	logger.Warn("Hello world!")
	logger.Error("Hello world!")

}
