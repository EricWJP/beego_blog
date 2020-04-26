package helpers

import (
	"github.com/astaxie/beego"
	"time"
)

func init() {
	_ = beego.AddFuncMap("timeFormat", timeFormat)
	_ = beego.AddFuncMap("add", add)
}

func timeFormat(time time.Time) string {
	return time.Format("2006-01-02 07:00")
}

func add(base, num int64) int64 {
	return base + num
}
