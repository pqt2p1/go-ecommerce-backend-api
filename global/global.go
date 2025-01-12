package global

import (
	"github.com/pqt2p1/go-ecommerce-backend-api/pkg/logger"
	"github.com/pqt2p1/go-ecommerce-backend-api/pkg/setting"
)

var (
	Config = setting.Config{}
	Logger = *logger.NewLogger(Config.Logger)
	Mdb    = nil
)
