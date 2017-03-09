package page

import "errors"

func calcTotalPages(pageSize, totalRecords int64) (totalPages int64, err error) {
	if pageSize == 0 {
		return 0, errors.New("divide by zero")
	}
	if totalRecords%pageSize == 0 {
		totalPages = totalRecords / pageSize
	} else {
		totalPages = totalRecords/pageSize + 1
	}
	return totalPages, nil
}

func calcOffset(page, pageSize, totalRecords int64) (offset int64, err error) {
	totalPages, err := calcTotalPages(pageSize, totalRecords)
	if err != nil {
		return 0, err
	}

	if totalPages > 0 && page > totalPages {
		page = totalPages
	}
	return (page - 1) * pageSize, nil
}
