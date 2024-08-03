package initialize

import (
	"fmt"
	"go-ecommerce-backend-api/global"

	"go.uber.org/zap"
)

func Run() {
	LoadConfig()
	fmt.Println("Loading configuration mysql: ", global.Config.Mysql.User)
	InitLogger()
	global.Logger.Info("Logger init success", zap.String("info", "success"))
	InitMySQL()
	InitRedis()
	r := InitRouter()
	r.Run("localhost:8080")
}
