package domain

type Links struct {
	Id           int    `json:"id"`
	Original_URL string `json:"original_url"`
	Short_URL    string `json:"short_url"`
	Clicks       string `json:"clicks"`
	Created_at   string `json:"created_at"`
	Expires_at   string `json:"expires_at"`
	Is_active    int    `json:"is_active"`
}
