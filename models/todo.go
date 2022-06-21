package models

type Todo struct {
	ID       uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Text     string `gorm:"type:varchar(250)" json:"text" binding:"min=2,max=250"`
	Day      string `gorm:"type:varchar(100)" json:"day" binding:"min=2,max=100"`
	Reminder bool   `gorm:"default:false" json:"reminder"`
}
