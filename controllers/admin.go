package controllers

import (
	"beego-project/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type AdminController struct {
	beego.Controller
}

func (this *AdminController) JsonEncode(code int, msg string, data interface{}, count int) {
	this.Data["json"] = map[string]interface{}{"code": code, "msg": msg, "data": data, "count": count}
	this.ServeJSON()
	this.StopRun()
}

/*会员列表*/
func (this *AdminController) MemberList() {
	this.TplName = "member/userlist.html"
}
func (this *AdminController) MemberListJson() {
	o := orm.NewOrm()
	var member []models.Member
	lim, _ := this.GetInt("limit")
	page, _ := this.GetInt("page")
	page = (page - 1) * lim
	_, err := o.QueryTable("member").Limit(lim, page).All(&member)
	if err != nil {
		fmt.Println(err)
	}
	count, err := o.QueryTable("member").Count()
	this.JsonEncode(0, "success", member, int(count))
}

/*添加用户*/
func (this *AdminController) MemberAdd() {
	o := orm.NewOrm()
	id, _ := this.GetInt64("id")
	name := this.GetString("name")
	nickname := this.GetString("nickname")
	password := this.GetString("password")
	is_show, _ := this.GetInt8("is_show")
	level, _ := this.GetInt8("level")
	pid, _ := this.GetInt64("pid")
	var member = models.Member{}

	if id != 0 {
		member.Id = id
		err := o.Read(&member, "Id")
		if err != nil {
			this.JsonEncode(0, "查询失败", nil, 0)
		}
		member.Name = name
		member.Nickname = nickname
		member.Password = password
		member.Pid = pid
		member.IsShow = is_show
		member.Level = level
		_,err1 := o.Update(&member);
		if err1 != nil {
			this.JsonEncode(2, "更新失败", nil, 0)
		}
		this.JsonEncode(1, "更新成功", nil, 0)
	} else {
		member.Name = name
		member.Nickname = nickname
		member.Password = password
		member.Pid = pid
		member.IsShow = is_show
		member.Level = level
		_, err := o.Insert(&member)
		if err != nil {
			this.JsonEncode(2, "添加失败", nil, 0)
		}
		this.JsonEncode(1, "添加成功", nil, 0)

	}
}

/*删除用户*/
func (this *AdminController) MemberDelete() {
	o := orm.NewOrm()
	id, _ := this.GetInt64("id")
	var member = models.Member{Id: id}
	//  member.Id = id
	_, err := o.Delete(&member, "Id")
	if err != nil {
		this.JsonEncode(2, "删除失败", nil, 0)
	}
	this.JsonEncode(1, "删除成功", nil, 0)
}
