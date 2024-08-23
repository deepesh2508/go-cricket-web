package logging

import (
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/deepesh2508/go-cricket-web/env"
)

var Log *zap.Logger

func init() {
	level := getLevel(env.ENV.LOG_LEVEL)

	devConfig := zap.NewDevelopmentEncoderConfig()
	prodConfig := zap.NewProductionEncoderConfig()

	devConfig.EncodeTime = zapcore.RFC3339NanoTimeEncoder
	prodConfig.EncodeTime = zapcore.RFC3339NanoTimeEncoder

	filewriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "logs/" + strings.ToLower(strings.ReplaceAll(env.ENV.PROCESS_NAME, " ", "_")) + ".log",
		MaxSize:    50,
		MaxAge:     30,
		MaxBackups: 100,
		Compress:   true, // disabled by default
	})
	core := zapcore.NewCore(zapcore.NewJSONEncoder(prodConfig), filewriter, level)

	if env.ENV.DEPL_ENV == "DEV" {
		core = zapcore.NewTee(
			core,
			zapcore.NewCore(zapcore.NewConsoleEncoder(devConfig), zapcore.Lock(os.Stdout), level),
		)
	}

	if env.ENV.KAFKA_LOG == "Y" {
		kafkaSync := zapcore.AddSync(getKafkaWriter())
		core = zapcore.NewTee(
			core,
			zapcore.NewCore(zapcore.NewJSONEncoder(prodConfig), kafkaSync, level),
		)
	}

	Log = zap.New(core,
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		// zap.AddStacktrace(zap.ErrorLevel),
	)
}

func Info(ctx *gin.Context, msg string, args ...zap.Field) {
	Log.Info(msg, append([]zap.Field{zap.String("uuid", ctx.GetString("uuid"))}, args...)...)
}

func Error(ctx *gin.Context, msg string, args ...zap.Field) {
	Log.Error(msg, append([]zap.Field{zap.String("uuid", ctx.GetString("uuid"))}, args...)...)
}

func Warn(ctx *gin.Context, msg string, args ...zap.Field) {
	Log.Warn(msg, append([]zap.Field{zap.String("uuid", ctx.GetString("uuid"))}, args...)...)
}

func Debug(ctx *gin.Context, msg string, args ...zap.Field) {
	Log.Debug(msg, append([]zap.Field{zap.String("uuid", ctx.GetString("uuid"))}, args...)...)
}

func Fatal(ctx *gin.Context, msg string, args ...zap.Field) {
	Log.Fatal(msg, append([]zap.Field{zap.String("uuid", ctx.GetString("uuid"))}, args...)...)
}

func getLevel(lvl string) zapcore.Level {
	switch lvl {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	case "dpanic":
		return zap.DPanicLevel
	case "panic":
		return zap.PanicLevel
	case "fatal":
		return zap.FatalLevel
	default:
		return zap.ErrorLevel
	}
}

func getKafkaWriter() *kafkaLogger {
	kafkaLogger := &kafkaLogger{
		kafka.Writer{
			Addr:                   kafka.TCP(strings.Split(env.ENV.KAFKA_BROKERS, ",")...),
			Topic:                  env.ENV.KAFKA_BROKERS,
			Balancer:               &kafka.RoundRobin{},
			MaxAttempts:            2,
			BatchSize:              100,
			BatchTimeout:           time.Second * 1,
			WriteTimeout:           time.Second * 2,
			RequiredAcks:           kafka.RequireNone,
			AllowAutoTopicCreation: true,
		},
	}
	return kafkaLogger
}
