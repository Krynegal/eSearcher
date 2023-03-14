package service

import (
	"crypto/sha1"
	"eSearcher/configs"
	"eSearcher/internal/models"
	"eSearcher/internal/storage"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

type Auth struct {
	store storage.AuthStorage
	cfg   *configs.Config
}

func NewAuth(cfg *configs.Config, storage storage.AuthStorage) *Auth {
	return &Auth{
		cfg:   cfg,
		store: storage,
	}
}

func (a *Auth) CreateUser(login, password string, role int) (int, error) {
	passwordHash := a.generatePasswordHash(password)
	userID, err := a.store.CreateUser(login, passwordHash, role)
	if err != nil {
		return -1, err
	}
	return userID, nil
}

func (a *Auth) AuthUser(login, password string) (*models.User, error) {
	passHash := a.generatePasswordHash(password)
	user, err := a.store.GetUser(login, passHash)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (a *Auth) GenerateToken(uid int, role int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": uid,
		"role_id": role,
	})
	tokenString, err := token.SignedString([]byte("KSFjH$53KSFjH6745u#uEQQjF349%835hFpzA"))
	return tokenString, err
}

func (a *Auth) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(a.cfg.Auth.SecretKey)))
}
