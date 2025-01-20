package global

import (
	"github.com/pqt2p1/go-ecommerce-backend-api/pkg/logger"
	"github.com/pqt2p1/go-ecommerce-backend-api/pkg/setting"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config = setting.Config{}
	Logger = *logger.NewLogger(Config.Logger)
	Mdb    *gorm.DB
	Rdb    *redis.Client
)
