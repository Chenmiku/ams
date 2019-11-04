package user

import (
	"ams/dapi/x/mlog"
	"time"

	"github.com/jinzhu/gorm"
)

var objectUserLog = mlog.NewTagLog("object_user")

// User struct
type User struct {
	gorm.Model
	Firstname    string    `gorm:"type:varchar(50)" bson:"first_name,omitempty" json:"first_name"`
	Lastname     string    `gorm:"type:varchar(50)" bson:"last_name,omitempty" json:"last_name"`
	Password     string    `gorm:"type:varchar(250)" bson:"password,omitempty" json:"password,omitempty"`
	Email        string    `gorm:"type:varchar(250)" bson:"email,omitempty" json:"email,omitempty"`
	Address      string    `gorm:"type:varchar(500)" bson:"address,omitempty" json:"address,omitempty"`
	PublicAvatar string    `gorm:"type:varchar(500)" bson:"public_avatar,omitempty" json:"public_avatar,omitempty"`
	RoleID       uint      `bson:"role_id,omitempty" json:"role_id"`
	Phone        string    `gorm:"type:varchar(250)" bson:"phone,omitempty" json:"phone"`
	DOB          time.Time `bson:"dob,omitempty" json:"dob"`
	Active       bool      `bson:"active,omitempty" json:"active"`
	Gender       string    `gorm:"type:varchar(50)" bson:"gender,omitempty" json:"gender"`
	Description  string    `gorm:"type:varchar(500)" bson:"description,omitempty" json:"description"`
}

type ChangePassword struct {
	OldPassword string `bson:"old_password" json:"old_password"`
	NewPassword string `bson:"new_password" json:"new_password"`
}

func NewCleanUser() interface{} {
	return &User{}
}

// func (v *User) ensureUniqueEmail() error {
// 	if err := TableUser.NotExist(map[string]interface{}{
// 		"email": v.Email,
// 	}); err != nil {
// 		return err
// 	}
// 	return nil
// }
