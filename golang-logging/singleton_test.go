package belajar_golang_logger

import (
	"testing"

	"github.com/sirupsen/logrus"
)

type FieldLogging struct {
	User     string
	Message  string
	Function string
}

func TestSigleton(t *testing.T) {
	// singleton itu menggunakan object baru jangan langusng menggunakan logrus example
	funcName := "[Hello Test]"
	// loggingField := FieldLogging{
	// 	User:     "HAHA",
	// 	Message:  "Cek Logging Field with struct",
	// 	Function: funcName,
	// }

	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{"FuncName": funcName}).Info("logging")
}
