package database

import (
	"pier/api/database/types"
)

func GetDatabaseIndex() ([]types.DatabaseIndex, error) {
	db := DB()

	rows, err := db.Query("SELECT * FROM `databases` ORDER BY `database`, `collection` ASC")
	if err != nil {
		return nil, err
	}

	var databaseIndex []types.DatabaseIndex
	for rows.Next() {
		var database types.DatabaseIndex
		err := rows.Scan(&database.Database, &database.Collection, &database.Facets)
		if err != nil {
			return nil, err
		}
		databaseIndex = append(databaseIndex, database)
	}

	return databaseIndex, nil
}
