package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

func OpenRedis(config *Config) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*2)
	rdb := redis.NewClient(config.Adapter())
	_, err := rdb.Ping(ctx).Result()
	cancel()
	if err != nil {
		panic(err)
	}
	set(config.Identity, &Client{Client: rdb, groups: make(map[string]*Group)})
}

// 供外部使用
func Default() *Client {
	if db := getDefault(); db != nil {
		return db
	}
	panic("无效的数据库")
}

func DefaultGroup() *Group {
	return Default().Default()
}

// 外部使用
func Get(key string) *Client {
	if db := get(key); db != nil {
		return db
	}
	panic("无效的数据库")
}

type Client struct {
	*redis.Client
	groups map[string]*Group
}

func (c *Client) Group(identity string) *Group {
	if group, has := c.groups[identity]; has {
		return group
	}
	group := &Group{
		c:        c,
		Identity: identity,
	}
	c.groups[identity] = group
	return group
}

func (c *Client) Default() *Group {
	return c.Group(DEFAULT)
}

type Group struct {
	c        *Client
	Identity string
	closed   bool
}

func (g *Group) key(k string) string {
	return g.Identity + "_" + k
}

func (g *Group) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	if g.closed {
		panic("使用已经关闭的redis group")
	}
	return g.c.Set(ctx, g.key(key), value, expiration)
}

func (g *Group) Get(ctx context.Context, key string) *redis.StringCmd {
	if g.closed {
		panic("使用已经关闭的redis group")
	}
	return g.c.Get(ctx, g.key(key))
}

func (g *Group) Ping(ctx context.Context) *redis.StatusCmd {
	return g.c.Ping(ctx)
}

func (g *Group) Close() error {
	delete(g.c.groups, g.Identity)
	g.closed = true
	if len(g.c.groups) == 0 {
		return g.c.Close()
	}
	return nil
}
