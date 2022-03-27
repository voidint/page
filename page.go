package page

import (
	"golang.org/x/exp/constraints"
)

// Integer Integer type constraints
type Integer constraints.Integer

// Limiter database limit clause in MySQL.
//
// eg. SELECT * FROM `user` LIMIT 0,2
//
// eg. SELECT * FROM `user` LIMIT 2 OFFSET 0
type Limiter[INT Integer] struct {
	Offset INT
	Limit  INT
}

// Page business layer pagination struct.
type Page[T any, INT Integer] struct {
	PageNo       INT `json:"page_no"`
	PageSize     INT `json:"page_size"`
	TotalPages   INT `json:"total_pages"`
	TotalRecords INT `json:"total_records"`
	Records      []T `json:"records"`
}

// Pager business layer pagination interface.
type Pager[T any, INT Integer] interface {
	// BuildLimiter get database limit clause for data access layer.
	BuildLimiter() *Limiter[INT]
	// AddRecords append record to collection.
	AddRecords(records ...T)
	// BuildPage get page object for serialization.
	BuildPage() *Page[T, INT]
}

// NewPager return business layer pagination instance.
func NewPager[T any, INT Integer](pageNo, pageSize, totalRecords INT) Pager[T, INT] {
	if pageNo <= 0 {
		pageNo = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	if totalRecords < 0 {
		totalRecords = 0
	}
	capacity := pageSize
	if pageSize > totalRecords {
		capacity = totalRecords
	}

	return &pagerImpl[T, INT]{
		pageNo:       pageNo,
		pageSize:     pageSize,
		totalRecords: totalRecords,
		totalPages:   mustCalculateTotalPages(pageSize, totalRecords),
		records:      make([]T, 0, capacity),
	}
}

func mustCalculateTotalPages[INT Integer](pageSize, totalRecords INT) (totalPages INT) {
	if pageSize <= 0 {
		panic("page size should be positive integer")
	}

	if totalRecords < 0 {
		panic("total records should not be negative integer")
	}

	if totalRecords == 0 {
		return 0
	}

	if totalRecords%pageSize == 0 {
		return totalRecords / pageSize
	}
	return totalRecords/pageSize + 1
}

type pagerImpl[T any, INT Integer] struct {
	pageNo       INT
	pageSize     INT
	totalPages   INT
	totalRecords INT
	records      []T
}

func (p *pagerImpl[T, INT]) BuildLimiter() *Limiter[INT] {
	return &Limiter[INT]{
		Limit:  p.pageSize,
		Offset: (p.pageNo - 1) * p.pageSize,
	}
}

func (p *pagerImpl[T, INT]) AddRecords(records ...T) {
	p.records = append(p.records, records...)
}

func (p *pagerImpl[T, INT]) BuildPage() *Page[T, INT] {
	return &Page[T, INT]{
		PageNo:       p.pageNo,
		PageSize:     p.pageSize,
		TotalPages:   p.totalPages,
		TotalRecords: p.totalRecords,
		Records:      p.records,
	}
}

// EmptyPage return empty record object.
func EmptyPage[T any, INT Integer](pageNo, pageSize INT) *Page[T, INT] {
	return &Page[T, INT]{
		PageNo:       pageNo,
		PageSize:     pageSize,
		TotalPages:   0,
		TotalRecords: 0,
		Records:      make([]T, 0),
	}
}
