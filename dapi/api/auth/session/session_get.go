package session

import (
	"ams/dapi/o/auth/session"
	"http/web"

	"github.com/jinzhu/gorm"
)

const (
	errReadSessonFailed   = web.InternalServerError("read_session_failed")
	errSessionNotFound    = web.Unauthorized("session_not_found")
	errUnauthorizedAccess = web.Unauthorized("unauthorized_access")
)

func Get(sessionID string, db *gorm.DB) (*session.Session, error) {
	s, err := session.GetByID(db, sessionID)
	if err != nil {
		sessionLog.Error(err)
		return nil, errReadSessonFailed
	}

	return s, nil
}
