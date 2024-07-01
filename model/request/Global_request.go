package request

type DataTableParam struct {
	PerPage  int    `json:"per_page"`
	Page     int    `json:"page"`
	OrderBy  string `json:"order_by"`
	OrderDir string `json:"order_dir"`
}
