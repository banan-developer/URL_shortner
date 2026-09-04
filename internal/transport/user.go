package transport

import (
	"URL_shortner/internal/domain"
	"URL_shortner/internal/service"
	"URL_shortner/pkg/auth"
	"log"
	"net/http"
)

type UserTransport struct {
	service *service.UserService
}

func NewUserTransport(service *service.UserService) *UserTransport {
	return &UserTransport{
		service: service,
	}
}

func (t *UserTransport) RegistratingUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.ServeFile(w, r, "./web/html/registration.html")
		return
	}
	if r.Method == http.MethodPost {
		User := &domain.User{
			UserName: r.FormValue("Username"),
			Email:    r.FormValue("email"),
			Password: r.FormValue("password"),
		}
		err := t.service.RegistratingUser(User)
		if err != nil {
			log.Printf("Ошибка регистрации: %v", err)
			return
		}

	}
	http.Redirect(w, r, "/login", http.StatusSeeOther)

}

func (t *UserTransport) LoginUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.ServeFile(w, r, "./web/html/login.html")
		return
	}
	if r.Method == http.MethodPost {
		login := r.FormValue("email")
		password := r.FormValue("password")
		UserID, err := t.service.LoginUser(login, password)

		if err != nil {
			http.Error(w, "invalid credentials", http.StatusUnauthorized)
			return
		}
		auth.SetUserID(w, r, UserID)
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}
}
