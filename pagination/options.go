package pagination

type Options struct {
	Page  int
	Limit int
}

func (x *Options) parse() (page, offset, from, limit int) {
	page, limit = max(x.Page, 1), max(x.Limit, 10)

	if page > 1 {
		offset = page - 1
		offset *= limit
	}

	from = 1

	if offset > 0 {
		from += offset
	}

	return
}
