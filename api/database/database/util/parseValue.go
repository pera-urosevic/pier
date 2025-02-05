package util

import (
	"errors"
	"fmt"
)

func ParseValue(val *interface{}, col ColumnInfo) (interface{}, error) {
	var value interface{}

	switch col.Type {

	case "VARCHAR":
		value = fmt.Sprintf("%s", *val)

	case "TEXT":
		value = fmt.Sprintf("%s", *val)

	case "DATE":
		value = *val

	case "DATETIME":
		value = *val

	case "TINYINT":
		value = *val

	case "BIGINT":
		value = *val

	case "INT":
		value = *val

	case "BLOB":
		value = *val

	default:
		return nil, errors.New("parse value unknown type: column [" + col.Name + "], type " + col.Type)

	}

	return value, nil
}
