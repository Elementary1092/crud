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

func (s *service) GetPostsByUserId(ctx context.Context, id int64) ([]Post, error) {
	return s.storage.GetAllByUserId(ctx, id)
}

func (s *service) GetAllPosts(ctx context.Context) ([]Post, error) {
	return s.storage.GetAll(ctx)
}

func (s *service) Save(ctx context.Context, user SavePostDTO) (*Post, error) {
	var post Post
	s.storage.Save(ctx, &post)
	return nil, nil
}

func (s *service) Update(ctx context.Context, user UpdatePostDTO) error {
	var post Post
	return s.storage.Update(ctx, post)
}

func (s *service) Delete(ctx context.Context, id int64) error {
	return s.storage.Delete(ctx, id)
}
