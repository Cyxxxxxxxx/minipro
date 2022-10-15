package mysql

//insert authID into minipro_auth
func InsertAuth() {

}

//select auth id from minipro_person
func SelectAuth(userid string) error {
	sqlAuth := "SELECT * FROM minipro_auth WHERE userid = ? "
	err := db.Get(sqlAuth, userid)
	if err != nil {
		return err
	}
	return err
}

func DeleteAuth(userid string) error {
	sqlDelAuth := "DELETE FROM minipro_auth WHERE userid = ? "
	err := db.Get(sqlDelAuth, userid)
	return err
}
