package generate

import (
	"fmt"
	"github.com/wudaoluo/go-codegen/internal"
)


func (g *Generate)WithImport(args string) string {
	if args != "import" {
		return ""
	}

	var iPacket string
	for k,_ := range g.iType {
		fmt.Println(k)
		iPacket += "\""+k+"\""
	}
	var a = "import (%s)\n"
	return fmt.Sprintf(a,iPacket)

}

func (g *Generate)TypeToGo(t string) string {
	 v,ok := typeToGo[t]
	 if !ok {
		return internal.GoDefaultType
	}

	switch v {
	case "time.Time":
		g.iType["time"] = struct {}{}
	}

	return v
}

//mysql类型 <=> golang类型
var typeToGo = map[string]string{
	"tinyint":    "int64",
	"smallint":   "int64",
	"mediumint":  "int64",
	"int":        "int64",
	"integer":    "int64",
	"bigint":     "int64",
	"float":      "float64",
	"double":     "float64",
	"decimal":    "float64",
	"date":       "string",
	"time":       "string",
	"year":       "string",
	"datetime":   "time.Time",
	"timestamp":  "time.Time",
	"char":       "string",
	"varchar":    "string",
	"tinyblob":   "string",
	"tinytext":   "string",
	"blob":       "string",
	"text":       "string",
	"mediumblob": "string",
	"mediumtext": "string",
	"longblob":   "string",
	"longtext":   "string",
}