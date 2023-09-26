package domain

import (
	"errors"
	"regexp"
	"time"

	"gorm.io/gorm"
)

type Student struct {
	ID       uint      `gorm:"-"`
	IDLain   uint      `gorm:"primaryKey;column:id_lain"`
	Email    string    `gorm:"not null;unique;type:varchar(100)"`
	Products []Product `gorm:"foreignKey:UserID;association_foreignkey:IDLain"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (s *Student) BeforeCreate(tx *gorm.DB) error {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailRegex.MatchString(s.Email) {
		return errors.New("email is not valid")
	}
	return nil
}

type Product struct {
	ID     uint   `gorm:"primaryKey"`
	Brand  string `gorm:"not null;type:varchar(100)"`
	Name   string `gorm:"not null;type:varchar(100)"`
	UserID uint

	CreatedAt time.Time
	UpdatedAt time.Time
}
