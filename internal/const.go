package internal

const (
	CONFIG_NAME  = "codegen.yaml"
	PROJECT_NAME = "go-codegen"
	VERSION      = "version 0.1"
)

type Gen int

const (
	GEN_MYSQL_CONN Gen = iota + 1
	GEN_MYSQL_TABLE
	GEN_MYSQL_DOC
)

const (
	SUFFIX_DOC = ".doc"
	SUFFIX_GO  = ".go"
)

func (g Gen) String() string {
	switch g {
	case GEN_MYSQL_CONN:
		return "mysql init"
	case GEN_MYSQL_TABLE:
		return "mysql add table"
	case GEN_MYSQL_DOC:
		return "mysql add doc"
	default:
		return ""
	}

}

func (g Gen) TplFile() string {
	switch g {
	case GEN_MYSQL_CONN:
		return TPL_MYSQL_CONN

	case GEN_MYSQL_TABLE:
		return TPL_MYSQL_TABLE

	case GEN_MYSQL_DOC:
		return TPL_MYSQL_DOC

	default:
		return ""
	}
}

func (g Gen) FileSuffix() string {
	var suffix string
	switch g {
	case GEN_MYSQL_DOC:
		suffix = SUFFIX_DOC
	default:
		suffix = SUFFIX_GO
	}

	return suffix
}

const (
	TPL_MYSQL_CONN  = "tpl/mysql/mysql_conn.tpl"
	TPL_MYSQL_TABLE = "tpl/mysql/mysql_table.tpl"
	TPL_MYSQL_DOC   = "tpl/mysql/mysql_doc.tpl"
)

const GoDefaultType = "interface"
