## 表总榄

| 表名 | 备注 |
|:----| :----|
{{range $j, $item := .}}| {{$item.Name }} |{{$item.Comment}}|
{{end}}

{{range $key, $item := .}}
### {{$item.Name}}
> {{$item.Comment}}

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
{{range $item.Fields}}|{{.Name}}  | {{.DataType}} | {{.IsNull}}| {{.Default }} | {{.Key}} | {{.Comment}}|
{{end}}

- 表索引

|名称 | 字段 | 备注|
|:----|:---|:----|
{{range $item.Indexs}}|{{.Name}}  | {{.Key}} | {{.Desc}}|
{{end}}
[TOP](#表总榄)
{{end}}