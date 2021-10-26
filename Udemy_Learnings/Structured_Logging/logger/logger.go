package logger

import (
	"go.uber.org/zap" //https://pkg.go.dev/go.uber.org/zap#section-readme
	"go.uber.org/zap/zapcore"
)

//var Log *zap.logger
var log *zap.Logger

func init() {
	var err error
	//Log,err=zap.NewProduction()
	log, err = zap.NewProduction() // "caller":"logger/logger.go
	//log, err = zap.NewProduction(zap.AddCallerSkip(1))  --->"caller":"Structured_Logging/main.go

	// To change the timestamp to ISO
	config := zap.NewProductionConfig()
	encoderconfig := zap.NewProductionEncoderConfig()
	encoderconfig.TimeKey = "TimeStamp"
	encoderconfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderconfig.StacktraceKey = "" //sets the stacktrace to empty inside the log
	config.EncoderConfig = encoderconfig

	log, err = config.Build()

	if err != nil {
		panic(err)
	}

}

//helper function to return message and fields
func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}
func Debug(message string, fields ...zap.Field) {
	log.Debug(message, fields...)
}
func Error(message string, fields ...zap.Field) {
	log.Error(message, fields...)
}
