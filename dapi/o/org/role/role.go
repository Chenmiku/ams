package role

import (
	"time"
)

type Role struct {
	ID        string `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Name      string     `gorm:"type:varchar(250);not null" bson:"name" json:"name"` //
}
