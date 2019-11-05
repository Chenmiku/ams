package org

import (
	"ams/dapi/o/org/role"
	"http/web"
	"net/http"

	"github.com/jinzhu/gorm"
)

type RoleServer struct {
	web.JsonServer
	*http.ServeMux
	db *gorm.DB
}

func NewRoleServer() *RoleServer {
	var s = &RoleServer{
		ServeMux: http.NewServeMux(),
	}
	s.HandleFunc("/create", s.HandleCreate)
	s.HandleFunc("/get", s.HandleGetByID)
	s.HandleFunc("/update", s.HandleUpdateByID)
	s.HandleFunc("/mark_delete", s.HandleMarkDelete)
	s.HandleFunc("/get_all", s.HandleAllRole)
	return s
}

func (s *RoleServer) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var u = &role.Role{}
	s.MustDecodeBody(r, u)
	role, err := u.Create(s.db)
	if err != nil {
		s.ErrorMessage(w, err.Error())
	} else {
		s.SendDataSuccess(w, role)
	}
}

func (s *RoleServer) mustGetRole(r *http.Request) *role.Role {
	var id = r.URL.Query().Get("id")
	var u, err = role.GetByID(s.db, id)
	web.AssertNil(err)
	return u
}

func (s *RoleServer) HandleUpdateByID(w http.ResponseWriter, r *http.Request) {
	var newRole = &role.Role{}
	s.MustDecodeBody(r, newRole)
	var u = s.mustGetRole(r)
	role, err := u.UpdateById(s.db, newRole.ID)
	if err != nil {
		s.ErrorMessage(w, err.Error())
	} else {
		s.SendDataSuccess(w, role)
	}
}

func (s *RoleServer) HandleGetByID(w http.ResponseWriter, r *http.Request) {
	var u = s.mustGetRole(r)
	s.SendDataSuccess(w, u)
}

func (s *RoleServer) HandleMarkDelete(w http.ResponseWriter, r *http.Request) {
	var u = s.mustGetRole(r)
	err := role.MarkDelete(s.db, u.ID)
	if err != nil {
		s.ErrorMessage(w, err.Error())
	} else {
		s.Success(w)
	}
}

func (s *RoleServer) HandleAllRole(w http.ResponseWriter, r *http.Request) {

	// sortBy := r.URL.Query().Get("sort_by")
	// sortOrder := r.URL.Query().Get("sort_order")

	// pageSize := StrToInt(r.URL.Query().Get("page_size"))
	// pageNumber := StrToInt(r.URL.Query().Get("page_number"))

	res, err := role.GetAll(s.db)

	if err != nil {
		s.ErrorMessage(w, err.Error())
	} else {
		s.SendDataSuccess(w, map[string]interface{}{
			"roles": res,
			//"count": count,
		})
	}
}
