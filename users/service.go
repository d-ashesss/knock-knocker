package users

import (
	"fmt"
	"github.com/d-ashesss/knock-knocker/datastore"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

const firstUserUsername = "admin@localhost"
const firstUserPassword = "admin"

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
	if err := initFirstUser(db); err != nil {
		panic(err)
	}
	return &service{db: db}
}

func initFirstUser(db *gorm.DB) error {
	var usersCount int64
	if err := db.Model(&User{}).Count(&usersCount).Error; err != nil {
		return fmt.Errorf("initFirstUser: failed to get user count: %v", err)
	}
	if usersCount > 0 {
		return nil
	}
	password, err := bcrypt.GenerateFromPassword([]byte(firstUserPassword), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("initFirstUser: failed to generate password hash: %v", err)
	}
	user := User{Username: firstUserUsername, Password: string(password)}
	if err := db.Create(&user).Error; err != nil {
		return fmt.Errorf("initFirstUser: failed to create user record: %v", err)
	}
	log.Printf("[initFirstUser] created user %s", user.Username)
	return nil
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
