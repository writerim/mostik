package entity

type ModelField struct {
	Id           int `gorm:"primary_key"`
	Name         string
	DefaultValue string
	Pattern      string
	MaxLength    int
	TypeValue    string
	ModelId      int
	IsSystem     int
}
