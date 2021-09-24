package store

import (
	"context"
	"database/sql"
	"github.com/informeai/shorten/entities"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

//StoreSqlite is struct for db sqlite.
type StoreSqlite struct {
	db *sql.DB
}

//NewStoreSqlite return instance the StoreSqlite.
func NewStoreSqlite() *StoreSqlite {
	db, err := sql.Open("sqlite3", "./shortlinks.db")
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS short (id VARCHAR(255) PRIMARY KEY, url VARCHAR(255), visits INTEGER)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}
	return &StoreSqlite{db: db}
}

//Get return first Shorten from database.
func (s *StoreSqlite) Get(id string) (entities.Shorten, error) {
	srt := entities.Shorten{}
	rows, err := s.db.QueryContext(context.Background(), "SELECT * FROM short")
	defer rows.Close()
	if err != nil {
		return srt, err
	}
	var storeId string
	var storeUrl string
	var storeVisits int

	for rows.Next() {
		err = rows.Scan(&storeId, &storeUrl, &storeVisits)

		if err != nil {
			return srt, err
		}
		if id == storeId {
			srt.Id = storeId
			srt.Url = storeUrl
			srt.Visits = storeVisits
			return srt, nil

		}
	}

	return srt, nil
}

//Insert add new Shorten to database.
func (s *StoreSqlite) Insert(srt entities.Shorten) error {
	stmt, err := s.db.Prepare("INSERT INTO short(id,url,visits) values(?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(srt.Id, srt.Url, srt.Visits)
	if err != nil {
		return err
	}
	return nil

}

//Update change the shorten and save in database.
func (s *StoreSqlite) Update(srt entities.Shorten) error {
	stmt, err := s.db.Prepare("UPDATE short SET visits=? WHERE id=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(srt.Visits+1, srt.Id)
	if err != nil {
		return err
	}
	return nil
}
