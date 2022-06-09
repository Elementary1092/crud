package domain

import (
	"strings"
)

type Post struct {
	Id     int64  `json:"id"`
	UserId int64  `json:"user_id omitempty"`
	Title  string `json:"title omitempty"`
	Body   string `json:"body omitempty"`
}

func (p *Post) String() string {
	var str strings.Builder
	str.WriteString("Post{")
	str.WriteString("title: " + shorten(p.Title))
	str.WriteString("; body" + shorten(p.Body))
	return str.String()
}

func shorten(str string) string {
	if len(str) > 20 {
		return str[0:20] + "..."
	}

	return str
}
