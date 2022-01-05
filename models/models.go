package models

import (
	"fmt"
	"time"

	"github.com/luenci/go-gin-example/config"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

var db *gorm.DB

// InitDB 数据库初始化.
func InitDB() {

	conf := config.GetConf()
	fmt.Printf("%+v\n", conf)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.Mysql.User,
		conf.Mysql.Password,
		conf.Mysql.Host,
		conf.Mysql.Port,
		conf.Mysql.DBName)
	fmt.Println(dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db = db.Debug()

	sqlDB, err := db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	db.AutoMigrate(&CasbinTable{}, &Article{}, &Tag{}, &User{})

}

// GetDB 获取一个数据库链接
func GetDB() *gorm.DB {
	return db
}
