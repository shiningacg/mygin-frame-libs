package redis

import (
	"fmt"
)

const DEFAULT = "DEFAULT"

var Pool = make(map[string]*Client)

func getDefault() *Client {
	return get(DEFAULT)
}

func get(key string) *Client {
	db, has := Pool[key]
	if !has {
		return nil
	}
	return db
}

func set(key string, db *Client) {
	if db := get(key); db != nil {
		err := db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}
	Pool[key] = db
}
