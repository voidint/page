package page

import "errors"

var (
	// ErrDivideByZero divide by zero error
	ErrDivideByZero = errors.New("divide by zero")
	// ErrNonPositiveInteger invalid parameter error
	ErrNonPositiveInteger = errors.New("invalid non-positive integer parameter")
)

func calcTotalPages(pageSize, totalRecords int64) (totalPages int64, err error) {
	if pageSize <= 0 {
		return 0, errors.New("page size should be positive integer")
	}

	if totalRecords < 0 {
		return 0, errors.New("total records should not be negative integer")
	}

	if totalRecords == 0 {
		return 0, nil
	}

	if totalRecords%pageSize == 0 {
		totalPages = totalRecords / pageSize
	} else {
		totalPages = totalRecords/pageSize + 1
	}
	return totalPages, nil
}

func calcOffset(page, pageSize, totalRecords int64) (offset int64, err error) {
	if page <= 0 {
		return 0, errors.New("page number should be positive integer")
	}

	if pageSize <= 0 {
		return 0, errors.New("page size should be positive integer")
	}

	if totalRecords < 0 {
		return 0, errors.New("total records should not be negative integer")
	}

	totalPages, err := calcTotalPages(pageSize, totalRecords)
	if err != nil {
		return 0, err
	}

	if totalPages > 0 && page > totalPages {
		page = totalPages
	}
	return (page - 1) * pageSize, nil
}
