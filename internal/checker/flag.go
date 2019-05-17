package checker

import (
	"flag"
	"fmt"

	"github.com/Sirupsen/logrus"
)

/*
	// PanicLevel level, highest level of severity. Logs and then calls panic with the
	// message passed to Debug, Info, ...
	PanicLevel Level = iota
	// FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	FatalLevel
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel
	// TraceLevel level. Designates finer-grained informational events than the Debug.
	TraceLevel
*/

var logLevelMap = map[string]logrus.Level{
	"panic": logrus.PanicLevel,
	"fatal": logrus.FatalLevel,
	"error": logrus.ErrorLevel,
	"warn":  logrus.WarnLevel,
	"info":  logrus.InfoLevel,
	"debug": logrus.DebugLevel,
	"trace": logrus.TraceLevel,
}

var envFile string
var reqFile string
var onlinePOIHost string
var debugPOIHost string
var showColor bool
var logLevelStr string
var logLevel logrus.Level
var outputDir string

func EnvFile() string {
	return envFile
}

func ReqFile() string {
	return reqFile
}

func OnlinePOIHost() string {
	return onlinePOIHost
}

func DebugPOIHost() string {
	return debugPOIHost
}

func ShowColor() bool {
	return showColor
}

func LogLevel() logrus.Level {
	if level, ok := logLevelMap[logLevelStr]; ok {
		return level
	}
	return logrus.WarnLevel
}

func OutputDir() string {
	return outputDir
}

func StartString(name, prefix, suffix string) string {
	num := 30
	nameLen := len(name)
	halfNum := (num - nameLen) / 2
	sig := ""
	for i := 0; i < halfNum; i = i + 1 {
		sig += "="
	}
	res := sig + name + sig
	for len(res) < num {
		res += "="
	}
	return prefix + res + suffix
}

func ParseFlag() {
	flag.StringVar(&envFile, "envFile", "./configs/env.json", "env json file")
	flag.StringVar(&reqFile, "reqFile", "./configs/req.json", "req json file")
	flag.StringVar(&onlinePOIHost, "onlinePOIHost", "", "online POI host")
	flag.StringVar(&debugPOIHost, "debugPOIHost", "", "debug POI host")
	flag.BoolVar(&showColor, "showColor", false, "show color")
	flag.StringVar(&logLevelStr, "logLevel", "info", "log level")
	flag.StringVar(&outputDir, "outputDir", "./logs", "output dir")

	flag.Parse()

	fmt.Printf(StartString("CONF", "\n", "\n"))
	fmt.Printf("Flag parsed: %v\n", flag.Parsed())
	fmt.Printf("EnvFile: %s\n", envFile)
	fmt.Printf("ReqFile: %s\n", reqFile)
	fmt.Printf("OnlinePOIHost: %s\n", onlinePOIHost)
	fmt.Printf("DebugPOIHost: %s\n", debugPOIHost)
	fmt.Printf("ShowColor: %v\n", showColor)
	fmt.Printf("LogLevel: %s\n", logLevelStr)
	fmt.Printf("OutputDir: %s\n", outputDir)
}
