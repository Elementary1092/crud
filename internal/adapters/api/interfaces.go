package api

import (
	"context"
	"github.com/elem1092/crud/internal/domain"
)

type Service interface {
	GetPostById(ctx context.Context, id int32) (*domain.Post, error)
	GetPostsByUserId(ctx context.Context, id int32) ([]*domain.Post, error)
	GetAllPosts(ctx context.Context) ([]*domain.Post, error)
	Save(ctx context.Context, user domain.SavePostDTO) (*domain.Post, error)
	Update(ctx context.Context, user domain.UpdatePostDTO) error
	Delete(ctx context.Context, id int32) error
}
