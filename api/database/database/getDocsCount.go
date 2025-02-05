package database

import (
	"fmt"
)

func GetDocsCount(database string, collection string, where string) (int, error) {
	db := DB()

	var count int
	table := fmt.Sprintf("`database_%s_%s`", database, collection)
	query := fmt.Sprintf("SELECT COUNT(id) FROM %s WHERE %s", table, where)
	err := db.QueryRow(query).Scan(&count)

	return count, err
}
