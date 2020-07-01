package mysql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func OpenMysql(config *Config, options ...ConfigOption) {
	if options == nil {
		options = []ConfigOption{OptionDefault}
	}
	useOption(config, options)
	db, err := gorm.Open("mysql", config.DSN())
	if err != nil {
		panic(err)
	}
	set(config.Identity, db)
}

// 供外部使用
func Default() *gorm.DB {
	if db := getDefault(); db != nil {
		return db
	}
	panic("无效的数据库")
}

// 外部使用
func Get(key string) *gorm.DB {
	if db := get(key); db != nil {
		return db
	}
	panic("无效的数据库")
}
