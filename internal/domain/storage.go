package domain

import "context"

type Storage interface {
	GetOne(ctx context.Context, id int64) (*Post, error)
	GetAll(ctx context.Context, limit, offset int) ([]*Post, error)
	GetAllByUserId(ctx context.Context, id int64) ([]*Post, error)
	Save(ctx context.Context, user SavePostDTO) (*Post, error)
	Update(ctx context.Context, user UpdatePostDTO) (*Post, error)
	Delete(ctx context.Context, id int64) error
}
