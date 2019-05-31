// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/wudaoluo/go-codegen/generate"
	"github.com/wudaoluo/go-codegen/internal"
	"github.com/wudaoluo/go-codegen/model/mysql"
	"github.com/wudaoluo/go-codegen/util"
	"github.com/wudaoluo/golog"
)

type mysqlInfo struct {
	Name     string
	Comment  string
	Fields   []*mysql.Field
	Indexs   []*mysql.Index
}

// mysqlCmd represents the mysql command
var mysqlCmd = &cobra.Command{
	Use:   "mysql",
	Short: "生成mysql语句和markdown文档",
	Long: ``,
	PreRun: func(cmd *cobra.Command, args []string) {
		if mysqlF.Add == "" && !mysqlF.Doc {
			util.Exit(1, "mysql -add not value")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		mysql.InitDB()
		defer mysql.DisconnectDB()

		g, err := getGen(getMysqlType(mysqlF.Doc, mysqlF.Add), mysqlF.Add)
		if err != nil {
			util.Exit(1, err.Error())
		}
		err = g.Gen()
		if err != nil {
			util.Exit(1, err.Error())
		}
	},
}

var mysqlF = new(internal.FlagMysql)

func init() {
	rootCmd.AddCommand(mysqlCmd)

	mysqlCmd.Flags().StringVar(&mysqlF.Add, "add", "", "--add 表名")
	mysqlCmd.Flags().BoolVar(&mysqlF.Doc, "doc", false, "--doc")
	mysqlCmd.Flags().BoolVar(&mysqlF.Prepare,"prepare",false,"--prepare")
}

func getGen(t mysqlType, tableName string) (generate.Generater, error) {
	var genType internal.Gen
	switch t {
	case MYSQL_DOC:
		genType = internal.GEN_MYSQL_DOC
		tableName = mysql.DBname
	case MYSQL_INIT:
		genType = internal.GEN_MYSQL_CONN
	case MYSQL_TABLE:
		genType = internal.GEN_MYSQL_TABLE
		if mysqlF.Prepare {
			genType = internal.GEN_MYSQL_TABLE_PREPARE
		}

	default:
		golog.Warn("输出类型不匹配", "type", t)
		return nil, errors.New("输出类型不匹配 请使用 table|doc")
	}

	data,err := t.GetData()
	if err != nil {
		golog.Error("t.Setdata", "err", err)
		return nil, err
	}
	g := generate.Generate(genType)
	g.SetDest(destFile(tableName) + genType.FileSuffix())
	g.SetPacket(rootF.OutPath)
	g.SetData(data)
	return g, nil
}

type mysqlType string

const (
	MYSQL_DOC   mysqlType = "doc"
	MYSQL_INIT  mysqlType = "init"
	MYSQL_TABLE mysqlType = "table"
)

func getMysqlType(mysqlDoc bool, tableName string) mysqlType {
	if mysqlDoc {
		return MYSQL_DOC
	}

	if tableName == string(MYSQL_INIT) {
		return MYSQL_INIT
	}

	return MYSQL_TABLE

}

func (m mysqlType) GetData() (interface{},error) {
	if m == MYSQL_INIT {
		return nil,nil
	}

	if m == MYSQL_TABLE {
		d:= new(mysqlInfo)
		tableInfo, err := mysql.DBTable.GetTable(mysqlF.Add)
		if err != nil {
			return nil,err

		}
		fileds, err := mysql.DBField.GetFields(mysqlF.Add)
		if err != nil {
			return nil,err
		}

		d.Name = tableInfo.Name
		d.Comment = tableInfo.Comment
		d.Fields = fileds
		return d,nil
	}

	if m == MYSQL_DOC {
		tableInfoList := []*mysqlInfo{}


		tableList,err := mysql.DBTable.TableList()
		if err != nil {
			util.Exit(1, err.Error())
		}
		for _,list := range tableList{
			var tableInfo = new(mysqlInfo)

			tableInfo.Name = list.Name
			tableInfo.Comment = list.Comment
			fileds, err := mysql.DBField.GetFields(tableInfo.Name)
			if err != nil {
				return nil,err
			}

			tableInfo.Fields = fileds

			indexs,err := mysql.DBIndex.GetIndexs(tableInfo.Name)
			if err != nil {
				return nil, err
			}

			tableInfo.Indexs = indexs

			tableInfoList = append(tableInfoList,tableInfo)
		}

		fmt.Println(tableList)
		return tableInfoList,nil

	}

	return nil, nil
}
