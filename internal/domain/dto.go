package domain

type SavePostDTO struct {
	UserId int32  `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type UpdatePostDTO struct {
	Id    int32  `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}
