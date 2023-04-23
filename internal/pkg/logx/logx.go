package logx

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

const logKey = "LogKey"

var Logger *zap.Logger

func init() {
	level := zap.DebugLevel

	conf := zap.NewProductionEncoderConfig()
	conf.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(conf),                            // json格式日志（ELK渲染收集）
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), // 打印到控制台和文件
		level, // 日志级别
	)

	// 开启文件及行号
	development := zap.Development()
	Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.DPanicLevel), development) // error级别日志，打印堆栈
}

// NewContext 给指定的context添加字段
func NewContext(ctx context.Context, fields ...zapcore.Field) context.Context {
	return context.WithValue(ctx, logKey, WithContext(ctx).With(fields...))
}

func WithContext(ctx context.Context) *zap.Logger {
	if ctx == nil {
		return Logger
	}
	if ctxLogger, ok := ctx.Value(logKey).(*zap.Logger); ok {
		return ctxLogger
	}
	return Logger
}

func String(key string, val string) zap.Field {
	return zap.Field{Key: key, Type: zapcore.StringType, String: val}
}

func Error(err error) zap.Field {
	return zap.NamedError("error", err)
}
