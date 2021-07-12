package controller

import (
	"errors"
	"fmt"
	"gin_web_demo/models/dto"
	"gin_web_demo/pkg/jwt"
	"gin_web_demo/service"
	"gin_web_demo/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// @Summary 获取用户列表（分页）
// @Tags 用户模块
// @Accept application/json
// @Param data body dto.UserSearchDto true "请求参数结构体"
// @Success 200 object utils.ResponseData 返回值
// @Router /getusers [post]
func GetUserList(c *gin.Context) {
	// 方式1
	search := dto.UserSearchDto{}
	if err := c.ShouldBind(&search); err != nil {
		zap.L().Error("dto.UserSearchDto param faild", zap.Error(err))
		utils.ResponseError(c, utils.CodeInvalidParam)
		return
	}
	userList, err := service.GetUserList(&search)

	// 方式2
	// search := new(dto.UserSearchDto)
	// if err := c.ShouldBind(search); err != nil {
	// 	zap.L().Error("dto.UserSearchDto param faild", zap.Error(err))
	//	utils.ResponseError(c, utils.CodeInvalidParam)
	// 	return
	// }
	// userList, err := service.GetUserList(search)

	if err != nil {
		zap.L().Error("service.GetUserList(&search) failed", zap.Error(err))
		utils.ResponseError(c, utils.CodeInvalidParam)
		return
	}
	utils.ResponseSuccess(c, userList)
}

// @Summary 获取单个用户信息
// @Tags 用户模块
// @Accept application/json
// @Param Authorization header string true "登录信息"
// @Param userId path integer true "用户ID"
// @Success 200 object utils.ResponseData 返回值
// @Router /getuser/{userId} [get]
func GetUserByUserId(c *gin.Context) {
	userId, _ := strconv.ParseInt(c.Param("userId"), 10, 64)
	user, err := service.GetUserByUserId(userId)
	if err != nil {
		zap.L().Error("service.GetUserByUserId(userId) failed", zap.Error(err))
		utils.ResponseError(c, utils.CodeInvalidParam)
		return
	}
	utils.ResponseSuccess(c, user)
}

// @Summary 更新用户信息
// @Tags 用户模块
// @Accept application/json
// @Param data body dto.UpdateUserDto true "请求参数结构体"
// @Success 200 object utils.ResponseData 返回值
// @Router /updateuser [post]
func UpdateUser(c *gin.Context) {
	var user dto.UpdateUserDto
	if err := c.Bind(&user); err != nil {
		zap.L().Error("dto.UpdateUserDto param failed", zap.Error(err))
		utils.ResponseError(c, utils.CodeInvalidParam)
		return
	}
	if err := service.UpdateUser(&user); err != nil {
		zap.L().Error("service.UpdateUser(&user) failed", zap.Error(err))
		utils.ResponseError(c, utils.CodeInvalidParam)
		return
	}
	utils.ResponseSuccess(c, nil)
}

// @Summary 新创建用户信息
// @Tags 用户模块
// @Accept application/json
// @Param data body dto.CreateUserDto true "请求参数结构体"
// @Success 200 object utils.ResponseData 返回值
// @Router /createuser [post]
func CreateUser(c *gin.Context) {
	user := new(dto.CreateUserDto)
	if err := c.ShouldBind(user); err != nil {
		zap.L().Error("dto.CreateUserDto param faild", zap.Error(err))
		utils.ResponseError(c, utils.CodeInvalidParam)
		return
	}
	if err := service.CreateUser(user); err != nil {
		zap.L().Error("service.CreateUser(user) faild", zap.Error(err))
		utils.ResponseError(c, utils.CodeInvalidParam)
		return
	}
	utils.ResponseSuccess(c, nil)
}

// @Summary 删除指定用户
// @Tags 用户模块
// @Accept application/json
// @Param userId path integer true "用户ID"
// @Success 200 object utils.ResponseData 返回值
// @Router /deleteuser/{userId} [get]
func DeleteUserByUserId(c *gin.Context) {
	userId, _ := strconv.ParseInt(c.Param("userId"), 10, 64)
	if err := service.DeleteUserByUserId(userId); err != nil {
		zap.L().Error("service.DeleteUserByUserId(userId) faild", zap.Error(err))
		utils.ResponseError(c, utils.CodeInvalidParam)
		return
	}
	utils.ResponseSuccess(c, nil)
}

// @Summary 用户登录
// @Tags 用户模块
// @Accept application/json
// @Param data body dto.LoginDto true "请求参数结构体"
// @Success 200 object utils.ResponseData 返回值
// @Router /login [post]
func Login(c *gin.Context) {
	dto := new(dto.LoginDto)
	if err := c.ShouldBind(dto); err != nil {
		zap.L().Error("dto.LoginDto param failed", zap.Error(err))
		// 判断err是不是validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			utils.ResponseError(c, utils.CodeInvalidParam)
			return
		}
		utils.ResponseErrorWithMsg(c, utils.CodeInvalidParam, utils.RemoveTopStruct(errs.Translate(utils.Trans)))
		return
	}
	userId, err := service.Login(dto)
	if err != nil {
		zap.L().Error("service.Login(&dto) faild", zap.String("username", dto.UserName), zap.Error(err))
		if errors.Is(err, utils.ErrorInvalidPassword) {
			utils.ResponseError(c, utils.CodeInvalidPassword)
			return
		}
		utils.ResponseError(c, utils.CodeInvalidPassword)
		return
	}

	// 生成JWT
	token, err := jwt.GenToken(userId, dto.UserName, "gin_login")
	if err != nil {
		return
	}

	utils.ResponseSuccess(c, gin.H{
		"userId":   fmt.Sprintf("%d", userId), // id值大于1<<53-1  int64类型的最大值是1<<63-1
		"userName": dto.UserName,
		"token":    token,
	})
}
