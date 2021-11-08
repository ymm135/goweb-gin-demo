package common

type UserInfo struct {
	ID   uint   `json:"id" form:"id"`
	Name string `json:"name"`
}
