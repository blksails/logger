package logger

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log = zap.L()

type Logger = zap.Logger
type SugaredLogger = zap.SugaredLogger

// Mode represents the logging mode
type Mode string

const (
	// ModeDevelopment represents development mode with human-readable output
	ModeDevelopment Mode = "development"
	// ModeProduction represents production mode with JSON output
	ModeProduction Mode = "production"
)

// SetMode sets the logging mode (development or production)
func SetMode(mode Mode) error {
	var newLog *zap.Logger
	var err error

	switch mode {
	case ModeDevelopment:
		newLog, err = zap.NewDevelopment(zap.AddCallerSkip(0))
	case ModeProduction:
		config := zap.NewProductionConfig()
		config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		newLog, err = config.Build(zap.AddCallerSkip(0))
	default:
		return fmt.Errorf("invalid mode: %s", mode)
	}

	if err != nil {
		return err
	}

	log = newLog
	return nil
}

func NewDevelopment() (*zap.Logger, error) {
	return zap.NewDevelopment()
}

func GetLogger() *zap.SugaredLogger {
	return log.Sugar()
}

func NewNop() *zap.Logger {
	return zap.NewNop()
}

func init() {
	// Default to production mode
	SetMode(ModeDevelopment)
}

func Any(key string, value any) zap.Field {
	return zap.Any(key, value)
}

func String(key string, value string) zap.Field {
	return zap.String(key, value)
}

func Int(key string, value int) zap.Field {
	return zap.Int(key, value)
}

func Int64(key string, value int64) zap.Field {
	return zap.Int64(key, value)
}

func Int32(key string, value int32) zap.Field {
	return zap.Int32(key, value)
}

func Int16(key string, value int16) zap.Field {
	return zap.Int16(key, value)
}

func Int8(key string, value int8) zap.Field {
	return zap.Int8(key, value)
}

func Float64(key string, value float64) zap.Field {
	return zap.Float64(key, value)
}

func Float32(key string, value float32) zap.Field {
	return zap.Float32(key, value)
}

func Bool(key string, value bool) zap.Field {
	return zap.Bool(key, value)
}

func Time(key string, value time.Time) zap.Field {
	return zap.Time(key, value)
}

func Duration(key string, value time.Duration) zap.Field {
	return zap.Duration(key, value)
}

func Error(err error) zap.Field {
	return zap.Error(err)
}

func Errors(key string, errs []error) zap.Field {
	return zap.Errors(key, errs)
}

func Object(key string, value zapcore.ObjectMarshaler) zap.Field {
	return zap.Object(key, value)
}

func Stringer(key string, value fmt.Stringer) zap.Field {
	return zap.Stringer(key, value)
}
