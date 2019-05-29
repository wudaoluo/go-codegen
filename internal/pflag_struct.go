package internal

type FlagRoot struct {
	Debug   bool   `json:"debug"`
	OutPath string `json:"out_path"`
}

type FlagUpdate struct {
	List bool `json:"list"`
}

type FlagMysql struct {
	Add     string `json:"add"`
	Doc     bool `json:"doc"`
	Context bool   `json:"context"` //true:使用Context包;false:不使用,默认不使用
}
