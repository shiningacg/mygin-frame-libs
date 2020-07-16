package mygin_frame_libs

import (
	"encoding/json"
	"errors"
	"github.com/shiningacg/mygin-frame-libs/log"
	"github.com/shiningacg/mygin-frame-libs/mysql"
	"github.com/shiningacg/mygin-frame-libs/redis"
	"io/ioutil"
	"os"
)

func Load(confDir string) {
	rc := loadRedisConfig(confDir)
	mc := loadMysqlConfig(confDir)
	lg := loadLogConfig(confDir)
	mysql.OpenMysql(mc)
	redis.OpenRedis(rc)
	log.OpenLog(lg)
}

func loadRedisConfig(path string) *redis.Config {
	config := &redis.Config{}
	redisConfigPath := path + "/redis.json"
	err := readJsonFile(redisConfigPath, config)
	if err != nil {
		panic("mysql加载失败：" + err.Error())
	}
	return config
}

func loadMysqlConfig(path string) *mysql.Config {
	config := &mysql.Config{}
	mysqlConfigPath := path + "/mysql.json"
	err := readJsonFile(mysqlConfigPath, config)
	if err != nil {
		panic("mysql加载失败：" + err.Error())
	}
	return config
}

func loadLogConfig(path string) *log.Config {
	config := &log.Config{}
	mysqlConfigPath := path + "/log.json"
	err := readJsonFile(mysqlConfigPath, config)
	if err != nil {
		panic("log加载失败：" + err.Error())
	}
	return config
}

func readJsonFile(path string, target interface{}) error {
	f, err := os.Open(path)
	if err != nil {
		return errors.New("没有找到文件:" + path)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, target)
	if err != nil {
		return err
	}
	return nil
}
