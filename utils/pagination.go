package utils

func DecodePage(page, size int) (int, int, int) {
	if page <= 0 {
		page = 1
	}

	if size <= 0 {
		size = 10
	}

	offset := page*size - size

	return page, size, offset
}
