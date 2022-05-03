package apiserver

import (
	"Diplom/internal/app/store/sqlstore"
	"database/sql"
	"fmt"
	"net/http"
)

func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}
	defer db.Close()
	store := sqlstore.New(db)
	srv := newServer(store)

	return http.ListenAndServe(config.BindAddr, srv.router)
}

func newDB(databaseURL string) (*sql.DB, error) {
	fmt.Println("In new db", databaseURL)
	db, err := sql.Open("mysql", databaseURL)
	if err != nil {
		fmt.Println("Error open " + string(err.Error()))
		return nil, err
	}

	if err := db.Ping(); err != nil {
		fmt.Println("In ping", databaseURL)
		return nil, err
	}

	return db, nil
}
