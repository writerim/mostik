package entity

type Model struct {
	Id            int    `gorm:"primary_key"`
	Library       string `gorm:"type:varchar(255)"`
	Title         string `gorm:"type:varchar(255)"`
	MarkId        int    `gorm:"column:mark_id"`
	IsSystem      int
	Functionality string `gorm:"column:functionality;type:json"`
}
