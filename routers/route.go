package routers

import (
	"GinCasbin/middleware"
	"GinCasbin/utils/ACS"
	"GinCasbin/utils/APIResponse"
	"GinCasbin/utils/Cache"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

var (
	R *gin.Engine
)

type RolePolicy struct {
	Subject string `json:"subject,omitempty"`
	Object  string `json:"object,omitempty"`
	Action  string `json:"action,omitempty"`
}

func init() {
	R = gin.Default()
	R.NoRoute(func(c *gin.Context) {
		c.JSON(400, gin.H{"code": 400, "message": "Bad Request"})
	})
	api()
}
func api() {
	auth := R.Group("/api")
	{
		// 模拟添加一条Policy策略
		auth.POST("acs", func(c *gin.Context) {
			APIResponse.C = c
			var rolePolicy RolePolicy
			err := c.ShouldBind(&rolePolicy)
			if err != nil {
				APIResponse.Error("bind fail")
				return
			}
			cacheName := rolePolicy.Subject + rolePolicy.Object + rolePolicy.Action
			fmt.Println(cacheName)
			result, err := ACS.Enforcer.AddPolicy(rolePolicy.Subject, rolePolicy.Object, rolePolicy.Action)
			if err != nil {
				fmt.Println(err.Error())
				APIResponse.Error("add fail " + err.Error())
				return
			}
			if result {
				APIResponse.Success("add success")
			} else {
				APIResponse.Error("add fail, maybe exist")
			}
		})
		// 模拟删除一条Policy策略
		auth.DELETE("acs/:id", func(c *gin.Context) {
			APIResponse.C = c
			var rolePolicy RolePolicy
			var err error
			err = c.ShouldBind(&rolePolicy)
			if err != nil {
				APIResponse.Error("bind fail")
				return
			}
			cacheName := rolePolicy.Subject + rolePolicy.Object + rolePolicy.Action
			result, err := ACS.Enforcer.RemovePolicy(rolePolicy.Subject, rolePolicy.Object, rolePolicy.Action)

			if err != nil {
				APIResponse.Error("delete Policy fail")
				return
			}
			if result {
				// 清除缓存
				err = Cache.GlobalCache.Delete(cacheName)
				if err != nil {
					log.Fatalf("delete cache fail")
				}
			}
			APIResponse.Success("delete success")
		})
		// 获取路由列表
		auth.POST("/routers", middleware.Privilege(), func(c *gin.Context) {
			type data struct {
				Method string `json:"method"`
				Path   string `json:"path"`
			}
			var datas []data
			routers := R.Routes()
			for _, v := range routers {
				var temp data
				temp.Method = v.Method
				temp.Path = v.Path
				datas = append(datas, temp)
			}
			APIResponse.C = c
			APIResponse.Success(datas)
			return
		})
	}
	// 定义路由组
	user := R.Group("/api/v1")
	// 使用访问控制中间件
	user.Use(middleware.Privilege())
	{
		user.POST("user", func(c *gin.Context) {
			c.JSON(200, gin.H{"code": 200, "message": "user add success"})
		})
		user.DELETE("user/:id", func(c *gin.Context) {
			id := c.Param("id")
			c.JSON(200, gin.H{"code": 200, "message": "user delete success " + id})
		})
		user.PUT("user/:id", func(c *gin.Context) {
			id := c.Param("id")
			c.JSON(200, gin.H{"code": 200, "message": "user update success " + id})
		})
		user.GET("user/:id", func(c *gin.Context) {
			id := c.Param("id")
			c.JSON(200, gin.H{"code": 200, "message": "user Get success " + id})
		})
	}
}
