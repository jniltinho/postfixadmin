package log

import (
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const LOGGER_KEY = "zapLogger"

type Logger struct {
	*zap.Logger
}

func NewLog(conf *viper.Viper) *Logger {
	return initZap(conf)
}

func initZap(conf *viper.Viper) *Logger {
	// 日志地址 "out.log" 自定义
	lp := conf.GetString("log.log_file_name")
	// 日志级别 DEBUG,ERROR, INFO
	lv := conf.GetString("log.log_level")
	var level zapcore.Level
	//debug<info<warn<error<fatal<panic
	switch lv {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}
	hook := lumberjack.Logger{
		Filename:   lp,                             // 日志文件路径
		MaxSize:    conf.GetInt("log.max_size"),    // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: conf.GetInt("log.max_backups"), // 日志文件最多保存多少个备份
		MaxAge:     conf.GetInt("log.max_age"),     // 文件最多保存多少天
		Compress:   conf.GetBool("log.compress"),   // 是否压缩
	}

	var encoder zapcore.Encoder
	if conf.GetString("log.encoding") == "console" {
		encoder = zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "Logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseColorLevelEncoder,
			EncodeTime:     timeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.FullCallerEncoder,
		})
	} else {
		encoder = zapcore.NewJSONEncoder(zapcore.EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.EpochTimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		})
	}
	core := zapcore.NewCore(
		encoder, // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
		level, // 日志级别
	)
	if conf.GetString("env") != "prod" {
		return &Logger{zap.New(core, zap.Development(), zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))}
	}
	return &Logger{zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))}

}

// 自定义时间编码器
func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	//enc.AppendString(t.Format("2006-01-02 15:04:05"))
	enc.AppendString(t.Format("2006-01-02 15:04:05.000000000"))
}

// NewContext 给指定的context添加字段
func (l *Logger) NewContext(ctx echo.Context, fields ...zapcore.Field) {
	ctx.Set(LOGGER_KEY, l.WithContext(ctx).With(fields...))
}

// WithContext 从指定的context返回一个zap实例
func (l *Logger) WithContext(ctx echo.Context) *Logger {
	if ctx == nil {
		return l
	}
	zl := ctx.Get(LOGGER_KEY)
	ctxLogger, ok := zl.(*zap.Logger)
	if ok {
		return &Logger{ctxLogger}
	}
	return l
}

func (l *Logger) EchoLog() middleware.RequestLoggerConfig {

	return middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			l.Info("request", zap.String("remote_ip", c.RealIP()),
				zap.String("host", c.Request().Host),
				zap.String("method", c.Request().Method),
				zap.String("uri", v.URI),
				zap.String("user_agent", c.Request().UserAgent()),
				zap.Int("status", v.Status),
				zap.Int64("bytes_in", c.Request().ContentLength),
				zap.Int64("bytes_out", c.Response().Size),
			)
			return nil
		},
	}
}
