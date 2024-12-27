package entity

type User struct {
	UserId        string `json:"userId"`
	UserName      string `json:"username"`
	NickName      string `json:"nickName"`
	Gender        string `json:"gender"`
	Phone         string `json:"phone"`
	LastLoginTime int64  `json:"lastLoginTime"`
	CreateTime    int64  `json:"createTime"`
	UpdateTime    int64  `json:"updateTime"`
	OpenId        string `json:"openId"`
}
