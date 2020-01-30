package utils

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var Logger *zap.Logger

func init() {
	hook := lumberjack.Logger{
		Filename:   "./logs/spikeProxy.log", // 日志文件路径
		MaxSize:    128,                     // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 30,                      // 日志文件最多保存多少个备份
		MaxAge:     7,                       // 文件最多保存多少天
		Compress:   true,                    // 是否压缩
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "Logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.DebugLevel)

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),                                           // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
		atomicLevel, // 日志级别
	)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段
	filed := zap.Fields(zap.String("serviceName", "GoLearn"))
	// 构造日志
	Logger = zap.New(core, caller, development, filed)
}

func Debugf(format string, args ...interface{}) {
	Logger.Debug(fmt.Sprintf(format, args...))
}

func Infof(format string, args ...interface{}) {
	Logger.Info(fmt.Sprintf(format, args...))
}

func Warnf(format string, args ...interface{}) {
	Logger.Warn(fmt.Sprintf(format, args...))
}

func Errorf(format string, args ...interface{}) {
	Logger.Error(fmt.Sprintf(format, args...))
}
