package datastore

import (
	"github.com/gofrs/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

// Model defines fields common for most models.
type Model struct {
	UUID      uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func Open(cfg *Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.Dsn()), &gorm.Config{
		Logger: logger.Discard,
	})
	if err != nil {
		return nil, err
	}
	if cfg.Debug {
		db.Logger = logger.Default
		db = db.Debug()
	}
	if len(cfg.TablePrefix) > 0 {
		db.NamingStrategy = schema.NamingStrategy{TablePrefix: cfg.TablePrefix}
	}
	return db, nil
}
