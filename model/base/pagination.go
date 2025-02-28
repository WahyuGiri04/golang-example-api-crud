package base

type Pagination struct{
	Page int `json:"page"`
	PageSize int `json:"page_size"`
	TotalRows int64 `json:"total_rows"`
	TotalPages int `json:"total_pages"`
	Data  interface{} `json:"data"`
}