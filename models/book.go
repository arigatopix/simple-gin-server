package models

type Book struct {
	ID     uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Title  string `gorm:"type:varchar(100)" json:"title" binding:"min=2,max=100"`
	Author string `json:"author"`
}
