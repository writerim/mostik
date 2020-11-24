package usecases

import (
	"entity"
	"repositories"
	"time"
)

type userUseCase struct {
	repo           entity.UserRepository
	contextTimeout time.Duration
}

func NewUserUseCase(ur entity.UserRepository, timeout time.Duration) entity.UserUseCases {
	return &userUseCase{
		repo:           ur,
		contextTimeout: timeout,
	}
}

/*
	Проверка что пользователь может видеть пользователя
*/
func (u *entity.UserUseCases) AccessUser(viewer *entity.UserUseCases, r *repositories.UserRepository) error {
	if u.GetPersonalArea().GetId() != viewer.GetPersonalArea().GetId() {
		return errors.New(FORBIDDEN)
	}
	return nil
}

func (u *entity.UserUseCases) GetById(r *repositories.UserRepository, id int) (entity.User, error) {

	// Получим по идентификатору

	user_viewer, err := r.GetById(id)
	if err != nil {
		return user_viewer, err
	}

	// Проверим что у пользователя есть права на эту запись
	if err := u.AccessUser(user_viewer, r); err != nil {
		return user_viewer, err
	}

	return user_viewer, nil
}
