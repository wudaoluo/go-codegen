package generate

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"

	"github.com/wudaoluo/go-codegen/internal"
	"github.com/wudaoluo/goutil"
)

type Generate struct {
	genType internal.Gen
	Src     string //tempalte name
	Dest    string //generate name
	funcMap template.FuncMap
	Data    interface{}
	iType   map[string]struct{}
}

func (g *Generate) Gen() error {
	var err error
	if !goutil.FileIsExist(g.Src) {
		err = &internal.TPlError{
			Op: "FileIsExist", TplName: g.Src, Err: internal.ERROR_TPL_NOT_FOUND,
		}
		return err
	}

	tmpl, err := template.New(filepath.Base(g.Src)).Funcs(g.funcMap).ParseFiles(g.Src)
	if err != nil {
		err = &internal.TPlError{
			Op: "ParseFiles", TplName: g.Src, Err: err,
		}
		return err
	}

	// create TempFile in Dest directory to avoid cross-filesystem issues
	temp, err := ioutil.TempFile(filepath.Dir(g.Dest), filepath.Base(g.Dest))
	if err != nil {
		err = &internal.TPlError{
			Op: "TempFile", TplName: g.Src, Err: err,
		}
		return err
	}

	if err = tmpl.Execute(temp, g.Data); err != nil {
		temp.Close()
		os.Remove(temp.Name())
		return err
	}

	defer temp.Close()

	err = os.Rename(temp.Name(), g.Dest)

	if err != nil {
		return err
	}

	return nil
}

func (g *Generate) SetTpl(tpl string) {
	g.Src = tpl
}

func (g *Generate) SetDest(dest string) {
	g.Dest = dest
	fmt.Println(g.Dest)
}

func (g *Generate) SetData(data interface{}) {
	g.Data = data
}

func (g *Generate) AddFuncMap() {
	if g.funcMap == nil {
		g.funcMap = template.FuncMap{}
	}

	if g.iType == nil {
		g.iType = make(map[string]struct{})
	}

	g.funcMap["WithComment"] = g.WithComment
	g.funcMap["TypeToGo"] = g.TypeToGo
	g.funcMap["WithTitle"] = g.WithTitle
	g.funcMap["WithNotFirstTitle"] = g.WithNotFirstTitle
	g.funcMap["WithImport"] = g.WithImport
}
