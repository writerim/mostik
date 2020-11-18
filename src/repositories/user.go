package repositories

import(
	"entity"
	"database/sql"
)

type mysqlUserRepo struct {
	DB *sql.DB
}

type UserRepository interface {
	GetByID(id int64) (entity.User, error)
}

func (m *mysqlUserRepo) GetByID(id int64) (entity.User, error){
	// select parse return
	return entity.User{} , nil
}

func NewMysqlUserRepository(db *sql.DB) entity.UserRepository {
	return &mysqlUserRepo{
		DB: db,
	}
}