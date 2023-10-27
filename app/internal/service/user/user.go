package user

import (
	"context"
	"errors"
	"github.com/vildan-valeev/melvad_test/internal/domain"
	"github.com/vildan-valeev/melvad_test/internal/transport/dto"
)

// Repository - методы для работы с БД (интерфейс реализован в инфре)
type Repository interface {
	InsertUser(ctx context.Context, u domain.User) (id int64, err error)
	UpdateUser(ctx context.Context, u domain.User) error
	UpdateUserInCache(ctx context.Context, u domain.User) error
}

// Service - бизнес логика.
type Service struct {
	db Repository
}

func New(db Repository) *Service {
	return &Service{
		db: db,
	}
}

// CreateUser Создание пользователя.
func (s Service) CreateUser(ctx context.Context, u dto.UserCreateDtoRequest) (id int64, err error) {

	if u.Name == "" {
		// TODO: создать кастомные бизнесовые ошибки
		return id, errors.New("Введите Имя")
	}

	user := domain.User{
		Name: u.Name,
		Age:  u.Age,
	}

	id, err = s.db.InsertUser(ctx, user)
	if err != nil {
		return id, err
	}

	return id, nil
}

// UpdateUser Обновление пользователя.
func (s Service) UpdateUserInCache(ctx context.Context, user *dto.UserUpdateDtoRequest) (uint8, error) {
	incrementedAgeUser := domain.User{
		Name: user.Key,
		Age:  user.Value + 1,
	}
	err := s.db.UpdateUserInCache(ctx, incrementedAgeUser)
	if err != nil {
		return user.Value, err
	}
	//return s.db.UpdateUser(ctx, &itemID)
	// query to redis
	return incrementedAgeUser.Age, nil
}
