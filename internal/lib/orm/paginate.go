package orm

import "github.com/vulcangz/gf2-authz/internal/model/entity"

type Paginated[T entity.Models] struct {
	Data  []*T  `json:"data"`
	Total int64 `json:"total"`
	Page  int64 `json:"page"`
	Size  int64 `json:"size"`
}

func NewPaginated[T entity.Models](data []*T, total, page, size int64) *Paginated[T] {
	return &Paginated[T]{
		Data:  data,
		Total: total,
		Page:  page,
		Size:  size,
	}
}
