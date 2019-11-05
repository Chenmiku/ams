package session

import (
	"ams/dapi/o/org/user"
	"ams/dapi/x/math"
	"encoding/json"

	"github.com/jinzhu/gorm"
)

var idMaker = math.RandStringMaker{Length: 40, Prefix: "s"}

type Session struct {
	gorm.Model
	Email  string    `json:"email"`
	UserID string    `json:"userid"`
	Role   user.Role `json:"role"`
}

func (a *Session) MarshalBinary() ([]byte, error) {
	return json.Marshal(a)
}

func (a *Session) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, a)
}
