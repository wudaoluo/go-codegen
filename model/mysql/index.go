package mysql

import (
	"github.com/wudaoluo/golog"
)
type Index struct {
	Name     string      //索引名称
	Desc string
	Key     string      //包含的字段
}

type indexService struct{}

func (i *indexService) GetIndexs(tableName string) ([]*Index,error) {
	sqlText := "SELECT index_name,stat_name,stat_description FROM mysql.innodb_index_stats WHERE database_name = ?  AND `table_name` = ? and stat_name like 'n_diff_pfx%'"

	rows, err := db.Query(sqlText, DBname, tableName)
	if err != nil {
		golog.Error("GetIndexs", "table", tableName, "err", err)
		return nil, err
	}

	defer rows.Close()

	list := make([]*Index, 0)
	for rows.Next() {
		msg := new(Index)
		err = rows.Scan(&msg.Name, &msg.Desc, &msg.Key)
		if err != nil {
			golog.Error("GetIndexs", "table", tableName, "err", err)
			return nil, err
		}

		list = append(list, msg)
	}

	//为了检查是否是迭代正常退出还是异常退出，需要检查rows.Err
	if rows.Err() != nil {
		golog.Error("GetIndexs", "table", tableName, "err", err)
		return nil, err
	}

	return list, nil
}


