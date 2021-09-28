package request

type UserPagerRequest struct {
	UserName string `json:"userName"` //用户名
	PageNo   int    `json:"pageNo"`   //页数
	PageSize int    `json:"pageSize"` //条数
}
