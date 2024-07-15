package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

/*
* init initializes the logger with the default configuration.
 */
func init() {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.StacktraceKey = ""

	var err error
	config := zap.NewProductionConfig()
	config.EncoderConfig = encoderConfig
	log, err = config.Build(zap.AddCallerSkip(1))

	if err != nil {
		panic(err)
	}
}

/*
* Info logs a message at InfoLevel. The message includes any fields passed at the log site, as well as any fields accumulated on the logger.
 */
func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

/*
* Debug logs a message at DebugLevel. The message includes any fields passed at the log site, as well as any fields accumulated on the logger.
 */
func Debug(message string, fields ...zap.Field) {
	log.Debug(message, fields...)
}

/*
* Error logs a message at ErrorLevel. The message includes any fields passed at the log site, as well as any fields accumulated on the logger.
 */
func Error(message string, fields ...zap.Field) {
	log.Error(message, fields...)
}
