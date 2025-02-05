package database

import (
	"fmt"
	"pier/api/database/database/util"
	"pier/api/database/types"
)

func GetDoc(database string, collection string, id int64) (types.Doc, error) {
	db := DB()

	table := fmt.Sprintf("`database_%s_%s`", database, collection)
	query := fmt.Sprintf("SELECT * FROM %s WHERE `id`=?", table)
	rows, err := db.Query(query, id)
	if err != nil {
		return nil, err
	}

	docs, err := util.MapDocs(rows)
	if err != nil {
		return nil, err
	}

	return docs[0], nil
}
