package session

import (
	"ams/dapi/o/org/user"
	"encoding/json"
	"github.com/jinzhu/gorm"
)

type Session struct {
	gorm.Model
	Email  string    `json:"email"`
	UserID uint    `json:"userid"`
	Role   user.Role `json:"role"`
}

func (s *Session) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}

func (s *Session) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, s)
}
