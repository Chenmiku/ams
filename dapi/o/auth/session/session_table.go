package session

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

func GetAll(db *gorm.DB) (*[]Session, error) {
	sessions := []Session{}
	err := db.Debug().Model(&Session{}).Limit(100).Find(&sessions).Error
	if err != nil {
		return &[]Session{}, err
	}

	return &sessions, nil
}

func GetByID(db *gorm.DB, id string) (*Session, error) {
	var s Session
	err := db.Debug().Model(&Session{}).Where("id = ?", id).Take(&s).Error
	if err != nil {
		return &Session{}, err
	}

	if gorm.IsRecordNotFoundError(err) {
		return &Session{}, errors.New("session_not_found")
	}

	return &s, err
}

func GetByUserID(db *gorm.DB, uid string) (*Session, error) {
	var s Session
	err := db.Debug().Model(&Session{}).Where("userid = ?", uid).Take(&s).Error
	if err != nil {
		return &Session{}, err
	}

	if gorm.IsRecordNotFoundError(err) {
		return &Session{}, errors.New("session_not_found")
	}

	return &s, err
}

func (s *Session) Create(db *gorm.DB) (*Session, error) {
	s.CreatedAt = time.Now()
	s.ID = uuid.New().String()

	err := db.Debug().Create(&s).Error
	if err != nil {
		return &Session{}, err
	}

	return s, nil
}

func MarkDelete(db *gorm.DB, id string) error {
	var s Session
	err := db.Debug().Model(&Session{}).Where("id = ?", id).Take(&s).Delete(&s).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *Session) Update(db *gorm.DB, id string) (*Session, error) {
	db = db.Debug().Model(&Session{}).Where("id = ?", id).Take(&s).UpdateColumns(
		map[string]interface{}{
			"email":      s.Email,
			"userid":     s.UserID,
			"role":       s.Role,
			"updated_at": s.UpdatedAt,
		},
	)
	if db.Error != nil {
		return &Session{}, db.Error
	}

	// This is the display the updated user
	err := db.Debug().Model(&Session{}).Where("id = ?", id).Take(&s).Error
	if err != nil {
		return &Session{}, err
	}
	return s, nil
}
