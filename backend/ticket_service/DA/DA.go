package TS_DA

import (
	"fmt"
	"log"
	"os"

	"github.com/lapayka/rsoi-2/Common/Logger"
	TS_structs "github.com/lapayka/rsoi-2/ticket_service/structs"
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

func (db *DB) GetTicketByUUID(uuid, username string) (TS_structs.Ticket, error) {
	ticket := TS_structs.Ticket{TicketUid: uuid, Username: username}

	err := db.db.First(&ticket).Error

	return ticket, err
}

func (db *DB) GetTicketsByUsername(username string) (TS_structs.Tickets, error) {
	tickets := TS_structs.Tickets{}

	err := db.db.Find(&tickets).Where(&TS_structs.Ticket{Username: username}).Error

	return tickets, err
}

func (db *DB) CreateTicket(ticket TS_structs.Ticket) error {
	err := db.db.Create(&ticket).Error

	if err != nil {
		Logger.GetLogger().Print(err)
	}

	return err
}

func (db *DB) DeleteTicket(ticket TS_structs.Ticket) error {
	tx := db.db.Begin()
	err := tx.First(&ticket).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	ticket.Status = "CANCELED"
	err = tx.Save(&ticket).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return err
}
