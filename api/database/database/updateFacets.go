package database

func UpdateFacets(database string, collection string, facets string) error {
	db := DB()

	_, err := db.Exec("UPDATE `databases` SET `facets`=? WHERE `database`=? AND `collection`=?", facets, database, collection)
	if err != nil {
		return err
	}

	return nil
}
