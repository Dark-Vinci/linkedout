package account

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `gorm:"primary key"`
	FirstName   string
	LastName    string
	MiddleName  *string
	DateOfBirth time.Time
	Email       string
	Password    *string
	GoogleToken *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}
