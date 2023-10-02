package middleware

import (
	"GinCasbin/utils/ACS"
	"GinCasbin/utils/APIResponse"
	"GinCasbin/utils/Cache"
	"github.com/gin-gonic/gin"
	"log"
)

func Privilege() gin.HandlerFunc {
	return func(c *gin.Context) {
		APIResponse.C = c
		var userName = c.GetHeader("userName")
		if userName == "" {
			APIResponse.Error("header miss userName")
			c.Abort()
			return
		}
		path := c.Request.URL.Path
		method := c.Request.Method
		cacheName := userName + path + method
		// 从缓存中读取&判断
		entry, err := Cache.GlobalCache.Get(cacheName)
		if err == nil && entry != nil {
			if string(entry) == "true" {
				c.Next()
			} else {
				APIResponse.Error("access denied")
				c.Abort()
				return
			}
		} else {
			// 从数据库中读取&判断
			// 加载策略规则
			err := ACS.Enforcer.LoadPolicy()
			if err != nil {
				log.Println("loadPolicy error")
				panic(err)
			}
			// 验证策略规则
			result, err := ACS.Enforcer.Enforce(userName, path, method)
			if err != nil {
				APIResponse.Error("No permission found")
				c.Abort()
				return
			}
			if !result {
				// 添加到缓存中
				Cache.GlobalCache.Set(cacheName, []byte("false"))
				APIResponse.Error("access denied")
				c.Abort()
				return
			} else {
				Cache.GlobalCache.Set(cacheName, []byte("true"))
			}
			c.Next()
		}
	}
}
