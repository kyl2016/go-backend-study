package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	logMsg := map[string]interface{}{
		"a": "b",
		"c": "d",
		"e": map[string]interface{}{
			"f": "g",
		},
	}

	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, _ := cfg.Build()
	defer logger.Sync()
	sugar := logger.Sugar()
	sugar.Errorw("This is a short message description", "key1", "value1", "key2", logMsg)
}
