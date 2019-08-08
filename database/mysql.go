package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"go-gin/config"
	"log"
	"time"
)

var DB *gorm.DB

func init()	{
	var err  error
	cfg := config.Cfg
	optionStr	:= cfg.Mysql.SqlUsername + ":"	+ cfg.Mysql.SqlPassword + "@tcp(" + cfg.Mysql.SqlHost + ":" + cfg.Mysql.SqlPort + ")/" + cfg.Mysql.SqlDatabase + "?parseTime=true"
	DB, err = gorm.Open("mysql", optionStr)
	if err != nil {
		fmt.Print(err)
		log.Fatal(err.Error())
		panic(err)
	}
	DB.DB().SetMaxIdleConns(10) //连接池空闲数大小
	DB.DB().SetMaxOpenConns(100)	//最大打开连接数
	DB.DB().SetConnMaxLifetime(time.Hour * 6) //连接最长存活时间
	DBMigrate()
}

func NewDb() * gorm.DB  {
	return DB
}

func DBMigrate()  {

	DB.SingularTable(true) //禁用表名

}



/*var SqlDB *sql.DB*/

/*func init()  {
	var err  error
	SqlDB, err = sql.Open("mysql", "root:tf2019@tcp(192.168.41.99:3306)/gin?parseTime=true")

	if err != nil {
		log.Fatal(err.Error())
	}

	err = SqlDB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}*/
