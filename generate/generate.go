package generate

import (
	"github.com/wudaoluo/go-codegen/generate/generate"
	"github.com/wudaoluo/go-codegen/generate/mysql"
	"github.com/wudaoluo/go-codegen/internal"
	"github.com/wudaoluo/golog"
)

type Generater interface {
	Gen() error //生成
	SetTpl(tpl string)
	SetDest(dest string)
	SetData(data interface{})
	SetPacket(outPath string)
	AddFuncMap()
}

func Generate(genType internal.Gen) Generater {
	var g Generater
	switch genType {
	case internal.GEN_MYSQL_CONN:
		g = &mysql.MysqlConn{&generate.Generate{}}

	case internal.GEN_MYSQL_TABLE:
		g = &mysql.MysqlTable{&generate.Generate{}}

	case internal.GEN_MYSQL_DOC:
		g = &mysql.MysqlDoc{&generate.Generate{}}

	default:
		golog.Error("Generate 不匹配", "genType", genType)
		panic(genType)
	}

	g.AddFuncMap()

	g.SetTpl(genType.TplFile())
	return g
}
