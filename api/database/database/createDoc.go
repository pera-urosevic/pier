package database

import (
	"fmt"
	"pier/api/database/types"
	"strings"
)

func CreateDoc(database string, collection string, record types.Doc) (int64, error) {
	db := DB()

	table := fmt.Sprintf("`database_%s_%s`", database, collection)
	docKeys := []string{}
	values := []interface{}{}
	for key := range record {
		docKeys = append(docKeys, "`"+key+"`")
		value := fmt.Sprintf("%v", record[key])
		values = append(values, value)
	}

	columns := strings.Join(docKeys, ", ")
	qms := []string{}
	for range docKeys {
		qms = append(qms, "?")
	}
	placeholders := strings.Join(qms, ", ")
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", table, columns, placeholders)
	res, err := db.Exec(query, values...)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
