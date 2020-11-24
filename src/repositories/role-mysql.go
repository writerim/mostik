package repositories

import (
	"entity"
	"github.com/jinzhu/gorm"
)

type mysqlRoleRepo struct {
	DB *gorm.DB
}

const (
	DEFAULT_IDENT = "guest"
)

type RoleRepository interface {
	GetById(id int) (entity.Role, error)
	Save(role entity.Role) (entity.Role, error)
	GetByIdentByPersonalAreaId(string, int) (entity.Role, error)
	GetDefaultByPersonalAreaId(int) (entity.Role, error)
}

func NewMysqlRoleRepository(db *gorm.DB) entity.RoleRepository {
	return &mysqlRoleRepo{
		DB: db,
	}
}

func (m *mysqlRoleRepo) GetById(id int) (entity.Role, error) {
	// select parse return
	return entity.Role{}, nil
}

func (m *mysqlRoleRepo) Save(role entity.Role) (entity.Role, error) {
	m.DB.Create(&role)
	return role, nil
}

func (m *mysqlRoleRepo) GetDefaultByPersonalAreaId(personal_area_id int) (entity.Role, error) {

	if r, err := m.GetByIdentByPersonalAreaId(DEFAULT_IDENT, personal_area_id); err == nil && r.Id == 0 {
		r := entity.Role{}
		r.Title = DEFAULT_IDENT
		r.PersonalAreaId = personal_area_id
		return m.Save(r)
	}

	return m.GetByIdentByPersonalAreaId(DEFAULT_IDENT, personal_area_id)
}

func (m *mysqlRoleRepo) GetByIdent(title string) (entity.Role, error) {
	record := entity.Role{}
	m.DB.Where("title = ?", title).First(&record)
	return record, nil
}

func (m *mysqlRoleRepo) GetByIdentByPersonalAreaId(title string, personal_area_id int) (entity.Role, error) {
	record := entity.Role{}
	m.DB.
		Where("title = ?", title).
		Where("personal_area_id = ?", personal_area_id).
		First(&record)
	return record, nil
}
