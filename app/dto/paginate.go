package dto

//分页相关

// Paginate 分页器
type Paginate struct {
	Limit   int `json:"limit"`
	Count   int `json:"count"`
	Total   int `json:"total"`
	Last    int `json:"last"`
	Current int `json:"current"`
	Next    int `json:"next"`
}
