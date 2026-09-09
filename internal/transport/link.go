package transport

import (
	"URL_shortner/internal/domain"
	"URL_shortner/internal/service"
	"URL_shortner/pkg/auth"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type LinksTransport struct {
	service *service.LinksService
}

func (t *LinksTransport) Link(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		t.GetLinkByUserID(w, r)
	case http.MethodPost:
		t.CreateLink(w, r)
	case http.MethodDelete:
		t.DeleteLinkByID(w, r)
	case http.MethodPut:
		t.UpdateIsActive(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func NewLinksTransport(service *service.LinksService) *LinksTransport {
	return &LinksTransport{
		service: service,
	}
}

func (t *LinksTransport) GetLinkByUserID(w http.ResponseWriter, r *http.Request) {
	UserID, ok := auth.GetUserId(r)
	if ok != true {
		fmt.Println("ошибка при получении айди пользователя")
		return
	}

	Link, err := t.service.GetLinksByUserID(UserID)
	if err != nil {
		fmt.Println("Ошибка при получении ссылки", err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Link)
}

func (t *LinksTransport) CreateLink(w http.ResponseWriter, r *http.Request) {
	var request domain.CreateLinkRequest
	UserID, ok := auth.GetUserId(r)
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

	if link.IsAcitve == 0 {
		http.Error(w, "Ссылка отключена", http.StatusGone)
		return
	}

	err = t.service.IncrementClicks(link.ID)
	if err != nil {
		http.Error(w, "Ошибка при увеличении числа переходов", http.StatusGone)
		return
	}

	http.Redirect(w, r, link.Original_URL, http.StatusFound)

}

func (t *LinksTransport) DeleteLinkByID(w http.ResponseWriter, r *http.Request) {
	strLinkID := r.URL.Query().Get("LinkID")
	LinkID, err := strconv.Atoi(strLinkID)
	UserID, ok := auth.GetUserId(r)
	if ok != true {
		fmt.Println("Пользователь не зарегистрован")
	}

	if err != nil {
		http.Error(w, "Invalid note id", http.StatusBadRequest)
		return
	}
	t.service.DeleteLinkByID(LinkID, UserID)

}

func (t *LinksTransport) UpdateIsActive(w http.ResponseWriter, r *http.Request) {
	strActive := r.URL.Query().Get("active")
	Active, err := strconv.Atoi(strActive)

	strLinkID := r.URL.Query().Get("LinkID")
	LinkID, err := strconv.Atoi(strLinkID)

	if err != nil {
		http.Error(w, "Invalid note Active", http.StatusBadRequest)
	}

	t.service.UpdateIsActive(LinkID, Active)

}
