package log

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"hinccvi/go-template/config"
	"os"
	"strconv"
	"time"
)

var log *zap.Logger
var sqlLog *zap.Logger

type GormLogger struct {
	ZapLogger                 *zap.Logger
	LogLevel                  gormlogger.LogLevel
	SlowThreshold             time.Duration
	SkipCallerLookup          bool
	IgnoreRecordNotFoundError bool
}

func Init(env string) {
	writeSyncer := getLogWriter(
		"./loggings/"+config.Conf.LogConfig.FileName,
		config.Conf.LogConfig.MaxSize,
		config.Conf.LogConfig.MaxBackup,
		config.Conf.LogConfig.MaxAge,
	)

	var core zapcore.Core
	if env == "dev" {
		core = zapcore.NewTee(zapcore.NewCore(getEncoder(env), zapcore.Lock(os.Stdout), zap.DebugLevel))
		sqlLog = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(3))
	} else {
		core = zapcore.NewTee(zapcore.NewCore(getEncoder(env), writeSyncer, zap.DebugLevel))
	}
	log = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	log.Info("Logger successfully init")
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func getEncoder(env string) zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Local().Format("2006-01-02T15:04:05Z0700"))
	}
	if env == "dev" {
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		return zapcore.NewConsoleEncoder(encoderConfig)
	}
	return zapcore.NewJSONEncoder(encoderConfig)
}

func Info(msg string, args ...interface{}) {
	log.Sugar().Infow(msg, args...)
}

func Debug(msg string, args ...interface{}) {
	log.Sugar().Debugw(msg, args...)
}

func Warn(msg string, args ...interface{}) {
	log.Sugar().Warnw(msg, args...)
}

func Error(msg string, args ...interface{}) {
	log.Sugar().Errorw(msg, args...)
}

func Panic(msg string, args ...interface{}) {
	log.Sugar().Panicw(msg, args...)
}

//	Configure new gorm logger setting and replacing it with zap
func New() GormLogger {
	return GormLogger{
		ZapLogger:                 sqlLog,
		LogLevel:                  gormlogger.Info,
		SlowThreshold:             100 * time.Millisecond,
		SkipCallerLookup:          false,
		IgnoreRecordNotFoundError: true,
	}
}

func (l GormLogger) SetAsDefault() {
	gormlogger.Default = l
}

func (l GormLogger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	return GormLogger{
		ZapLogger:                 l.ZapLogger,
		SlowThreshold:             l.SlowThreshold,
		LogLevel:                  level,
		SkipCallerLookup:          l.SkipCallerLookup,
		IgnoreRecordNotFoundError: l.IgnoreRecordNotFoundError,
	}
}

func (l GormLogger) Info(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel < gormlogger.Info {
		return
	}
	l.ZapLogger.Sugar().Debugf(str, args...)
}

func (l GormLogger) Warn(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel < gormlogger.Warn {
		return
	}
	l.ZapLogger.Sugar().Warnf(str, args...)
}

func (l GormLogger) Error(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel < gormlogger.Error {
		return
	}
	l.ZapLogger.Sugar().Errorf(str, args...)
}

func (l GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= 0 {
		return
	}
	elapsed := time.Since(begin)
	switch {
	case err != nil && l.LogLevel >= gormlogger.Error && (!l.IgnoreRecordNotFoundError || !errors.Is(err, gorm.ErrRecordNotFound)):
		l.ZapLogger.Sugar().Errorf("SQL\t\"elapsed\": %*s\"error\": %*s", getPadding(elapsed), elapsed, getPadding(err), err)
	case l.SlowThreshold != 0 && elapsed > l.SlowThreshold && l.LogLevel >= gormlogger.Warn:
		sql, rows := fc()
		l.ZapLogger.Sugar().Warnf("SQL\t\"elapsed\": %*s\"rows\": %*d\"query\": %s", getPadding(elapsed), elapsed, getPadding(rows), rows, sql)
	case l.LogLevel >= gormlogger.Info:
		sql, rows := fc()
		l.ZapLogger.Sugar().Debugf("SQL\t\"elapsed\": %*s\"rows\": %*d\"query\": %s", getPadding(elapsed), elapsed, getPadding(rows), rows, sql)
	}
}

func getPadding(data interface{}) int {
	switch v := data.(type) {
	case int64:
		s := strconv.Itoa(int(v))
		return (len(s) + 2 + 4) * -1
	case time.Duration:
		s := v.String()
		return (len(s) + 2 + 4) * -1
	case error:
		s := v.Error()
		return (len(s) + 2 + 4) * -1
	}

	return 0
}
