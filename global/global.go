package global

import (
	"go-ecommerce-backend-api/pkg/logger"
	"go-ecommerce-backend-api/pkg/settings"

	"gorm.io/gorm"
)

var (
	Config settings.Config
	Logger *logger.ZapLogger
	Mdb    *gorm.DB
)
