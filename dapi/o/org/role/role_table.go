package role

import (
	"errors"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"time"
)

func GetByID(db *gorm.DB, id string) (*Role, error) {
	var r Role
	err := db.Debug().Model(&Role{}).Where("id = ?", id).Take(&r).Error
	if err != nil {
		return &Role{}, err
	}

	if gorm.IsRecordNotFoundError(err) {
		return &Role{}, errors.New("role_not_found")
	}

	return &r, err
}

func GetAll(db *gorm.DB) (*[]Role, error) {
	roles := []Role{}
	err := db.Debug().Model(&Role{}).Limit(100).Find(&roles).Error
	if err != nil {
		return &[]Role{}, err
	}
	return &roles, err
}

// func (u *User) GetAllPaging(pageSize int, pageNumber int, sortBy string, sortOrder string, user *[]User) (int, err) {

// }

func (r *Role) Create(db *gorm.DB) (*Role, error) {
	var err error

	r.CreatedAt = time.Now()
	r.ID = uuid.New().String()

	err = db.Debug().Create(&r).Error
	if err != nil {
		return &Role{}, err
	}

	return r, nil
}

func (r *Role) UpdateById(db *gorm.DB, id string) (*Role, error) {
	db = db.Debug().Model(&Role{}).Where("id = ?", id).Take(&r).UpdateColumns(
		map[string]interface{}{
			"name":      r.Name,
			"update_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &Role{}, db.Error
	}

	// This is the display the updated Role
	err := db.Debug().Model(&Role{}).Where("id = ?", id).Take(&r).Error
	if err != nil {
		return &Role{}, err
	}
	return r, nil
}

func MarkDelete(db *gorm.DB, id string) error {
	var r Role
	err := db.Debug().Model(&Role{}).Where("id = ?", id).Take(&r).Delete(&r).Error
	if err != nil {
		return err
	}
	return nil
}
