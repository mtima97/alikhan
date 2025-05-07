package internal

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type Db struct {
	conn *sql.DB
}

func Init() Db {
	connStr := "postgres://user:password@localhost:5432/registration?sslmode=disable"

	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err = conn.Ping(); err != nil {
		log.Fatal(err)
	}

	return Db{conn: conn}
}

func (db *Db) Close() {
	err := db.conn.Close()
	if err != nil {
		log.Fatal(err)
	}
}

//goland:noinspection GoUnhandledErrorResult
func (db *Db) Get() ([]map[int64]string, error) {
	rows, err := db.conn.Query("select * from emails;")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var emails []map[int64]string

	if rows.Next() {
		var id int64
		var email string

		rows.Scan(&id, &email)

		emails = append(emails, map[int64]string{
			id: email,
		})
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return emails, nil
}

func (db *Db) Save(email string) error {
	_, err := db.conn.Exec("insert into emails (email) values ($1);", email)

	return err
}
