package org

import (
	"ams/dapi/api/auth/session"
	"ams/dapi/o/org/user"
	"github.com/jinzhu/gorm"
	"http/web"
	"net/http"
	"strconv"
	"strings"
)

type UserServer struct {
	web.JsonServer
	*http.ServeMux
	db *gorm.DB
}

func NewUserServer() *UserServer {
	var s = &UserServer{
		ServeMux: http.NewServeMux(),
	}

	s.HandleFunc("/get_all", s.HandleAllUser)
	s.HandleFunc("/create", s.HandleCreate)
	s.HandleFunc("/get", s.HandleGetByID)
	s.HandleFunc("/update", s.HandleUpdateByID)
	s.HandleFunc("/mark_delete", s.HandleMarkDelete)
	s.HandleFunc("/change_password", s.ChangePassword)
	return s
}

func StrToInt(s string) int {
	i, _ := strconv.ParseInt(s, 10, 64)
	return int(i)
}

func (s *UserServer) Session(w http.ResponseWriter, r *http.Request) *session.Session {

	ses, err := session.FromContext(r.Context())

	if err != nil {
		s.SendError(w, err)
	}

	return ses
}

func (s *UserServer) HandleAllUser(w http.ResponseWriter, r *http.Request) {

	// sortBy := r.URL.Query().Get("sort_by")
	// sortOrder := r.URL.Query().Get("sort_order")

	// pageSize := StrToInt(r.URL.Query().Get("page_size"))
	// pageNumber := StrToInt(r.URL.Query().Get("page_number"))

	//var res = []user.User{}
	res, err := user.GetAll(s.db)

	if err != nil {
		s.SendErrorMessage(w, err)
	} else {
		s.SendDataSuccess(w, map[string]interface{}{
			"users": res,
			//"count": count,
		})
	}
}

func (s *UserServer) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var u = &user.User{}
	s.MustDecodeBody(r, u)
	u.Email = strings.ToLower(u.Email)
	user, err := u.Create(s.db)
	if err != nil {
		s.ErrorMessage(w, err.Error())
	} else {
		s.SendDataSuccess(w, user)
	}
}

func (s *UserServer) mustGetUser(r *http.Request) *user.User {
	var id = r.URL.Query().Get("id")
	var u, err = user.GetByID(s.db, id)
	if err != nil {
		return &user.User{}
	} else {
		return u
	}
}

func (s *UserServer) HandleUpdateByID(w http.ResponseWriter, r *http.Request) {
	var newUser = &user.User{}
	s.MustDecodeBody(r, newUser)
	newUser.Email = strings.ToLower(newUser.Email)
	u := s.mustGetUser(r)
	user, err := u.UpdateById(s.db, u.ID)
	if err != nil {
		s.ErrorMessage(w, err.Error())
	} else {
		s.SendDataSuccess(w, user)
	}
}

func (s *UserServer) HandleGetByID(w http.ResponseWriter, r *http.Request) {
	u := s.mustGetUser(r)
	s.SendDataSuccess(w, u)
}

func (s *UserServer) HandleMarkDelete(w http.ResponseWriter, r *http.Request) {
	u := s.mustGetUser(r)
	err := user.MarkDelete(s.db, u.ID)
	if err != nil {
		s.ErrorMessage(w, err.Error())
	} else {
		s.Success(w)
	}
}

func (s *UserServer) ChangePassword(w http.ResponseWriter, r *http.Request) {

	var change = &user.ChangePassword{}

	s.MustDecodeBody(r, change)

	u, err := user.GetByID(s.db, s.Session(w, r).UserID)
	if err != nil {
		s.ErrorMessage(w, "user_not_found")
		return
	}

	if err = u.ComparePassword(change.OldPassword); err != nil {
		s.ErrorMessage(w, "password_not_campare")
		return
	}

	err = u.UpdatePass(s.db, change.NewPassword)
	if err != nil {
		s.ErrorMessage(w, err.Error())
	} else {
		s.Success(w)
	}
}
