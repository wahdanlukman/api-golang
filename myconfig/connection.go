package myconfig

import (
	"database/sql"
)

//GetMysqlConnect KONEKSI KE DATABSE MYSQL
func GetMysqlConnect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/simasneg-go")
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil

}
