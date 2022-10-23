package models

type Item struct {
	ID          int    `gorm:"primaryKey" json:"line_item_id,omitempty"`
	ItemCode    string `gorm:"not null;type:varchar(20)" json:"item_code" example:"A12B3C"`
	Description string `gorm:"type:varchar(255)" json:"description" example:"This is the description of the item"`
	Quantity    int    `gorm:"not null" json:"quantity" example:"10"`
	OrderID     int    `json:"-"`
}
