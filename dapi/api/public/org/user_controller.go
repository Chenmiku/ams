package org

import (
	"ams/dapi/o/org/user"
	"http/web"
	"net/http"
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
)

type userAPI struct {
	web.JsonServer
	*http.ServeMux
	db *gorm.DB
}

func newPublicUserAPI() *userAPI {
	u := new(userAPI)
	u.ServeMux = http.NewServeMux()
	u.HandleFunc("/create", u.handleCreate)
	return u
}

func (uapi *userAPI) handleCreate(w http.ResponseWriter, r *http.Request) {
	var u = &user.User{}
	uapi.MustDecodeBody(r, u)
	u.Email = strings.ToLower(u.Email)
	user, err := u.Create(uapi.db)
	if err != nil {
		uapi.ErrorMessage(w, err.Error())
	} else {
		uapi.SendData(w, user)
	}
}

func StrToInt(s string) int {
	i, _ := strconv.ParseInt(s, 10, 64)
	return int(i)
}
