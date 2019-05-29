package {{ "" | WithPacket }}

import (
	"bytes"
	"database/sql"
	"github.com/spf13/viper"
	"github.com/wudaoluo/golog"
	_ "github.com/go-sql-driver/mysql"
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
		golog.Fatal("mysql连接失败", "err",err)
	}

	//设置连接池
	db.SetMaxOpenConns(viper.GetInt("mysql.DBmaxconn"))
	db.SetMaxIdleConns(viper.GetInt("mysql.DBidleconn"))

	err = db.Ping()
	if err != nil {
		golog.Fatal("mysql ping失败","err", err)
	}
	golog.Info("mysql连接成功")

}

// DisconnectDB disconnects from the database.
func DisconnectDB() {
	if err := db.Close(); nil != err {
		golog.Error("Disconnect from database failed: " ,"err", err.Error())
	}
}
