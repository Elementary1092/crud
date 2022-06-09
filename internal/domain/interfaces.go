package domain

import (
	"context"
)

type Service interface {
	GetPostById(ctx context.Context, id int32) (*Post, error)
	GetPostsByUserId(ctx context.Context, id int32) ([]*Post, error)
	GetAllPosts(ctx context.Context) ([]*Post, error)
	Save(ctx context.Context, user SavePostDTO) (*Post, error)
	Update(ctx context.Context, user UpdatePostDTO) error
	Delete(ctx context.Context, id int32) error
}
