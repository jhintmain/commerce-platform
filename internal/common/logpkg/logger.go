package logpkg

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New(service string) *zap.Logger {
	encoderCfg := zapcore.EncoderConfig{
		TimeKey:      "ts",
		LevelKey:     "level",
		CallerKey:    "caller",
		MessageKey:   "msg",
		EncodeLevel:  zapcore.LowercaseLevelEncoder,
		EncodeTime:   zapcore.ISO8601TimeEncoder, // 한국시간
		EncodeCaller: zapcore.ShortCallerEncoder,
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.AddSync(os.Stdout),
		zap.InfoLevel,
	)

	return zap.New(core, zap.AddCaller()).
		With(zap.String("service", service))
}
