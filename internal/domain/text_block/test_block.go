package textblock

type TextBlock struct {
	ID     int    `json:"id,omitempty"`
	Text   string `json:"text"`
	PageId int    `json:"page_id"`
	Order  int    `json:"order"`
	Type   string `json:"type"`
}
