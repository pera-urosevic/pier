package database

import (
	"fmt"
	"pier/api/database/database/util"
	"pier/api/database/types"
)

func GetDocs(database string, collection string, where string, order string, offset int) ([]types.Doc, error) {
	db := DB()

	table := fmt.Sprintf("`database_%s_%s`", database, collection)
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s ORDER BY %s LIMIT %d OFFSET %d", table, where, order, 100, offset)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	docs, err := util.MapDocs(rows)
	if err != nil {
		return nil, err
	}

	return docs, nil
}
