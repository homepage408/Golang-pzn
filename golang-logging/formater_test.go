package belajar_golang_logger

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestFormater(t *testing.T) {
	logger := logrus.New()
	file, _ := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger.Out = os.Stdout
	logger.SetOutput(file)
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.Info("Testing formater")

}
