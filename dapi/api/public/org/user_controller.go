package org

import (
	"ams/dapi/o/org/user"
	//"encoding/json"
	"http/web"
	//"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"fmt"
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
	u.Prepare()
	u.Email = strings.ToLower(u.Email)
	fmt.Println(*u)
	user, err := u.Create(uapi.db)
	fmt.Println(&user)
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
