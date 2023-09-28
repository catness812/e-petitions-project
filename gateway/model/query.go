package model

type PaginationQuery struct {
	Page  uint32 `form:"page"`
	Limit uint32 `form:"limit"`
}
