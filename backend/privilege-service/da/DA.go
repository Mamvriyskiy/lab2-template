package PS_DA

import (
	"fmt"
	"log"
	"os"

	PS_structs "github.com/lapayka/rsoi-2/privilege-service/structs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	db *gorm.DB
}

func GetConnectionString() (string) {
	// Берём обязательные части из переменных окружения
	host, ok := os.LookupEnv("DB_HOST")
	if !ok {
		return ""
	}

	port, ok := os.LookupEnv("DB_PORT")
	if !ok {
		return ""
	}

	user, ok := os.LookupEnv("DB_USER")
	if !ok {
		return ""
	}

	password, ok := os.LookupEnv("DB_PASSWORD")
	if !ok {
		return ""
	}

	dbname, ok := os.LookupEnv("DB_NAME")
	if !ok {
		return ""
	}

	// Формируем строку подключения
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		user,
		password,
		host,
		port,
		dbname,
	)

	return connStr
}

func New(host, user, db_name, password string) (*DB, error) {
	dsn := GetConnectionString()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("unable to connect database", err)
	}

	return &DB{db: db}, nil
}

func (db *DB) GetPrivilegeAndHistoryByUserName(username string) (PS_structs.Privilege_with_history, error) {
	p := PS_structs.Privilege{Username: username}

	tx := db.db.Begin()

	err := tx.First(&p).Error

	if err != nil {
		tx.Rollback()
		return PS_structs.Privilege_with_history{}, nil
	}

	transactions := PS_structs.Privileges_history{}

	err = db.db.Find(&transactions).Where("Privilege_id = ", p.ID).Error

	if err != nil {
		tx.Rollback()
		return PS_structs.Privilege_with_history{Privilege_info: p}, nil
	}

	tx.Commit()

	return PS_structs.Privilege_with_history{Privilege_info: p, History: transactions}, err
}

func (db *DB) CreateTicket(username string, price int64, is_paid_from_balance bool, privelege_item PS_structs.Privilege_history) error {
	privelege := PS_structs.Privilege{Username: username}

	tx := db.db.Begin()
	err := tx.First(&privelege).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	privelege_item.BalanceDiff = 0
	privelege_item.PrivilegeID = privelege.ID
	if is_paid_from_balance {
		diff := price
		if price > privelege.Balance {
			diff = privelege.Balance
		}
		privelege.Balance -= diff
		privelege_item.BalanceDiff = diff

		err = tx.Save(&privelege).Error

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Create(&privelege_item).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}
