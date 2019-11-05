package user

import (
	"errors"
	"fmt"
	"html"
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

func (u *User) Prepare() {
	u.ID = ""
	u.Firstname = html.EscapeString(strings.TrimSpace(u.Firstname))
	u.Lastname = html.EscapeString(strings.TrimSpace(u.Lastname))
	u.Gender = ""
	u.Phone = ""
	u.PublicAvatar = ""
	u.RoleID = ""
	u.UpdatedAt = time.Now()
	u.CreatedAt = time.Now()
	u.DOB = time.Now()
	u.Description = ""
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.Password = html.EscapeString(strings.TrimSpace(u.Password))
	u.Active = true
	u.Address = ""
}

func GetByID(db *gorm.DB, id string) (*User, error) {
	var u User
	err := db.Debug().Model(&User{}).Where("id = ?", id).Take(&u).Error
	if err != nil {
		return &User{}, err
	}

	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("user_not_found")
	}

	return &u, err
}

func GetByemail(db *gorm.DB, email string) (*User, error) {
	var u User
	err := db.Debug().Model(&User{}).Where("email = ?", email).Take(&u).Error
	if err != nil {
		return &User{}, err
	}

	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("user_not_found")
	}

	return &u, err
}

func GetAll(db *gorm.DB) (*[]User, error) {
	users := []User{}
	err := db.Debug().Model(&User{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}

// func GetAllPaging(pageSize int, pageNumber int, sortBy string, sortOrder string, user *[]User) (int, err) {

// }

func (u *User) Create(db *gorm.DB) (*User, error) {
	fmt.Println("create")
	var err error
	if err = u.validate(); err != nil {
		return &User{}, errors.New("validate_user_failed")
	}

	fmt.Println(u.Email)
	// if !ensureUniqueEmail(db, u.Email) {
	// 	return &User{}, errors.New("email_not_unique")
	// }

	//pass := randSeq(6)
	u.Password = "123456"

	var p = password(u.Password)
	// replace
	if err = p.HashTo(&u.Password); err != nil {
		return &User{}, errors.New("hash_password_failed")
	}

	u.CreatedAt = time.Now()
	u.ID = uuid.New().String()
	fmt.Println(u.ID)

	err = db.Debug().Create(&u).Error
	fmt.Println("created")
	if err != nil {
		return &User{}, err
	} else {
		return u, nil
	}
}

func (u *User) UpdateById(db *gorm.DB, id string) (*User, error) {
	if !ensureUniqueEmail(db, u.Email) {
		return &User{}, errors.New("email_not_unique")
	}
	db = db.Debug().Model(&User{}).Where("id = ?", id).Take(&u).UpdateColumns(
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

func MarkDelete(db *gorm.DB, id string) error {
	var u User
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

func ensureUniqueEmail(db *gorm.DB, email string) bool {
	fmt.Println("ensure")
	var u User
	err := db.Debug().Model(&User{}).Where("email = ?", email).Take(&u).Error
	fmt.Println(err.Error())
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
