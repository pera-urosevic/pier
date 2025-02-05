package database

import (
	"fmt"
	"pier/api/database/types"
	"pier/lib"
	"strings"
)

func UpdateDoc(database string, collection string, id int64, doc types.Doc) error {
	db := DB()

	table := fmt.Sprintf("`database_%s_%s`", database, collection)
	sets := []string{}
	values := []interface{}{}

	for key := range doc {
		sets = append(sets, "`"+key+"`=?")
		value := doc[key]
		valueString := fmt.Sprintf("%v", value)

		if strings.HasPrefix(valueString, "🕸") && strings.HasSuffix(valueString, "🕸") {
			url := valueString[4 : len(valueString)-4]
			image, err := lib.Download(url)
			if err != nil {
				return err
			}

			value = lib.ResizeImage(image, 256, 256, 70)
		}

		values = append(values, value)
	}

	set := strings.Join(sets, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE `id`=%d", table, set, id)
	_, err := db.Exec(query, values...)
	if err != nil {
		return err
	}

	return nil
}
