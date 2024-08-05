package global

import (
	"go-ecommerce-backend-api/pkg/logger"
	"go-ecommerce-backend-api/pkg/settings"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config settings.Config
	Logger *logger.ZapLogger
	Rdb    *redis.Client
	Mdb    *gorm.DB
)
