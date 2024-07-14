package repository

import (
	"fmt"

	"github.com/AyBalatsan/TimeTracker/models"
	"github.com/go-sqlx/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) CreateUser(user models.People) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, surname, patronymic, address, passport_number, passport_serie) values ($1, $2, $3, $4, $5, $6) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Surname, user.Patronymic, user.Address, user.PassportNumber, user.PassportSerie)
	// тут мы записываем переменную id
	if err := row.Scan(&id); err != nil {
		return 0, nil
	}
	return id, nil //
}

func (r *UserPostgres) GetAllUser() ([]models.People, error) {
	var users []models.People
	query := fmt.Sprintf("SELECT tl.id, tl.name, tl.surname, tl.patronymic, tl.address  FROM %s", usersTable)
	err := r.db.Select(&users, query)
	return users, err
}

// ID             int    `json:"id" db:"id"`
// 	Name           string `json:"name" binding:"required" db:"name"`
// 	Surname        string `json:"surname" binding:"required" db:"surname"`
// 	Patronymic     string `json:"patronymic" db:"patronymic"`
// 	Address        string `json:"address" binding:"required" db:"address"`
// 	PassportNumber int    `json:"passport_number" binding:"required" db:"passport_number"`
// 	PassportSerie  int    `json:"passport_serie" binding:"required" db:"passport_serie"`
