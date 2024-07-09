package logger

import (
	"io"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/go-logr/logr"
	"github.com/go-logr/zerologr"
	gcrLog "github.com/google/go-containerregistry/pkg/logs"
	"github.com/rs/zerolog"
	runtimeLog "sigs.k8s.io/controller-runtime/pkg/log"
)

func NewConsoleLogger(verbose bool, jsonFormat bool) logr.Logger {
	color.NoColor = !verbose
	zconfig := zerolog.ConsoleWriter{
		Out:        os.Stderr,
		NoColor:    !verbose,
		TimeFormat: time.Kitchen,
	}

	if jsonFormat {
		zconfig.NoColor = true
	}

	// Configure timestamp output
	if !verbose {
		zconfig.PartsExclude = []string{zerolog.TimestampFieldName}
	}

	zlog := zerolog.New(zconfig).Level(zerolog.InfoLevel)

	if verbose {
		zlog = zlog.Level(zerolog.DebugLevel)
	}

	zlog = zlog.With().Timestamp().Logger()

	gcrLog.Warn.SetOutput(io.Discard)

	zerologr.VerbosityFieldName = "v"
	log := zerologr.New(&zlog)

	runtimeLog.SetLogger(log)

	return log
}
