package session

import (
	"ams/dapi/o/org/user"
	"encoding/json"
	"time"
)

type Session struct {
	ID        string `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Email     string     `gorm:"type:varchar(250)" json:"email"`
	UserID    string     `json:"userid"`
	Role      user.Role  `json:"role"`
}

func (s *Session) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}

func (s *Session) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, s)
}
