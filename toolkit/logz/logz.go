package logz

import (
	"os"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Zap *zap.Logger

// SettingGlobalLogger 设置全局的日志输出
// output 日志输出的文件路径
// level  参考 zapcore.Level
func SettingGlobalLogger(output string, level zapcore.Level) {
	// 配置参数
	// 设置时间输出字段的key值，和格式化时间
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "time"
	encoderCfg.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05"))
	}

	// 以JSON格式输出日志
	encoder := zapcore.NewJSONEncoder(encoderCfg)

	// 终端输出log核心
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	// 设置多端输出log核心
	core := zapcore.NewTee(
		// 输出到日志文件
		zapcore.NewCore(encoder, logOutput(output), level),
		// 输出到终端
		zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), level),
	)

	// 实例化
	Zap = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	// zap.ReplaceGlobals(log)
}

// 日志输出与分割
func logOutput(path string) zapcore.WriteSyncer {
	// 默认配置
	lumberJackLogger := &lumberjack.Logger{
		Filename:   path,  // 日志文件名
		MaxSize:    10,    // 日志文件最大容量  MB
		MaxBackups: 10,    // 备份数量
		MaxAge:     30,    // 最大备份天数
		Compress:   false, // 是否压缩
	}
	return zapcore.AddSync(lumberJackLogger)
}
