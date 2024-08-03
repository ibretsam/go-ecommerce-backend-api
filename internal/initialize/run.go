package initialize

import (
	"fmt"
	"go-ecommerce-backend-api/global"
)

func Run() {
	LoadConfig()
	fmt.Println("Loading configuration mysql: ", global.Config.Mysql.User)
	InitLogger()
	InitMySQL()
	InitRedis()
	r := InitRouter()
	r.Run("localhost:8080")
}
