package models

type People struct {
	ID             int    `json:"id" db:"id"`
	Name           string `json:"name" binding:"required" db:"name"`
	Surname        string `json:"surname" binding:"required" db:"surname"`
	Patronymic     string `json:"patronymic" db:"patronymic"`
	Address        string `json:"address" binding:"required" db:"address"`
	PassportNumber int    `json:"passport_number" binding:"required" db:"passport_number"`
	PassportSerie  int    `json:"passport_serie" binding:"required" db:"passport_serie"`
}
