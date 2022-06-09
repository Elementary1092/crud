package domain

import "context"

type Storage interface {
	GetOne(ctx context.Context, id int32) (*Post, error)
	GetAll(ctx context.Context) ([]Post, error)
	GetAllByUserId(ctx context.Context, id int32) ([]Post, error)
	Save(ctx context.Context, post *Post) error
	Update(ctx context.Context, post Post) error
	Delete(ctx context.Context, id int32) error
	Close()
}
