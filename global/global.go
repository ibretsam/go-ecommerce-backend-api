package global

import (
	"go-ecommerce-backend-api/pkg/logger"
	"go-ecommerce-backend-api/pkg/settings"
)

var (
	Config settings.Config
	Logger *logger.ZapLogger
)
