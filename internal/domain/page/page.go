package page

type Page struct {
	ID        int    `json:"id,omitempty"`
	Title     string `json:"title"`
	IconSrc   string `json:"iconSrc"`
	IconClass string `json:"iconClass"`
	CoverSrc  string `json:"coverSrc"`
}
