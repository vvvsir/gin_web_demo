package models

import (
	"github.com/golang-module/carbon"
)

type User struct {
	Id         int64                   `db:"id" json:"id"`
	UserId     int64                   `db:"user_id" json:"userId"`
	UserName   string                  `db:"username" json:"username"`
	Password   string                  `db:"password" json:"password"`
	Email      string                  `db:"email" json:"email"`
	Gender     int                     `db:"gender" json:"gender"`
	CreateTime carbon.ToDateTimeString `json:"createTime" db:"create_time"`
	UpdateTime carbon.ToDateTimeString `json:"updateTime" db:"update_time"`
}
