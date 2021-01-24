package service

import (
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq" // import to support Postgres
	"github.com/saad-karim/kaleido-project-backend/biddingservice/pkg/api"
)

type DB struct {
	*sqlx.DB
}

func DBOpen(connStr string) (*DB, error) {
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

func (db *DB) Ping() error {
	return db.Ping()
}

func (db *DB) GetOpenAuctions() ([]api.AuctionDBRow, error) {
	query := "SELECT * from auction WHERE closed = $1"

	auctionRows := []api.AuctionDBRow{}
	err := db.DB.Select(&auctionRows, query, false)
	if err != nil {
		return nil, err
	}

	// auctionRows := make([]api.AuctionDBRow, 0)
	// for rows.Next() {
	// 	auction := api.AuctionDBRow{}
	// 	if err := rows.Scan(&auction); err != nil {
	// 		return nil, err
	// 	}
	// 	auctionRows = append(auctionRows, auction)
	// }

	return auctionRows, nil
}
