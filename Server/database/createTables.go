package database

import (
	"database/sql"
	"fmt"

	// To open a PSQL connection
	_ "github.com/lib/pq"
)

// CreateTables returns a DB pointer after creating necessary tables if they dont exist
func CreateTables() (*sql.DB, error) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", "localhost", 54320, "postgres", "", "groceries")
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return db, err
	}

	query := `CREATE TABLE IF NOT EXISTS cart (
				userID varchar(255) NOT NULL, 
				item varchar(255) NOT NULL, 
				itemID varchar(255) NOT NULL,
				PRIMARY KEY (userID,itemID))
				`

	_, err = db.Exec(query)

	return db, err
}
