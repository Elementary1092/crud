package domain

type SavePostDTO struct {
	Id     int64  `json:"id"`
	UserId int64  `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type UpdatePostDTO struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}
