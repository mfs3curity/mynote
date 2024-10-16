package dto

type TokenDetail struct {
	Token  string
	Expire int64
}
type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
