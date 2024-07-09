package logger

import (
	"io"

	"github.com/fatih/color"
	"github.com/go-logr/logr"
	"github.com/go-logr/zerologr"
	gcrLog "github.com/google/go-containerregistry/pkg/logs"
	"github.com/rs/zerolog"
	runtimeLog "sigs.k8s.io/controller-runtime/pkg/log"
)

func NewConsoleLogger(verbose bool, jsonFormat bool) logr.Logger {
	color.NoColor = !verbose
	zconfig := zerolog.ConsoleWriter{Out: color.Error, NoColor: !verbose}
	if jsonFormat {
		zconfig = zerolog.ConsoleWriter{Out: color.Error, NoColor: true}
	}

	zlog := zerolog.New(zconfig).With().Timestamp().Logger()

	gcrLog.Warn.SetOutput(io.Discard)

	zerologr.VerbosityFieldName = ""
	log := zerologr.New(&zlog)

	runtimeLog.SetLogger(log)

	return log
}
