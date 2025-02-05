package database

import (
	"fmt"
)

func GetDuplicate(database string, collection string, id int64, name string) (bool, error) {
	db := DB()

	table := fmt.Sprintf("`database_%s_%s`", database, collection)
	query := fmt.Sprintf("SELECT COUNT(id) FROM %s WHERE `id`!=? AND `name`=?", table)
	row := db.QueryRow(query, id, name)
	var count int
	err := row.Scan(&count)
	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}
