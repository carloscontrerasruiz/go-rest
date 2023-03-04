package dbsqlite

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"github.com/carloscontrerasruiz/gorest/models"
)

type Sqlite struct {
}

var database *sql.DB

func (db *Sqlite) Create(user models.User) (*models.User, error) {
	statement, err := database.Prepare("INSERT INTO users (email, password) VALUES (?,?)")
	if err != nil {
		return nil, err
	}
	result, err := statement.Exec(user.Email, user.Password)
	if err != nil {
		return nil, err
	}

	log.Println(result)
	return &models.User{}, nil
	//rows, _ := database.Query("Select")
}

func (db *Sqlite) GetUserById(id int64) (*models.User, error) {
	rows, err := database.Query("SELECT id, email FROM usuarios WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer func() {
		rows.Close()
	}()

	var user = models.User{}

	for rows.Next() {
		if err = rows.Scan(&user.Id, &user.Email); err == nil {
			return &user, nil
		}
	}

	return nil, err
}

func (db *Sqlite) Close() error {
	return database.Close()
}

func NewSqliteDB(dbName string) *Sqlite {
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal(err)
	}

	statement, err := database.Prepare(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY,
		password VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	);`)
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()

	database = db

	return &Sqlite{}
}
