package models

import "time"

type Order struct {
	ID           int       `gorm:"primaryKey" json:"order_id,omitempty"`
	CustomerName string    `gorm:"not null;type:varchar(50)" json:"customer_name" example:"Budi"`
	Items        []Item    `json:"items"`
	OrderedAt    time.Time `json:"ordered_at" example:"2022-10-07T21:39:28.165117+07:00"`
}
