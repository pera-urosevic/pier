package database

func RemoveCache(key string) error {
	db := DB()

	_, err := db.Exec("DELETE FROM `cache` WHERE `key`=?", key)
	if err != nil {
		return err
	}

	return nil
}
