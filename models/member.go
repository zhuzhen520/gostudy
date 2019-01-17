package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Member struct {
	Id     int64   `orm:"pk;auto"`
	Name   string   `orm:"size(25)"`
	Nickname  string   `orm:"size(50)"`
	Password string
	Pid    int64
	Level  int8
	IsShow  int8
	AddTime time.Time `orm:"auto_now"`
}
func init (){
	orm.RegisterDataBase("","mysql","root:root@tcp(127.0.0.1:3306)/rbac")
	orm.RegisterModel(new(Member))
	//orm.RunSyncdb("default",false,true)
}