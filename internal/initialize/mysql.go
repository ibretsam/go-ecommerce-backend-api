package initialize

import (
	"fmt"
	"go-ecommerce-backend-api/global"
	"go-ecommerce-backend-api/internal/po"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func checkErrorPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(fmt.Errorf("%s: %s ", errString, err))
	}
}

func InitMySQL() {
	// Initialize MySQL
	m := global.Config.Mysql
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.UserName, m.Password, m.Host, m.Port, m.Dbname)
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		SkipDefaultTransaction: false,
	})
	checkErrorPanic(err, "MySQL initialization failed")
	global.Logger.Info("MySQL init success", zap.String("info", "success"))
	global.Mdb = db

	SetPool()
	MigrateTables()
}

func SetPool() {
	// Set pool
	sqlDb, err := global.Mdb.DB()
	if err != nil {
		global.Logger.Error("Set pool failed", zap.Error(err))
		fmt.Println("Set pool failed", err)
	}
	m := global.Config.Mysql
	sqlDb.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))
	sqlDb.SetMaxOpenConns(m.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime))
}

func MigrateTables() {
	// Migrate tables
	err := global.Mdb.AutoMigrate(
		&po.User{},
		&po.Role{},
	)
	checkErrorPanic(err, "Migrate tables failed")

}
