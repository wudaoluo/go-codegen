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
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/wudaoluo/go-codegen/generate"
	"github.com/wudaoluo/go-codegen/internal"
	"github.com/wudaoluo/go-codegen/model/mysql"
	"github.com/wudaoluo/go-codegen/util"
	"github.com/wudaoluo/golog"
)

type mysqlInfo struct {
	BasePath string
	Name     string
	Comment  string
	Fields   []*mysql.Field
}

// mysqlCmd represents the mysql command
var mysqlCmd = &cobra.Command{
	Use:   "mysql",
	Short: "生成mysql语句和markdown文档",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if mysqlF.Add == "" {
			util.Exit(1, "mysql -add not value")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		mysql.InitDB()
		defer mysql.DisconnectDB()

		var data = &mysqlInfo{BasePath: basePath()}
		g, err := getGen(mysqlType(mysqlF.Type), mysqlF.Add, data)
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

	mysqlCmd.Flags().StringVar(&mysqlF.Add, "add", "", "")
	mysqlCmd.Flags().StringVar(&mysqlF.Type, "type", "table", "table|doc")
}

func getGen(t mysqlType, tableName string, data *mysqlInfo) (generate.Generater, error) {
	var genType internal.Gen

	switch t {
	case MYSQL_DOC:
		genType = internal.GEN_MYSQL_DOC
	case MYSQL_TABLE:
		genType = internal.GEN_MYSQL_TABLE
		if tableName == "init" {
			t = MYSQL_INIT
			genType = internal.GEN_MYSQL_CONN
		}
	default:
		golog.Warn("输出类型不匹配", "type", t)
		return nil, errors.New("输出类型不匹配 请使用 table|doc")
	}

	err := t.Setdata(data)
	if err != nil {
		golog.Error("t.Setdata", "err", err)
		return nil, err
	}
	g := generate.Generate(genType)
	g.SetDest(destFile(tableName) + genType.FileSuffix())
	g.SetData(data)
	return g, nil
}

type mysqlType string

const (
	MYSQL_DOC   = "doc"
	MYSQL_INIT  = "init"
	MYSQL_TABLE = "table"
)

func (m mysqlType) Setdata(data *mysqlInfo) error {
	if m == MYSQL_DOC || m == MYSQL_INIT {
		return nil
	}

	tableInfo, err := mysql.DBTable.GetTable(mysqlF.Add)
	if err != nil {
		return err

	}
	fileds, err := mysql.DBField.GetFields(mysqlF.Add)
	if err != nil {
		return err
	}

	data.Name = tableInfo.Name
	data.Comment = tableInfo.Comment
	data.Fields = fileds

	return nil
}
