package logger

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

type LoggerInterface interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	DPanic(args ...interface{})
	Panic(args ...interface{})
	Fatal(args ...interface{})
	Debugf(template string, args ...interface{})
	Infof(template string, args ...interface{})
	Warnf(template string, args ...interface{})
	Errorf(template string, args ...interface{})
	DPanicf(template string, args ...interface{})
	Panicf(template string, args ...interface{})
	Fatalf(template string, args ...interface{})
}

type log struct {
	zapLog *zap.Logger
}

type LoggerInfo struct {
	Level      string
	EncodeTime string
	LoggerFile struct {
		LogPath    string
		ErrPath    string
		MaxSize    int
		MaxBackups int
		MaxAge     int
		Compress   bool
	}
}

func NewLogger(conf LoggerInfo) LoggerInterface {
	var logger *zap.Logger
	zapConfig := zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    "level",
		EncodeLevel: zapcore.CapitalLevelEncoder, //将级别转换成大写
		TimeKey:     "ts",
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format(conf.EncodeTime))
		},
		CallerKey:    "file",
		EncodeCaller: zapcore.FullCallerEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	}
	encoder := zapcore.NewConsoleEncoder(zapConfig)
	// 设置级别
	logLevel := zap.DebugLevel
	switch conf.Level {
	case "debug":
		logLevel = zap.DebugLevel
	case "info":
		logLevel = zap.InfoLevel
	case "warn":
		logLevel = zap.WarnLevel
	case "error":
		logLevel = zap.ErrorLevel
	case "panic":
		logLevel = zap.PanicLevel
	case "fatal":
		logLevel = zap.FatalLevel
	default:
		logLevel = zap.InfoLevel
	}
	// 实现两个判断日志等级的interface  可以自定义级别展示
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.WarnLevel && lvl >= zap.InfoLevel
	})

	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel
	})
	infoLumberJackLogger := &lumberjack.Logger{
		Filename:   conf.LoggerFile.LogPath + time.Now().Format("2006-01-02") + ".log", //日志文件的位置
		MaxSize:    conf.LoggerFile.MaxSize,                                            //在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: conf.LoggerFile.MaxBackups,                                         //保留旧文件的最大个数
		MaxAge:     conf.LoggerFile.MaxAge,                                             //保留旧文件的最大天数
		Compress:   conf.LoggerFile.Compress,                                           //是否压缩/归档旧文件
	}
	errLumberJackLogger := &lumberjack.Logger{
		Filename:   conf.LoggerFile.ErrPath + time.Now().Format("2006-01-02") + ".log", //日志文件的位置
		MaxSize:    conf.LoggerFile.MaxSize,                                            //在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: conf.LoggerFile.MaxBackups,                                         //保留旧文件的最大个数
		MaxAge:     conf.LoggerFile.MaxAge,                                             //保留旧文件的最大天数
		Compress:   conf.LoggerFile.Compress,                                           //是否压缩/归档旧文件
	}
	var core zapcore.Core
	if conf.Level == "debug" {
		core = zapcore.NewTee(
			//日志都会在console中展示
			zapcore.NewCore(zapcore.NewConsoleEncoder(zapConfig), zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), logLevel),
		)
	} else {
		core = zapcore.NewTee(
			// 记录日志到文件
			zapcore.NewCore(encoder, zapcore.AddSync(infoLumberJackLogger), infoLevel),
			zapcore.NewCore(encoder, zapcore.AddSync(errLumberJackLogger), warnLevel),
			//日志都会在console中展示
			zapcore.NewCore(zapcore.NewConsoleEncoder(zapConfig), zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), logLevel),
		)
	}
	// 最后创建具体的Logger
	logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.WarnLevel)) // 需要传入 zap.AddCaller() 才会显示打日志点的文件名和行数, 有点小坑
	zap.ReplaceGlobals(logger)
	return &log{
		zapLog: logger,
	}
}

func (l *log) Debug(args ...interface{}) {
	l.zapLog.Sugar().Debug(args)
}

func (l *log) Info(args ...interface{}) {
	l.zapLog.Sugar().Info(args)
}

func (l *log) Warn(args ...interface{}) {
	l.zapLog.Sugar().Warn(args)
}

func (l *log) Error(args ...interface{}) {
	l.zapLog.Sugar().Error(args)
}

func (l *log) DPanic(args ...interface{}) {
	l.zapLog.Sugar().DPanic(args)
}

func (l *log) Panic(args ...interface{}) {
	l.zapLog.Sugar().Panic(args)
}

func (l *log) Fatal(args ...interface{}) {
	l.zapLog.Sugar().Fatal(args)
}

func (l *log) Debugf(template string, args ...interface{}) {
	l.zapLog.Sugar().Debugf(template, args)
}

func (l *log) Infof(template string, args ...interface{}) {
	l.zapLog.Sugar().Infof(template, args)
}

func (l *log) Warnf(template string, args ...interface{}) {
	l.zapLog.Sugar().Warnf(template, args)
}

func (l *log) Errorf(template string, args ...interface{}) {
	l.zapLog.Sugar().Errorf(template, args)
}

func (l *log) DPanicf(template string, args ...interface{}) {
	l.zapLog.Sugar().DPanicf(template, args)
}

func (l *log) Panicf(template string, args ...interface{}) {
	l.zapLog.Sugar().Panicf(template, args)
}

func (l *log) Fatalf(template string, args ...interface{}) {
	l.zapLog.Sugar().Fatalf(template, args)
}
