package repository

import (
	"time"
	"github.com/jmoiron/sqlx"
	model "github.com/Mamvriyskiy/lab2-template/src/bonus/model"
)

type TicketPostgres struct {
	db *sqlx.DB
}

func NewBonusPostgres(db *sqlx.DB) *TicketPostgres {
	return &TicketPostgres{db: db}
}

func (r *TicketPostgres) GetInfoAboutUserPrivilege(username string) (model.PrivilegeResponse, error) {
    var resp model.PrivilegeResponse
    var privilegeID int

	print("+++++++")
    
    err := r.db.QueryRow(`
        SELECT id, balance, status
        FROM privilege
        WHERE username = $1
    `, username).Scan(&privilegeID, &resp.Balance, &resp.Status)
    if err != nil {
        return resp, err
    }

	
    rows, err := r.db.Query(`
        SELECT ticket_uid, datetime, balance_diff, operation_type
        FROM privilege_history
        WHERE privilege_id = $1
        ORDER BY datetime DESC
    `, privilegeID)
    if err != nil {
        return resp, err
    }
    defer rows.Close()

    for rows.Next() {
        var item model.HistoryItem
        var dt time.Time
        if err := rows.Scan(&item.TicketUid, &dt, &item.BalanceDiff, &item.OperationType); err != nil {
            return resp, err
        }
        item.Date = dt.Format(time.RFC3339)
        resp.History = append(resp.History, item)
    }

	print(resp)
    return resp, nil
}

