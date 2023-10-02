package Cache

import (
	"github.com/allegro/bigcache/v2"
	"time"
)

var GlobalCache *bigcache.BigCache

func init() {
	// 初始化BigCache实例
	GlobalCache, _ = bigcache.NewBigCache(bigcache.DefaultConfig(30 * time.Minute))
}
