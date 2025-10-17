package models

type LibraryPaginationResponse struct {
	Page          int64             `json:"page"`
	PageSize      int64             `json:"pageSize"`
	TotalElements int64             `json:"totalElements"`
	Items         []LibraryResponse `json:"items"`
}

type LibraryBookPaginationResponse struct {
	Page          int64                 `json:"page"`
	PageSize      int64                 `json:"pageSize"`
	TotalElements int64                 `json:"totalElements"`
	Items         []LibraryBookResponse `json:"items"`
}
