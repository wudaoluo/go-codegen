package {{.BasePath}}


{{.Comment}}
type {{.Name | WithTitle}} struct {
{{range $j, $item := .Fields}}{{$item.Name | WithTitle}} {{$item.DataType | TypeToGo}}  {{$item.Comment | WithComment}}
{{end}}
}




type {{.Name | WithNotFirstTitle}}Service struct {}

var DB{{.Name | WithTitle}} *{{.Name | WithNotFirstTitle}}Service
func init() {
   DB{{.Name | WithTitle}} = &{{.Name | WithNotFirstTitle}}Service{}
}



func (t *{{.Name | WithNotFirstTitle}}Service) SelectById(id int64) (*{{.Name | WithTitle}},error) {
    sqlText := "SELECT  {{range $j, $item := .Fields}}{{$item.Name }},{{end}} FROM {{.Name}} WHERE id = ? limit 1"
    row := db.QueryRow(sqlText,id)

    msg := new({{.Name | WithTitle}})
    err := row.Scan(
        {{range $j, $item := .Fields}}&msg.{{$item.Name | WithTitle}},
        {{end}})
    if err != nil {
        return nil,err
    }

    return msg,nil
}


func (t *{{.Name | WithNotFirstTitle}}Service) Select() ([]*{{.Name | WithTitle}},error) {
    return nil,nil
}

func (t *{{.Name | WithNotFirstTitle}}Service) Insert(msg *{{.Name | WithTitle}}) (int64,error) {
    sqlText := "INSERT INTO {{.Name}} ({{range $j, $item := .Fields}}{{if ne  $j 0}} {{$item.Name }},{{end}} {{end}}) " +
        "VALUE ({{range $j, $item := .Fields}}{{if ne  $j 0}} ?,{{end}} {{end}})"

    ret, err := db.Exec(sqlText,
            {{range $i, $item := .Fields}}
            {{if ne  $i 0}}&msg.{{$item.Name | WithTitle}},{{end}}{{end}})

    if err != nil {
        return 0,err
    }
    return ret.LastInsertId()
}

func (t *{{.Name | WithNotFirstTitle}}Service) DeleteById (id int64) error {
    sqlText := "DELETE FROM {{.Name}} where id = ?"
    _, err := db.Exec(sqlText,id)
    return err
}

func (t *{{.Name | WithNotFirstTitle}}Service) UpdateById (msg *{{.Name | WithTitle}}) error {
    sqlText := "UPDATE {{.Name}} SET {{range $j, $item := .Fields}}{{if ne  $j 0}}{{$item.Name }}=?,{{end}} {{end}} WHERE id = ?"
	_, err := db.Exec(sqlText,{{range $j, $item := .Fields}}
	{{if ne  $j 0}}msg.{{$item.Name | WithTitle}},{{end}} {{end}}
	msg.Id)
	return err
}