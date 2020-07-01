package redis

import (
	"context"
	"fmt"
	"testing"
)

func TestOpenRedis(t *testing.T) {
	OpenRedis(&Config{
		Host:     "127.0.0.1:6379",
		Secret:   "",
		Identity: DEFAULT,
	})
	_, err := Default().Set(context.TODO(), "test", "aaa", 0).Result()
	if err != nil {
		panic(err)
	}
	r, err := Default().Get(context.TODO(), "test").Result()
	if err != nil {
		panic(err)
	}
	if r != "aaa" {
		panic(fmt.Errorf("got %v, want %v", r, "aaa"))
	}
	Default().Close()
}

func TestIdentity(t *testing.T) {
	OpenRedis(&Config{
		Host:     "127.0.0.1:6379",
		Secret:   "",
		Identity: DEFAULT,
	})
	_, err := Default().Group("db1").Set(context.TODO(), "test", "aaa", 0).Result()
	if err != nil {
		panic(err)
	}
	r, _ := Default().Group("db2").Get(context.TODO(), "test").Result()
	if r != "" {
		panic("mix database")
	}
	Default().Close()
}
