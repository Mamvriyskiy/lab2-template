package models

type LibraryPaginationResponse struct {
	Page          int64             `json:"page"`
	PageSize      int64             `json:"page_size"`
	TotalElements int64             `json:"total_elements"`
	Items         []LibraryResponse `json:"items"`
}

type LibraryBookPaginationResponse struct {
	Page          int64                 `json:"page"`
	PageSize      int64                 `json:"page_size"`
	TotalElements int64                 `json:"total_elements"`
	Items         []LibraryBookResponse `json:"items"`
}
