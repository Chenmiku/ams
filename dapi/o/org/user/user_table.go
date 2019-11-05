package user

import (
	"errors"
	"math/rand"
	"time"

	"github.com/jinzhu/gorm"
)

func (u *User) GetByID(db *gorm.DB, id uint) (*User, error) {
	err := db.Debug().Model(&User{}).Where("id = ?", id).Take(&u).Error
	if err != nil {
		return &User{}, err
	}

	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("user_not_found")
	}

	return u, err
}

func (u *User) GetByemail(db *gorm.DB, email string) (*User, error) {
	err := db.Debug().Model(&User{}).Where("email = ?", email).Take(&u).Error
	if err != nil {
		return &User{}, err
	}

	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("user_not_found")
	}

	return u, err
}

func (u *User) GetAll(db *gorm.DB) (*[]User, error) {
	users := []User{}
	err := db.Debug().Model(&User{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}

// func (u *User) GetAllPaging(pageSize int, pageNumber int, sortBy string, sortOrder string, user *[]User) (int, err) {

// }

func (u *User) Create(db *gorm.DB) (*User, error) {
	var err error
	if err = u.validate(); err != nil {
		return &User{}, errors.New("validate_user_failed")
	}

	if !u.ensureUniqueEmail(db, u.Email) {
		return &User{}, errors.New("email_not_unique")
	}

	//pass := randSeq(6)
	u.Password = "123456"

	var p = password(u.Password)
	// replace
	if err = p.HashTo(&u.Password); err != nil {
		return &User{}, errors.New("hash_password_failed")
	}

	u.CreatedAt = time.Now()

	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}

	return u, nil
}

func (u *User) UpdateById(db *gorm.DB, id uint) (*User, error) {
	if !u.ensureUniqueEmail(db, u.Email) {
		return &User{}, errors.New("email_not_unique")
	}
	db = db.Debug().Model(&User{}).Where("id = ?", id).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"first_name":    u.Firstname,
			"last_name":     u.Lastname,
			"email":         u.Email,
			"password":      u.Password,
			"address":       u.Address,
			"public_avatar": u.PublicAvatar,
			"role_id":       u.RoleID,
			"phone":         u.Phone,
			"dob":           u.DOB,
			"active":        u.Active,
			"gender":        u.Gender,
			"description":   u.Description,
			"update_at":     time.Now(),
		},
	)
	if db.Error != nil {
		return &User{}, db.Error
	}

	// This is the display the updated user
	err := db.Debug().Model(&User{}).Where("id = ?", id).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) MarkDelete(db *gorm.DB, id uint) error {
	err := db.Debug().Model(&User{}).Where("id = ?", id).Take(&u).Delete(&u).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *User) UpdatePass(db *gorm.DB, newValue string) error {
	var update = map[string]interface{}{
		"password": newValue,
	}

	if len(newValue) > 0 {
		var p = password(newValue)
		if err := p.HashTo(&newValue); err != nil {
			return errors.New("generate_hash_password_failed")
		}
		update["password"] = newValue
	}

	db = db.Debug().Model(&User{}).Where("id = ?", u.ID).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"password": newValue,
		},
	)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func (u *User) ensureUniqueEmail(db *gorm.DB, email string) bool {
	err := db.Debug().Model(&User{}).Where("email = ?", email).Take(&u).Error
	if gorm.IsRecordNotFoundError(err) {
		return true
	}
	return false
}

func NewCleanUser() interface{} {
	return &User{}
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
