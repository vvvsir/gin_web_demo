package dto

type UserSimpleDto struct {
	UserId   int64  `db:"user_id" json:"userId"`    //用户ID
	UserName string `db:"username" json:"username"` //用户名
	Email    string `db:"email" json:"email"`       //邮箱
	Gender   int    `db:"gender" json:"gender"`     //性别
}

type UpdateUserDto struct {
	UserId   int64  `form:"userId" json:"userId"`
	Password string `form:"password" json:"password"`
	Email    string `form:"email" json:"email"`
	Gender   int    `form:"gender" json:"gender"`
}

type CreateUserDto struct {
	UserName string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
	Email    string `form:"email" json:"email"`
	Gender   int    `form:"gender" json:"gender"`
}

type LoginDto struct {
	UserName string `form:"username" json:"username" binding:"required"` //用户名
	Password string `form:"password" json:"password" binding:"required"` //密码
}

type UserSearchDto struct {
	UserName string `db:"username" json:"username"`
	Email    string `db:"email" json:"email"`
	*PageDto
}
