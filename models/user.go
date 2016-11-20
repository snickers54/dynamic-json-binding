package models

type User struct {
	ID    int64
	Name  string `json:"firstname" db:"blabla"`
	Email string
}
