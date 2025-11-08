package repository

import (
	"fmt"
	"time"

	model "github.com/Mamvriyskiy/lab2-template/src/bonus/model"
	"github.com/jmoiron/sqlx"
)

type BonusPostgres struct {
	db *sqlx.DB
}

func NewBonusPostgres(db *sqlx.DB) *BonusPostgres {
	return &BonusPostgres{db: db}
}

func (r *BonusPostgres) GetInfoAboutUserPrivilege(username string) (model.PrivilegeResponse, error) {
	var resp model.PrivilegeResponse
	var privilegeID int

	err := r.db.QueryRow(`
        SELECT id, balance, status
        FROM privilege
        WHERE username = $1
    `, username).Scan(&privilegeID, &resp.Balance, &resp.Status)
	if err != nil {
		return resp, err
	}

	rows, err := r.db.Query(`
        SELECT datetime, ticket_uid, balance_diff, operation_type
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
		if err := rows.Scan(&dt, &item.TicketUid, &item.BalanceDiff, &item.OperationType); err != nil {
			return resp, err
		}
		item.Date = dt.Format(time.RFC3339)
	
		resp.History = append(resp.History, item)
	}

	return resp, nil
}

func (r *BonusPostgres) UpdateBonus(username, ticketUid string) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// 1. Получаем ID и текущий баланс пользователя
	var privilegeID, balance int
	query := `SELECT id, balance FROM privilege WHERE username = $1`
	err = tx.QueryRow(query, username).Scan(&privilegeID, &balance)
	if err != nil {
		return fmt.Errorf("user not found: %w", err)
	}

	// 2. Находим, сколько бонусов было списано при покупке
	var spent int
	err = tx.QueryRow(`
		SELECT balance_diff 
		FROM privilege_history
		WHERE privilege_id = $1
		  AND ticket_uid = $2
		  AND operation_type = 'DEBIT_THE_ACCOUNT'
		ORDER BY datetime DESC
		LIMIT 1
	`, privilegeID, ticketUid).Scan(&spent)
	if err != nil {
		return fmt.Errorf("no debit history found for ticket: %w", err)
	}

	if spent <= 0 {
		return fmt.Errorf("nothing to refund for ticket %s", ticketUid)
	}

	// 3. Обновляем баланс пользователя
	newBalance := balance + spent
	_, err = tx.Exec(`UPDATE privilege SET balance = $1 WHERE id = $2`, newBalance, privilegeID)
	if err != nil {
		return err
	}

	// 4. Добавляем запись в историю возврата
	_, err = tx.Exec(`
		INSERT INTO privilege_history (privilege_id, ticket_uid, datetime, balance_diff, operation_type)
		VALUES ($1, $2, NOW(), $3, 'FILL_IN_BALANCE')
	`, privilegeID, ticketUid, spent)
	if err != nil {
		return err
	}

	return tx.Commit()
}
