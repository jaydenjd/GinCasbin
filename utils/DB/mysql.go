package DB

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
import _ "github.com/go-sql-driver/mysql"

var (
	Mysql *gorm.DB
)

func init() {
	var err error
	dns := "root:root@(127.0.0.1:3306)/ginCasbin?charset=utf8&parseTime=True&loc=Local"
	Mysql, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		fmt.Println("connect DB error")
		panic(err)
	}
}
