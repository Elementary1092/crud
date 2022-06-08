package domain

type Post struct {
	Id     int64  `json:"id"`
	UserId int64  `json:"user_id omitempty"`
	Title  string `json:"title omitempty"`
	Body   string `json:"body omitempty"`
}
