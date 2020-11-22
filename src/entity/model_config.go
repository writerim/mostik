package entity

type ModelConfig struct {
	Id      int `gorm:"primary_key"`
	IsSet   bool
	Title   string
	Values  string `gorm:"type:json"`
	ModelId int
}
