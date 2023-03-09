package logger

import (
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

func Init(level string) {
	var logLevel log.Level

	switch strings.ToLower(level) {
	case "error":
		logLevel = log.ErrorLevel
	case "warn":
		logLevel = log.WarnLevel
	case "info":
		logLevel = log.InfoLevel
	case "debug":
		logLevel = log.DebugLevel
	default:
		logLevel = log.InfoLevel
	}

	// Output to stdout instead of the default stderr
	log.SetOutput(os.Stdout)
	// Only log the debug severity or above
	log.SetLevel(logLevel)
	// logrus show line number
	//log.SetReportCaller(true)

	log.SetFormatter(&log.TextFormatter{
		DisableColors:   false,
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
		ForceColors:     true,
	})
}
