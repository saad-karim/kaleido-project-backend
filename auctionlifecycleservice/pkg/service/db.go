package service

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // import to support Postgres
	"github.com/saad-karim/kaleido-project-backend/auctionlifecycleservice/pkg/api"
)

type DBConfig struct {
	ConnectionString string
}

type DB struct {
	*sql.DB
}

func DBOpen(connStr string) (*DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

func (db *DB) Ping() error {
	return db.Ping()
}

func (db *DB) InsertAuction(row *api.AuctionDBRow) error {
	query := "INSERT INTO auction (id, item, price, closed) VALUES ($1, $2, $3, $4)"
	_, err := db.DB.Exec(query, row.ID, row.Item, row.Price, row.Closed)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) GetOpenAuctions() ([]api.AuctionDBRow, error) {
	query := "SELECT * from auction WHERE closed = $1"
	rows, err := db.DB.Query(query, false)
	if err != nil {
		return nil, err
	}

	auctionRows := make([]api.AuctionDBRow, 0)
	for rows.Next() {
		auction := api.AuctionDBRow{}
		if err := rows.Scan(&auction); err != nil {
			return nil, err
		}
		auctionRows = append(auctionRows, auction)
	}

	return auctionRows, nil
}

func (db *DB) CloseAuction(auctionID string) error {
	query := fmt.Sprintf("UPDATE auction SET closed = TRUE WHERE id = '%s'", auctionID)
	fmt.Println("Close Auction Query: ", query)
	_, err := db.DB.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
