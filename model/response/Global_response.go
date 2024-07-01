package response

type GlobalJSONResponse struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type GlobalListDataTableResponse struct {
	TotalData int64 `json:"total_data"`
	PerPage   int   `json:"per_page"`
	Page      int   `json:"page"`
	// TotalPage int           `json:"total_page"`
	List []interface{} `json:"list"`
}

type GlobalListResponse struct {
	Code    int                         `json:"code"`
	Status  string                      `json:"status"`
	Message string                      `json:"message"`
	Data    GlobalListDataTableResponse `json:"data,omitempty"`
}

type GlobalSingleResponse struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
