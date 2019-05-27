package mysql

import (
	"github.com/wudaoluo/golog"
)

type Field struct {
	Name     string      //数据库原始字段
	DataType string      //数据库原始类型
	Key      string      //索引
	IsNull   string      //是否为空
	Default  interface{} //默认值
	Comment  string      //备注
}

type FieldService struct{}

func (f *FieldService) GetFields(tableName string) ([]*Field, error) {
	sqlText := "SELECT column_name,data_type, column_key, is_nullable,column_default, column_comment " +
		"FROM information_schema.columns WHERE table_schema = ? and table_name = ?"

	rows, err := db.Query(sqlText, DBname, tableName)
	if err != nil {
		golog.Error("GetFields", "table", tableName, "err", err)
		return nil, err
	}

	defer rows.Close()

	list := make([]*Field, 0)
	for rows.Next() {
		msg := new(Field)
		err = rows.Scan(&msg.Name, &msg.DataType, &msg.Key, &msg.IsNull, &msg.Default, &msg.Comment)
		if err != nil {
			golog.Error("GetFields", "table", tableName, "err", err)
			return nil, err
		}
		list = append(list, msg)
	}

	//为了检查是否是迭代正常退出还是异常退出，需要检查rows.Err
	if rows.Err() != nil {
		golog.Error("GetFields", "table", tableName, "err", err)
		return nil, err
	}

	return list, nil

}
