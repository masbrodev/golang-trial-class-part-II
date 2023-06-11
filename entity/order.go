package entity

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	*gorm.Model

	ID           uint   `db:"id"`
	BuyerEmail   string `db:"buyer_email"`
	BuyerAddress string `db:"buyer_address"`
	ProductID    uint   `db:"product_id"`
	OrderTime    time.Time
	Product      Product
}
