package entity

type Parameter struct {
	Id       int `gorm:"primary_key"`
	Ident    string
	Title    string
	GroupCat string `gorm:"type:text"`
	Divider  int
	Round    int
	Utils    string
	IsSystem int
}
