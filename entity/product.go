package entity

import "gorm.io/gorm"

type Product struct {
	*gorm.Model
	ID          uint   `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	Price       int    `db:"price"`
}
