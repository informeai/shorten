package store

import (
	"database/sql"
	"github.com/informeai/shorten/entities"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

//Shortener is interface for methods the store package.
type Shortener interface {
	Get(id string) (entities.Shorten, error)
	Insert(srt entities.Shorten) error
}

//Store is struct for db.
type Store struct {
	db *sql.DB
}

//NewStore return instance the store.
func NewStore() *Store {
	db, err := sql.Open("sqlite3", "./shortlinks.db")
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS short (id VARCHAR(64), url VARCHAR(255), visits INTEGER)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}
	return &Store{db: db}
}

//Get return first Shorten from database.
func (s *Store) Get(id string) (entities.Shorten, error) {
	srt := entities.Shorten{}
	rows, err := s.db.Query("SELECT * FROM short")
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
	storeId = ""

	return srt, nil
}

//Insert add new Shorten to database.
func (s *Store) Insert(srt entities.Shorten) error {
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
