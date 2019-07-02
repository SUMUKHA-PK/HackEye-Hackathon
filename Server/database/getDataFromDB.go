package database

import (
	"fmt"

	"github.com/SUMUKHA-PK/HackEye-Hackathon/Server/util"
)

func GetDataFromDB(UserID string) ([]util.CartData, error) {

	dbConn, err := CreateTables()
	if err != nil {
		return nil, err
	}

	query := "SELECT * FROM cart WHERE userID = '" + UserID + "'"
	fmt.Println(query)

	rows, err := dbConn.Query(query)
	defer rows.Close()

	var (
		userID string
		item   string
		itemID string
	)

	var data []util.CartData

	for rows.Next() {
		if err := rows.Scan(&userID, &item, &itemID); err != nil {
			return nil, err
		}
		data = append(data, util.CartData{UserID: userID, Item: item, ItemID: itemID})
	}

	return data, nil
}
