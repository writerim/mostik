package entity

type Rule struct {
	Id     int `gorm:"primary_key"`
	Object string
	Action string
	Title  string
}

type RuleRepository interface {
	Save(Rule) (Rule, error) // Сохранение
	GetAll() ([]Rule, error)
}
