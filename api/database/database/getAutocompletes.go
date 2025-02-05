package database

import (
	"fmt"
	"pier/api/database/types"
)

func GetAutocompletes(database string, collection string, fields types.AutocompleteFields) (types.Autocompletes, error) {
	autocompletes := types.Autocompletes{}

	db := DB()

	table := fmt.Sprintf("database_%s_%s", database, collection)
	for _, field := range fields.Selects {
		query := fmt.Sprintf("SELECT DISTINCT(`%s`) FROM `%s` ORDER BY `%s` ASC", field, table, field)
		rows, err := db.Query(query)
		if err != nil {
			return nil, err
		}

		var values []string
		for rows.Next() {
			var value string
			err := rows.Scan(&value)
			if err != nil {
				return nil, err
			}

			if value != "" {
				values = append(values, value)
			}
		}

		autocompletes[field] = values
	}

	for _, field := range fields.Tags {
		query := fmt.Sprintf("SELECT DISTINCT tags.tag from `%s` cross join json_table(`%s`.`%s`, '$[*]' COLUMNS (tag varchar(255) path '$')) tags order by tags.tag ASC", table, table, field)
		rows, err := db.Query(query)
		if err != nil {
			return nil, err
		}

		var values []string
		for rows.Next() {
			var value string
			err := rows.Scan(&value)
			if err != nil {
				return nil, err
			}

			if value != "" {
				values = append(values, value)
			}
		}

		autocompletes[field] = values
	}

	return autocompletes, nil
}
