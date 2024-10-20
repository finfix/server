package model

type Icon struct {
	ID   int    `json:"id" db:"id"`     // ID иконки
	Name string `json:"name" db:"name"` // Название иконки
	Url  string `json:"url" db:"img"`   // URL иконки
}
