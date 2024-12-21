package entities

type Mail struct {
	ID        int     `json:"id,omitempty"`
	Theme     string  `json:"theme"`
	Text      string  `json:"text"`
	Image     *string `json:"image,omitempty"`
	SenderID  int     `json:"sender_id"`
	Receivers []int   `json:"receivers"`
}
