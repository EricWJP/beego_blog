package helpers

import (
	"github.com/astaxie/beego"
	"html/template"
	"time"
)

func init() {
	beego.AddFuncMap("timeFormat", timeFormat)
	//beego.AddFuncMap("compileTemplate", compileTemplate)
}

func timeFormat(time time.Time) string {
	return time.Format("2006-01-02 07:00")
}

func compileTemplate(name string) *template.Template {
	tpl := template.New(name)
	//tpl, err := tpl.ParseFiles(
	//	"views/layouts/"+layout+".htm",
	//	"views/"+name+".htm",
	//)
	//if err != nil {
	//	return nil, err
	//}
	return tpl
}
