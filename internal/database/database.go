package database

import (
	"database/sql"
	"log"

	"github.com/Cyndir/tucows-app/internal/model"
	_ "github.com/lib/pq"
)

const connStr = "host=localhost port=15432 user=postgres dbname=tucows password=postgres sslmode=disable"

//go:generate mockgen -destination=../mocks/database.go -package=mocks -source=database.go
type Database interface {
	Connect() error
	Disconnect()
	InsertOrder(order *model.Order) error
	GetOrder(id string) (*model.Order, error)
	UpdateOrder(orderID, status string) error
}

type dbImpl struct {
	db *sql.DB
}

func New() Database {
	return &dbImpl{}
}

func (d *dbImpl) Connect() error {
	database, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	d.db = database
	return nil
}

func (d *dbImpl) Disconnect() {
	if d.db != nil {
		d.db.Close()
	}

}

func (d *dbImpl) InsertOrder(order *model.Order) error {
	_, err := d.db.Exec("INSERT into orders VALUES ($1, $2, $3, $4)", order.ID, order.CustomerID, order.ProductID, order.Status)

	return err
}

func (d *dbImpl) GetOrder(id string) (*model.Order, error) {
	var res model.Order

	row := d.db.QueryRow("SELECT id, customerid, productid, status FROM orders WHERE id = $1", id)
	err := row.Scan(&res.ID, &res.CustomerID, &res.ProductID, &res.Status) // only going to have one
	if err != nil {
		log.Printf("Database error: %s", err)
		return nil, err
	}

	return &res, nil
}

func (d *dbImpl) UpdateOrder(orderID, status string) error {
	_, err := d.db.Exec("UPDATE orders SET status = $1 WHERE id = $2", status, orderID)

	return err
}
