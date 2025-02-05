package util

import (
	"database/sql"
	"pier/api/database/types"
)

func MapDocs(rows *sql.Rows) ([]map[string]interface{}, error) {
	colTypes, err := rows.ColumnTypes()
	if err != nil {
		return nil, err
	}

	colInfos := make([]ColumnInfo, len(colTypes))
	for i, colType := range colTypes {
		colInfos[i] = ColumnInfo{Name: colType.Name(), Type: colType.DatabaseTypeName()}
	}

	cols := make([]interface{}, len(colInfos))
	colPtrs := make([]interface{}, len(colInfos))
	for i := range cols {
		colPtrs[i] = &cols[i]
	}

	docs := []types.Doc{}
	for rows.Next() {
		if err := rows.Scan(colPtrs...); err != nil {
			return nil, err
		}

		doc := make(map[string]interface{})
		for i := range cols {
			val := colPtrs[i].(*interface{})
			value, err := ParseValue(val, colInfos[i])
			if err != nil {
				return nil, err
			}
			doc[colInfos[i].Name] = value
		}

		docs = append(docs, doc)
	}

	return docs, nil
}
