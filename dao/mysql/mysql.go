package mysql

import (
	"gin_web_demo/setting"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

// Init初始化MySQL连接
func Init(config *setting.MySQLConfig) (err error) {
	// "user:password@tcp(host:port)/dbname"
	conn := config.User + ":" + config.Password + "@tcp(" + config.Host + ":" + strconv.Itoa(config.Port) + ")/" + config.DB + "?charset=utf8&parseTime=true&loc=Asia%2FShanghai&allowNativePasswords=true"
	db, err = sqlx.Connect("mysql", conn)
	if err != nil {
		return
	}
	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetMaxIdleConns(config.MaxIdleConns)
	return
}

// Close 关闭MySQL连接
func CLose() {
	_ = db.Close()
}
