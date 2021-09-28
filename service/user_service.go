package service

import (
	"beego_example/models/base"
	"beego_example/models/model"
)

func QueryById(id int) (model.User, error) {
	one, err := model.QueryOne(id)
	return one, err
}

func SaveOrUpdate(u *model.User) (int64, error) {
	id, err := model.AddUser(u)
	return id, err
}

func QueryAll(pager base.Pager, user *model.User) ([]model.User, error) {
	users, err := model.QueryAll(pager, user)
	if err != nil {
		return nil, err
	}
	return users, err
}
