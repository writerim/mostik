package entity

type Mark struct {
	Id       int    `gorm:"primary_key"`
	Title    string `gorm:"column:title;type:varchar(100)"`
	IsSystem int
}
