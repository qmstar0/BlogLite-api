package paginate

// Paginate 分页器
type Paginate struct {
	PageSize    int
	TotalItems  int
	TotalPages  int
	CurrentPage int
	PrevPage    int
	NextPage    int
}

func NewPaginate(total, limit, current int) Paginate {
	if total < 0 {
		total = 0
	}
	if limit <= 0 {
		limit = 1
	}
	if current <= 0 {
		current = 1
	}
	totalPages := (total + limit - 1) / limit
	prevPage := current - 1
	if prevPage < 1 {
		prevPage = 0
	}
	nextPage := current + 1
	if nextPage > totalPages {
		nextPage = 0
	}
	return Paginate{
		PageSize:    limit,
		TotalItems:  total,
		TotalPages:  totalPages,
		CurrentPage: current,
		PrevPage:    prevPage,
		NextPage:    nextPage,
	}
}
