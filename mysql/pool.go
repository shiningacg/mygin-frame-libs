package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

const DEFAULT = "default"

var Pool = make(map[string]*gorm.DB)

func getDefault() *gorm.DB {
	return get(DEFAULT)
}

func get(key string) *gorm.DB {
	db, has := Pool[key]
	if !has {
		return nil
	}
	return db
}

func set(key string, db *gorm.DB) {
	if db := get(key); db != nil {
		err := db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}
	Pool[key] = db
}
