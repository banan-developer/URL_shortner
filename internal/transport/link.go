package transport

import (
	"URL_shortner/internal/domain"
	"URL_shortner/internal/service"
	"URL_shortner/pkg/auth"
	"encoding/json"
	"fmt"
	"net/http"
)

type LinksTransport struct {
	service *service.LinksService
}

func (t *LinksTransport) Link(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		t.GetLinkByID(w, r)
	case http.MethodPost:
		t.CreateLink(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func NewLinksTransport(service *service.LinksService) *LinksTransport {
	return &LinksTransport{
		service: service,
	}
}

func (t *LinksTransport) GetLinkByID(w http.ResponseWriter, r *http.Request) {
	UserID := 1

	Link, err := t.service.GetLinkByID(UserID)
	if err != nil {
		fmt.Println("Ошибка при получении ссылки")
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Link)
}

func (t *LinksTransport) CreateLink(w http.ResponseWriter, r *http.Request) {
	var request domain.CreateLinkRequest
	UserID, ok := auth.GetUserId(r)
	fmt.Println("USER ID:", UserID)
	fmt.Println("USER ID ERROR:", ok)

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	result, err := t.service.CreateLink(request.Original_URL, UserID)
	if err != nil {
		http.Error(w, "error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (t *LinksTransport) GetLinkByShortlink(w http.ResponseWriter, r *http.Request) {
	shortCode := r.PathValue("shortCode")

	link, err := t.service.GetLinkByShortlink(shortCode)
	if err != nil {
		fmt.Println(err)
		return
	}

	http.Redirect(w, r, link.Original_URL, http.StatusFound)

}
