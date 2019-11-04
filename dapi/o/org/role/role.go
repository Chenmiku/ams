package role

import (
	"github.com/jinzhu/gorm"
)

type Role struct {
	gorm.Model
	Name       string   `bson:"name" json:"name"` //
	Permission []string `bson:"permission" json:"permission"`
}
