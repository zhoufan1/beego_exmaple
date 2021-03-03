package models

import "math"

type Pager struct {
	PageNo   int `json:"pageNo"`   //页数
	PageSize int `json:"pageSize"` //条数
	Total    int `json:"total"`    //总条数
	Pages    int `json:"pages"`    //总条数
}

func (p *Pager) Offset() int {
	return (p.PageNo - 1) * p.PageSize
}

func (p *Pager) Size() int {
	if p.PageSize == 0 {
		p.PageSize = 10
	}
	return p.PageSize
}

func (p *Pager) SetTotal(total int) {
	p.Total = total
	p.Pages = int(math.Ceil(float64(total) / float64(p.Size())))
}

type Model struct {
	Code    int         `json:"code"`    //错误码
	Message string      `json:"message"` //消息
	Data    interface{} `json:"data"`    //数据
}

type ModelPager struct {
	Records interface{} `json:"records"` //数据
	Pager   Pager       `json:"pager"`   //分页
}

func Response(code int, message string, data interface{}) *Model {
	m := new(Model)
	m.Code = code
	m.Message = message
	m.Data = data
	return m
}

func ResponseSuccess(data interface{}) *Model {
	return Response(0, "success", data)
}

func ResponseSuccessPager(data interface{}, pager Pager) *Model {
	m := new(ModelPager)
	m.Pager = pager
	m.Records = data
	return Response(0, "success", m)
}

func ResponseFailure(code int, message string) *Model {
	return Response(code, message, nil)
}
