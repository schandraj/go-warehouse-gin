package usecases

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
	"wh-gin/internal/entities"
	"wh-gin/internal/repositories"
	"wh-gin/utils"
)

type UserUsecase interface {
	Register(ctx context.Context, input entities.User) error
	Login(ctx context.Context, username, password string) (string, error)
	Me(ctx context.Context, username string) (entities.User, error)
	GetAll(ctx context.Context, limit int, offset int) ([]entities.User, int, error)
}

type userUsecase struct {
	repo repositories.UserRepository
}

func NewUserUsecase(repository repositories.UserRepository) UserUsecase {
	return &userUsecase{
		repo: repository,
	}
}

func (u *userUsecase) Register(ctx context.Context, input entities.User) error {
	return u.repo.Create(ctx, input)
}

func (u *userUsecase) Login(ctx context.Context, username, password string) (string, error) {
	user, err := u.repo.GetUser(ctx, username)
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	// Verify password
	if err = utils.VerifyPassword(user.Password, password); err != nil {
		return "", errors.New("invalid username or password")
	}

	// Create JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func (u *userUsecase) Me(ctx context.Context, username string) (entities.User, error) {
	return u.repo.GetUser(ctx, username)
}

func (u *userUsecase) GetAll(ctx context.Context, limit int, offset int) ([]entities.User, int, error) {
	return u.repo.GetAllUsers(ctx, limit, offset)
}
