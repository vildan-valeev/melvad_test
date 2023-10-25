package user

import (
	"context"
	"errors"
	"github.com/vildan-valeev/melvad_test/internal/domain"
	"github.com/vildan-valeev/melvad_test/internal/transport/dto"
	"math/rand"
)

// Repository - методы для работы с БД (интерфейс реализован в инфре)
type Repository interface {
	InsertUser(ctx context.Context, u domain.User) error
	UpdateUser(ctx context.Context, id int64) error
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
func (s Service) CreateUser(ctx context.Context, u dto.UserCreateDtoRequest) (int64, error) {
	id := rand.Int63()

	if u.Name == "" {
		// TODO: создать кастомные ошибки
		return id, errors.New("Введите Имя")
	}

	user := domain.User{
		ID:   rand.Int63(),
		Name: u.Name,
		Age:  u.Age,
	}

	if err := s.db.InsertUser(ctx, user); err != nil {
		return id, err
	}

	return id, nil
}

// UpdateUser Обновление пользователя.
func (s Service) UpdateUser(ctx context.Context, user dto.UserUpdateDto) (uint8, error) {
	//return s.db.UpdateUser(ctx, &itemID)
	// query to redis
	return user.Value + 1, nil
}