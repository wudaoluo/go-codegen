package mysql

import (
	"bytes"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"github.com/wudaoluo/golog"
)

var db *sql.DB

func InitDB() {
	var buf bytes.Buffer
	var err error

	buf.WriteString(viper.GetString("mysql.DBuser"))
	buf.WriteString(":")
	buf.WriteString(viper.GetString("mysql.DBpasswd"))
	buf.WriteString("@tcp(")
	buf.WriteString(viper.GetString("mysql.DBaddr"))
	buf.WriteString(":")
	buf.WriteString(viper.GetString("mysql.DBport"))
	buf.WriteString(")/")
	buf.WriteString(viper.GetString("mysql.DBname"))
	buf.WriteString("?charset=utf8")

	golog.Info(buf.String())
	db, err = sql.Open("mysql", buf.String())
	if err != nil {
		golog.Fatal("mysql连接失败", "err", err)
	}

	//设置连接池
	db.SetMaxOpenConns(viper.GetInt("mysql.DBmaxconn"))
	db.SetMaxIdleConns(viper.GetInt("mysql.DBidleconn"))

	err = db.Ping()
	if err != nil {
		golog.Fatal("mysql ping失败", "err", err)
	}
	golog.Info("mysql连接成功")

	initService()

}

// DisconnectDB disconnects from the database.
func DisconnectDB() {
	if err := db.Close(); nil != err {
		golog.Error("Disconnect from database failed: ", "err", err.Error())
	}
}

var DBname string
var DBTable *tableService
var DBField *fieldService
var DBIndex *indexService

func initService() {
	DBname = viper.GetString("mysql.DBname")

	DBTable = &tableService{}
	DBField = &fieldService{}
	DBIndex = &indexService{}
}
