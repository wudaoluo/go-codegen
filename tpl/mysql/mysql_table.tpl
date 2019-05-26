package {{.BasePath}}


{{.Comment}}
type {{.Name | WithTitle}} struct {
{{range $j, $item := .Fields}}{{$item.Name}} {{$item.DataType | TypeToGo}}  {{$item.Comment | WithComment}}
{{end}}
}




type {{.Name | WithNotFirstTitle}}Service struct {}

var DB{{.Name | WithTitle}} *{{.Name | WithNotFirstTitle}}Service
func init() {
   DB{{.Name | WithTitle}} = &{{.Name | WithNotFirstTitle}}Service{}
}