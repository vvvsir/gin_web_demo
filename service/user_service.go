package service

import (
	"gin_web_demo/dao/mysql"
	"gin_web_demo/dao/redis"
	"gin_web_demo/models"
	"gin_web_demo/models/dto"
)

func GetUserList(dto *dto.UserSearchDto) (userList []*dto.UserSimpleDto, err error) {
	return mysql.GetUserList(dto)
}

func GetUserByUserId(userId int64) (user *models.User, err error) {
	user, err = redis.GetUserByUserId(userId)
	if err != nil {
		user, err = mysql.GetUserByUserId(userId)
		if err == nil {
			redis.SetUserByUserId(user)
		}
	}

	return
}

func UpdateUser(dto *dto.UpdateUserDto) (err error) {
	return mysql.UpdateUser(dto)
}

func CreateUser(dto *dto.CreateUserDto) (err error) {
	return mysql.CreateUser(dto)
}

func DeleteUserByUserId(userId int64) (err error) {
	return mysql.DeleteUserByUserId(userId)
}

func Login(dto *dto.LoginDto) (userId int64, err error) {
	return mysql.Login(dto)
}
