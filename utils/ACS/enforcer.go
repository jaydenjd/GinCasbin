package ACS

import (
	"GinCasbin/utils/DB"
	"github.com/casbin/casbin"
	"github.com/casbin/gorm-adapter"
)

var Enforcer *casbin.Enforcer

func init() {
	// mysql 适配器
	adapter := gormadapter.NewAdapterByDB(DB.Mysql)
	// 通过mysql适配器新建一个enforcer
	Enforcer = casbin.NewEnforcer("config/keymatch2_model.conf", adapter)
	// 日志记录
	Enforcer.EnableLog(true)
}
