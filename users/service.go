package users

import (
	"github.com/d-ashesss/knock-knocker/datastore"
	"gorm.io/gorm"
)

type User struct {
	datastore.Model
	Username string
	Password string
}

type Service interface {
	GetUser(username string) (*User, error)
}

func NewService(db *gorm.DB) Service {
	if err := db.Migrator().AutoMigrate(&User{}); err != nil {
		panic(err)
	}
	return &service{db: db}
}

type service struct {
	db *gorm.DB
}

func (s *service) GetUser(username string) (*User, error) {
	var user User
	if err := s.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
