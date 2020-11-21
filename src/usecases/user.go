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

func (u *entity.UserUseCases) MoveTo(r *repositories.UserRepository, id, to_id int) (entity.User, error) {

	// move to move

	return u.repo.GetById(id), nil
}
