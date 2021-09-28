package model

import (
	"beego_example/constants"
	"beego_example/models/base"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

type User struct {
	Id       int    `orm:"pk;auto;int(32);description(主键)" json:"id"`
	UserName string `orm:"column(user_name);description(用户名)" json:"userName"`
	UserPass string `orm:"size(32);description(密码)" json:"userPass"`
	UserAge  int    `orm:"default(18);description(年龄)" json:"userAge"`
}

func init() {
	orm.RegisterModel(new(User))
}

func (u *User) TableName() string {
	return constants.UserTable
}

func AddUser(u *User) (int64, error) {
	o := orm.NewOrm()
	id, err := o.InsertOrUpdate(u, u.UserName)
	logs.Info("data : ", id)
	if err != nil {
		logs.Error("Insert Or Update ", err)
		return 0, err
	}
	return id, nil
}

func QueryOne(id int) (user User, err error) {
	o := orm.NewOrm()
	err = o.QueryTable(user).Filter("id", id).One(&user)
	if err != nil {
		logs.Error("query id error ", err.Error())
		return User{}, err
	}
	return user, nil
}

func QueryAll(pager base.Pager, user *User) (users []User, err error) {
	table := orm.NewOrm().QueryTable(new(User))
	if len(user.UserName) > 0 {
		table = table.Filter("user_name", user.UserName)
	}
	_, err = table.Limit(pager.PageSize, pager.Offset()).All(&users)
	if err != nil {
		logs.Error("query all error", err)
	}
	return users, err
}
