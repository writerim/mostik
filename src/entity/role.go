package entity

func (Role) TableName() string {
	return "role"
}

type Role struct {
	Id             int `gorm:"primary_key"`
	Title          string
	PersonalAreaId int
}

type RoleRepository interface {
	GetById(id int) (Role, error) // Получение по id
	Save(Role) (Role, error)      // Сохранение
	GetByIdentByPersonalAreaId(string, int) (Role, error)
	GetDefaultByPersonalAreaId(int) (Role, error)
}
