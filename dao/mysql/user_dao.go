package mysql

import (
	"database/sql"
	"errors"
	"gin_web_demo/models"
	"gin_web_demo/models/dto"
	"gin_web_demo/pkg/snowflake"
	"gin_web_demo/utils"
)

func GetUserList(dto *dto.UserSearchDto) (userList []*dto.UserSimpleDto, err error) {
	sql := "select user_id,username,email,gender from user where id > 0"
	if dto.UserName != "" {
		sql += " and username like '%" + dto.UserName + "%'"
	}
	if dto.Email != "" {
		sql += " and email like '%" + dto.Email + "%'"
	}
	sql += " limit ?,?"
	err = db.Select(&userList, sql, (dto.Page-1)*dto.PageSize, dto.PageSize)
	return
}

func GetUserByUserId(userId int64) (user *models.User, err error) {
	sqlStr := "select * from user where user_id = ?"
	user = new(models.User)
	if err = db.Get(user, sqlStr, userId); err != nil {
		if err == sql.ErrNoRows {
			err = errors.New("无效的userId")
		}
	}
	return
}

func UpdateUser(dto *dto.UpdateUserDto) (err error) {
	sql := `UPDATE user SET 
					password = ?, 
					email = ?, 
					gender = ?, 
					create_time = NOW(), 
					update_time = now() 
					WHERE user_id = ?`
	password := utils.EncryptPassword(dto.Password)
	if _, err := db.Exec(sql, password, dto.Email, dto.Gender, dto.UserId); err != nil {
		return err
	}
	return
}

func CreateUser(dto *dto.CreateUserDto) (err error) {
	sql := `INSERT INTO user (user_id, username, password, email, gender, create_time, update_time) 
	VALUES (?, ?, ?, ?, ?, now(), now());`
	uid := snowflake.GenID()
	password := utils.EncryptPassword(dto.Password)
	if _, err := db.Exec(sql, uid, dto.UserName, password, dto.Email, dto.Gender); err != nil {
		return err
	}
	return
}

func DeleteUserByUserId(userId int64) (err error) {
	sql := "delete from user where user_id = ?"
	if _, err = db.Exec(sql, userId); err != nil {
		return err
	}
	return
}

func Login(dto *dto.LoginDto) (userId int64, err error) {
	sqlStr := "select user_id from user where username = ? and password = ?"
	password := utils.EncryptPassword(dto.Password)
	if err := db.Get(&userId, sqlStr, dto.UserName, password); err != nil {
		if err == sql.ErrNoRows {
			err = utils.ErrorInvalidPassword
		}
		return 0, err
	}
	return
}
