package mygin_frame_libs

import (
	"context"
	"fmt"
	"github.com/shiningacg/mygin-frame-libs/log"
	"github.com/shiningacg/mygin-frame-libs/mysql"
	"github.com/shiningacg/mygin-frame-libs/redis"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	Load("../conf/dev")
	r, err := redis.Default().Get(context.TODO(), "test").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(r)
	redis.Default().Close()
	mysql.Default().Close()
}
func TestLoadLogJsonFile(t *testing.T) {
	config := loadLogConfig("./log")
	log.OpenLog(config)
	log.Default().Log("hi")
}
