package domain

import (
	"context"
	"github.com/elem1092/crud/internal/adapters/api"
)

type service struct {
	storage Storage
}

func NewService(storage Storage) api.Service {
	return &service{storage}
}

func (s *service) GetPostById(ctx context.Context, id int64) (*Post, error) {
	return s.storage.GetOne(ctx, id)
}

func (s *service) GetPostsByUserId(ctx context.Context, id int64) ([]*Post, error) {
	return s.storage.GetAllByUserId(ctx, id)
}

func (s *service) GetAllPosts(ctx context.Context, limit, offset int) ([]*Post, error) {
	return s.storage.GetAll(ctx, limit, offset)
}

func (s *service) Save(ctx context.Context, dto SavePostDTO) (*Post, error) {
	return s.storage.Save(ctx, dto)
}

func (s *service) Update(ctx context.Context, dto UpdatePostDTO) (*Post, error) {
	return s.storage.Update(ctx, dto)
}

func (s *service) Delete(ctx context.Context, id int64) error {
	return s.storage.Delete(ctx, id)
}
