package initialize

import (
	"github.com/pqt2p1/go-ecommerce-backend-api/global"
	"github.com/pqt2p1/go-ecommerce-backend-api/pkg/logger"
)

func InitLogger() {
	global.Logger = *logger.NewLogger(global.Config.Logger)
}
