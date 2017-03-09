package page

import (
	"fmt"
	"reflect"
)

type Limiter struct {
	Offset int64
	Limit  int64
}

// Page pagination struct
type Page struct {
	Page         int64         `json:"Page"`
	PageSize     int64         `json:"PageSize"`
	TotalPages   int64         `json:"TotalPages"`
	TotalRecords int64         `json:"TotalRecords"`
	Records      []interface{} `json:"Records"`
}

// Pager pagination interface
type Pager interface {
	BuildLimiter() *Limiter
	AddRecords(records ...interface{}) error
	BuildPage() *Page
}

type pagerImpl struct {
	page         int64
	pageSize     int64
	totalPages   int64
	totalRecords int64
	records      []interface{}
	elemType     reflect.Type
}

func NewPager(elemType reflect.Type, page, pageSize, totalRecords int64) Pager {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	totalPages, _ := calcTotalPages(pageSize, totalRecords)
	return &pagerImpl{
		page:         page,
		pageSize:     pageSize,
		totalPages:   totalPages,
		totalRecords: totalRecords,
		records:      make([]interface{}, 0, pageSize),
		elemType:     elemType,
	}
}

func (p *pagerImpl) BuildLimiter() *Limiter {
	offset, _ := calcOffset(p.page, p.pageSize, p.totalRecords)
	return &Limiter{
		Limit:  p.pageSize,
		Offset: offset,
	}
}

func (p *pagerImpl) isAcceptableElem(k interface{}) bool {
	return reflect.TypeOf(k) == p.elemType
}

func (p *pagerImpl) AddRecords(records ...interface{}) error {
	for _, record := range records {
		if !p.isAcceptableElem(record) {
			return fmt.Errorf("invalid element: %#v", record)
		}
	}
	p.records = append(p.records, records...)
	return nil
}

func (p *pagerImpl) BuildPage() *Page {
	return &Page{
		Page:         p.page,
		PageSize:     p.pageSize,
		TotalPages:   p.totalPages,
		TotalRecords: p.totalRecords,
		Records:      p.records,
	}
}

func EmptyPage(page, pageSize int64) *Page {
	return &Page{
		Page:         page,
		PageSize:     pageSize,
		TotalPages:   0,
		TotalRecords: 0,
		Records:      make([]interface{}, 0),
	}
}
