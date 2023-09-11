package model

// 用户信息
type UserInfo struct {
	Id         int    `json:"id"`
	AppPubKey  string `json:"appPubKey"`
	Email      string `json:"email"`
	LoginType  string `json:"loginType"`
	Address    string `json:"address"`
	KeyHash    string `json:"keyHash"`
	CreateTime int64  `json:"create_time"`
}
