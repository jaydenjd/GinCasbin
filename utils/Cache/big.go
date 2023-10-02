package Cache

import (
	"github.com/allegro/bigcache/v2"
	"time"
)

var GlobalCache *bigcache.BigCache

func init() {
	// 初始化BigCache实例
	var err error
	GlobalCache, err = bigcache.NewBigCache(bigcache.DefaultConfig(1 * time.Minute))
	if err != nil {
		panic(err)
	}
}
