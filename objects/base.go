package objects

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Base struct {
	ID        uuid.UUID  `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `sql:"index" json:"-"`
	UpdatedAt time.Time  `json:"update_at"`
}

func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.NewV4())
}
