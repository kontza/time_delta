package cmd

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"time"
	"go.uber.org/zap/zapcore"
	"os"
	"github.com/spf13/viper"
)

var logger *zap.SugaredLogger

func loggerInit(cmd *cobra.Command, args []string) {
	dec := zap.NewDevelopmentEncoderConfig()
	dec.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
	}
	dec.EncodeLevel = zapcore.CapitalColorLevelEncoder
	consoleDebugging := zapcore.Lock(os.Stdout)
	consoleEncoder := zapcore.NewConsoleEncoder(dec)
	loggingLevel := zap.InfoLevel
	if viper.GetBool("beVerbose") {
		loggingLevel = zap.DebugLevel
	}
	core := zapcore.NewCore(consoleEncoder, consoleDebugging, zap.NewAtomicLevelAt(loggingLevel))
	coreLogger := zap.New(core)
	defer coreLogger.Sync()
	logger = coreLogger.Sugar()
}