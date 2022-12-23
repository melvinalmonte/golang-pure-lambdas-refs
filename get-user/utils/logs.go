package utils

import (
	"go.elastic.co/ecszap"
	"go.uber.org/zap"
)

func InitLogger() {
	config := zap.NewProductionConfig()
	config.EncoderConfig = ecszap.ECSCompatibleEncoderConfig(config.EncoderConfig)
	logger, _ := config.Build(ecszap.WrapCoreOption(), zap.AddCaller())

	_ = zap.ReplaceGlobals(logger)
}
