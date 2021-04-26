package postgresql

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Store struct {
	db *sql.DB
}

//Создание структуры хранилища
func New() *Store {
	return &Store{}
}

//Открытие соединения с базой данных
func (s *Store) Open() error {
	sourceName := "user=serebryakov password=123 dbname=course_go_binance sslmode=disable"
	db, err := sql.Open("postgres", sourceName)
	if err != nil {
		panic(err)
	}
	s.db = db
	return nil
}

//INSERT запрос к базе данных
func (s *Store) InsertBinanceData(symbol string, price string) {
	_, err := s.db.Exec("insert into binance (Symbol, Price) values ($1, $2)", symbol, price)
	if err != nil {
		panic(err)
	}
}

//Закрытие соединения с базой данных
func (s *Store) Close() {
	s.db.Close()
}
