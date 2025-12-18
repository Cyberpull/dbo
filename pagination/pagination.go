package pagination

type Getter[T any] interface {
	GetCurrentPage() uint64
	GetFrom() uint64
	GetTo() uint64
	GetPerPage() int
	GetLastPage() uint64
	GetTotal() uint64
	GetData() []T
}

type Setter[T any] interface {
	SetCurrentPage(page uint64)
	SetFrom(page uint64)
	SetTo(page uint64)
	SetPerPage(page int)
	SetLastPage(page uint64)
	SetTotal(page uint64)
	SetData(data []T)
}

type Pagination[T any] struct {
	CurrentPage uint64 `json:"current_page"`
	From        uint64 `json:"from"`
	LastPage    uint64 `json:"last_page"`
	PerPage     int    `json:"per_page"`
	To          uint64 `json:"to"`
	Total       uint64 `json:"total"`
	Data        []T    `json:"data"`
}

func (x *Pagination[T]) GetCurrentPage() uint64 {
	return x.CurrentPage
}

func (x *Pagination[T]) GetFrom() uint64 {
	return x.From
}

func (x *Pagination[T]) GetTo() uint64 {
	return x.To
}

func (x *Pagination[T]) GetPerPage() int {
	return x.PerPage
}

func (x *Pagination[T]) GetLastPage() uint64 {
	return x.LastPage
}

func (x *Pagination[T]) GetTotal() uint64 {
	return x.Total
}

func (x *Pagination[T]) GetData() []T {
	return x.Data
}

func (x *Pagination[T]) SetCurrentPage(page uint64) {
	x.CurrentPage = page
}

func (x *Pagination[T]) SetFrom(page uint64) {
	x.From = page
}

func (x *Pagination[T]) SetTo(page uint64) {
	x.To = page
}

func (x *Pagination[T]) SetPerPage(page int) {
	x.PerPage = page
}

func (x *Pagination[T]) SetTotal(page uint64) {
	x.Total = page
}

func (x *Pagination[T]) SetLastPage(page uint64) {
	x.LastPage = page
}

func (x *Pagination[T]) SetData(data []T) {
	x.Data = data
}
