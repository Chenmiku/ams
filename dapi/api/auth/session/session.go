package session

import (
	"ams/dapi/o/org/user"
	"ams/dapi/x/math"
	"encoding/json"
	"time"
)

var idMaker = math.RandStringMaker{Length: 40, Prefix: "s"}

type Session struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Email     string     `json:"email"`
	UserID    string     `json:"userid"`
	Role      user.Role  `json:"role"`
}

func (a *Session) MarshalBinary() ([]byte, error) {
	return json.Marshal(a)
}

func (a *Session) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, a)
}
