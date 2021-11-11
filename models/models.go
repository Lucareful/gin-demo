package models

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"

	"github.com/luenci/go-gin-example/pkg/setting"
)

var db *gorm.DB

// init 数据库初始化.
func init() {
	var (
		err                          error
		dbName, user, password, host string
	)

	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println(err)
	}

	db = db.Debug()

	sqlDB, err := db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	db.AutoMigrate(&CasbinTable{}, &Article{}, &Tag{})

}

// GetDB 获取一个数据库链接
func GetDB() *gorm.DB {
	return db
}
