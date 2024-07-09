package cmd

import (
	"bytes"
	"testing"

	"github.com/go-logr/zapr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestCustomLogger(t *testing.T) {
	// Create a buffer to capture log output
	var buf bytes.Buffer

	// Create a custom zap logger
	zapConfig := zap.NewDevelopmentConfig()
	zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	zapLogger := zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(zapConfig.EncoderConfig),
		zapcore.AddSync(&buf),
		zapcore.DebugLevel,
	))

	// Create a logr.Logger from the zap logger
	customLogger := zapr.NewLogger(zapLogger)

	// Set the global cliLogger
	cliLogger = customLogger

	// Now when we run a command, it should use our custom logger
	cmd := rootCmd
	cmd.SetArgs([]string{"version"})
	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Check if our custom logger was used
	logOutput := buf.String()
	if logOutput == "" {
		t.Error("Expected log output, but got none")
	}

	// You can add more specific checks on the log output here
	t.Logf("Log output: %s", logOutput)
}
