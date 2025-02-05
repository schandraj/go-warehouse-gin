package repositories

import (
	"context"
	"gorm.io/gorm"
	"wh-gin/internal/entities"
)

type UserRepository interface {
	Create(ctx context.Context, user entities.User) error
	GetUser(ctx context.Context, username string) (user entities.User, err error)
	GetAllUsers(ctx context.Context, limit int, offset int) ([]entities.User, int, error)
}

type userRepository struct {
	*gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (u *userRepository) Create(ctx context.Context, user entities.User) error {
	err := u.DB.Create(&user).WithContext(ctx).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepository) GetUser(ctx context.Context, username string) (user entities.User, err error) {
	if err = u.DB.Where("username = ?", username).First(&user).WithContext(ctx).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (u *userRepository) GetAllUsers(ctx context.Context, limit int, offset int) ([]entities.User, int, error) {
	var result []entities.User
	var count int64
	if err := u.DB.WithContext(ctx).Order("id asc").Limit(limit).Offset(offset).Find(&result).Error; err != nil {
		return nil, 0, err
	}

	count = u.DB.WithContext(ctx).Find(&[]entities.User{}).RowsAffected

	return result, int(count), nil
}
