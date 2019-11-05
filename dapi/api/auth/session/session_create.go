package session

import (
	"ams/dapi/o/auth/session"
	"ams/dapi/o/org/user"
	"github.com/jinzhu/gorm"
	"http/web"
	"time"
)

func New(u *user.User, db *gorm.DB) (*session.Session, error) {

	var s = &session.Session{
		UserID:    u.ID,
		Email:     u.Email,
		CreatedAt: time.Now(),
	}

	var se, err = s.Create(db)
	if err != nil {
		sessionLog.Error(err)
		return nil, web.InternalServerError("save_session_failed")
	}
	return se, nil
}

func MustNew(u *user.User, db *gorm.DB) *session.Session {
	s, err := New(u, db)
	if err != nil {
		panic(err)
	}
	return s
}
