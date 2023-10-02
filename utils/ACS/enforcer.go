package ACS

import (
	"GinCasbin/utils/DB"
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

var Enforcer *casbin.Enforcer

func init() {
	// 通过mysql适配器新建一个enforcer
	adapter, _ := gormadapter.NewAdapterByDB(DB.Mysql)
	// 通过model.conf文件配置策略
	var err error
	Enforcer, err = casbin.NewEnforcer("config/keymatch2_model.conf", adapter)
	if err != nil {
		fmt.Println("NewCachedEnforcer err:", err)
		panic(err)
	}
	Enforcer.EnableAutoSave(true)
	err = Enforcer.LoadPolicy()
	if err != nil {
		return
	}

	// 日志记录
	Enforcer.EnableLog(true)
}
