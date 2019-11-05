package auth

import (
	"ams/dapi/api/auth/session"
	"ams/dapi/o/org/user"
	"http/web"
	"net/http"
	"strings"

	"github.com/jinzhu/gorm"
)

type AuthServer struct {
	*http.ServeMux
	web.JsonServer
	db *gorm.DB
}

func NewAuthServer() *AuthServer {
	var s = &AuthServer{
		ServeMux: http.NewServeMux(),
	}
	s.HandleFunc("/login", s.HandleLogin)
	s.HandleFunc("/get_profile", s.HandleGetProfile)
	s.HandleFunc("/logout", s.HandleLogout)
	s.HandleFunc("/change_pass", s.handleChangePass)
	return s
}

func (s *AuthServer) MustGetUser(r *http.Request) *user.User {
	var id = session.MustGet(r, s.db).UserID
	var v, err = user.GetByID(s.db, id)
	if err != nil {
		return &user.User{}
	} else {
		return v
	}
}

func (s *AuthServer) HandleGetProfile(w http.ResponseWriter, r *http.Request) {
	u := s.MustGetUser(r)
	s.SendDataSuccess(w, map[string]interface{}{
		"user": &u,
	})
}

func (s *AuthServer) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var body = struct {
		Email    string
		Password string
	}{}

	s.MustDecodeBody(r, &body)

	var u, err = user.GetByemail(s.db, strings.ToLower(body.Email))
	web.AssertNil(err)

	if err = u.ComparePassword(body.Password); err != nil {
		s.SendError(w, err)
		return
	}

	var ses = session.MustNew(u, s.db)
	s.SendData(w, map[string]interface{}{
		"user":  u,
		"token": ses.ID,
	})
}

func (s *AuthServer) HandleLogout(w http.ResponseWriter, r *http.Request) {
	session.MustClear(r, s.db)
	s.SendData(w, nil)
}

func (s *AuthServer) handleChangePass(w http.ResponseWriter, r *http.Request) {
	var body = struct {
		OldPass   string `json:"old_pass"`
		NewPass   string `json:"new_pass"`
		ReNewPass string `json:"re_new_pass"`
		Email     string `json:"email"`
	}{}

	s.MustDecodeBody(r, &body)

	var u, err = user.GetByemail(s.db, strings.ToLower(body.Email))

	if err := u.ComparePassword(body.OldPass); err != nil {
		s.ErrorMessage(w, "password_not_campare")
		return
	}
	err = u.UpdatePass(s.db, body.NewPass)
	if err != nil {
		s.ErrorMessage(w, err.Error())
	} else {
		s.SendData(w, map[string]interface{}{
			"status": "success",
		})
	}
}
