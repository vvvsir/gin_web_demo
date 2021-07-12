package dto

type PageDto struct {
	Page     int64 `form:"page" json:"page" binding:"required"`
	PageSize int64 `form:"pageSize" json:"pageSize" binding:"required"`
}
