package models

type User struct {
	Id       string `json:"id"`
	Nickname string `json:"nickname"`
	Passwd   string `json:"passwd"`
}
