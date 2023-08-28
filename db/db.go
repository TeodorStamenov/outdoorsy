package db

import (
	"database/sql"
	"fmt"

	"github.com/TeodorStamenov/outdoorsy/models"
	"github.com/TeodorStamenov/outdoorsy/util"
	_ "github.com/lib/pq"
)

type Db interface {
	GetVehicle(int) ([]models.Vehicle, error)
	GetUser(int) (models.User, error)
}

type Postgres struct {
	db *sql.DB
}

func NewDb(c util.Config) (Db, error) {
	conn := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		c.Db.Addr, c.Db.Port, c.Db.User, c.Db.Pass, c.Db.Name)

	fmt.Printf("Database trying to connect with: %s\n", conn)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	return &Postgres{
		db: db,
	}, err
}

func (p *Postgres) GetVehicle(id int) ([]models.Vehicle, error) {
	sqlStatement := `SELECT user_id, name, type FROM rentals WHERE user_id=$1;`

	rows, err := p.db.Query(sqlStatement, id)
	if err != nil {
		fmt.Println(err)
	}

	vehicles := make([]models.Vehicle, 0)
	for rows.Next() {
		v := models.Vehicle{}
		err = rows.Scan(&v.Id, &v.Name, &v.Type)
		if err != nil {
			fmt.Println(err)
		}
		vehicles = append(vehicles, v)
	}

	return vehicles, nil
}

func (p *Postgres) GetUser(id int) (models.User, error) {
	sqlStatement := `SELECT * FROM users WHERE id=$1;`

	row := p.db.QueryRow(sqlStatement, id)

	user := models.User{}

	err := row.Scan(&user.Id, &user.FirstName, &user.LastName)
	if err != nil {
		fmt.Println(err)
	}

	return user, nil
}
