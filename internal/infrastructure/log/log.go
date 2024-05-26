// Package log contains logging setup and utilities.
package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

// InitLog is the function used to initiate the log.
func InitLog() {
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(new(logrus.TextFormatter))
	logrus.SetReportCaller(true)
	logrus.Info("Dealls Technical Test - Dating Service")
}
