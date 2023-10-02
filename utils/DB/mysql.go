package DB

import (
	"fmt"
	"github.com/jinzhu/gorm"
)
import _ "github.com/go-sql-driver/mysql"

var (
	Mysql *gorm.DB
)

func init() {
	var err error
	dsn := "root:root@(127.0.0.1:3306)/ginCasbin?charset=utf8&parseTime=True&loc=Local"
	Mysql, err = gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println("connect DB error")
		panic(err)
	}
}
