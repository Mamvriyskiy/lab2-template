package model

type HistoryItem struct {
    Date          string `json:"date"`
    TicketUid     string `json:"ticketUid"`
    BalanceDiff   int    `json:"balanceDiff"`
    OperationType string `json:"operationType"`
}

type PrivilegeResponse struct {
    Balance int           `json:"balance"`
    Status  string        `json:"status"`
    History []HistoryItem `json:"history"`
}
