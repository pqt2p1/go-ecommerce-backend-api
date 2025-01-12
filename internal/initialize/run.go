package initialize

import (
	"fmt"

	"github.com/pqt2p1/go-ecommerce-backend-api/global"
)

func Run() {
	LoadConfig()
	fmt.Println("Loading mysql config", global.Config.Mysql.Username)
	InitLogger()
	global.Logger.Info("Logger init success")
	InitMysql()
	InitRedis()
	r := InitRouter()
	r.Run(":8002")
}
