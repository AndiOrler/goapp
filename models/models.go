package models

type User struct {
	ID   int    `Json: ID, gorm: "primaryKey; autoIncrement"`
	Name string `Json: Name`
}
