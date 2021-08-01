package services

import (
	"go.uber.org/zap"
)

func NewProductionLogger() (*zap.Logger, error) {
	c := zap.NewProductionConfig()
	c.DisableCaller = true
	c.DisableStacktrace = true
	return c.Build()
}
