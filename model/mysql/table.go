package mysql

import (
	"github.com/wudaoluo/golog"
)

type Table struct {
	Name    string
	Comment string
}

type tableService struct {
	dbName string
}

/*
	row := db.QueryRow("select id,user_name,weixin,email,status from user where id = ?",id)
	u := &User{}
	err := row.Scan(&u.Id,&u.UserName,&u.Weixin,&u.Email,&u.Status)
	if err != nil{
		//只有当查询的结果为空的时候，会触发一个sql.ErrNoRows错误
		if err == sql.ErrNoRows{
			return nil,fmt.Errorf("user not found id: ",id)
		}else {
			return nil,err
		}
	}

	return u,nil
*/
func (t *tableService) GetTable(tableName string) (*Table, error) {
	sqlText := "SELECT `table_name`, table_comment FROM information_schema.tables " +
		"WHERE table_schema = ? AND  `table_name` = ?"
	row := db.QueryRow(sqlText, DBname, tableName)

	msg := new(Table)
	err := row.Scan(&msg.Name, &msg.Comment)
	if err != nil {
		golog.Error("GetTable", "table", tableName, "err", err)
	}

	return msg, err
}
