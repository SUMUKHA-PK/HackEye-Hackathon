package database

import (
	"fmt"

	"github.com/SUMUKHA-PK/HackEye-Hackathon/Server/util"
)

// AddGroceryListToDatabase adds the grocery list to the DB
func AddGroceryListToDatabase(data util.AddToCartReq) error {

	dbConn, err := CreateTables()
	if err != nil {
		return err
	}
	fmt.Println(data)

	for i := range data.ItemList {
		query := "INSERT INTO cart VALUES ('" +
			data.UserID + "','" +
			data.ItemList[i] + "','" +
			data.ItemIDList[i] + "'" +
			")"
		fmt.Println(query)
		_, err = dbConn.Exec(query)
		if err != nil {
			return err
		}
	}

	return nil
}
