package api

import (
	"context"
	"github.com/elem1092/crud/internal/domain"
)

type service struct {
	storage domain.Storage
}

func NewService(storage domain.Storage) domain.Service {
	return &service{storage}
}

func (s *service) GetPostById(ctx context.Context, id int32) (*domain.Post, error) {
	return s.storage.GetOne(ctx, id)
}

func (s *service) GetPostsByUserId(ctx context.Context, id int32) ([]*domain.Post, error) {
	return s.storage.GetAllByUserId(ctx, id)
}

func (s *service) GetAllPosts(ctx context.Context) ([]*domain.Post, error) {
	return s.storage.GetAll(ctx)
}

func (s *service) Save(ctx context.Context, dto domain.SavePostDTO) (*domain.Post, error) {
	var post = domain.Post{
		UserId: dto.UserId,
		Title:  dto.Title,
		Body:   dto.Body,
	}

	err := s.storage.Save(ctx, &post)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (s *service) Update(ctx context.Context, dto domain.UpdatePostDTO) error {
	post := domain.Post{
		Id:    dto.Id,
		Title: dto.Title,
		Body:  dto.Body,
	}

	return s.storage.Update(ctx, post)
}

func (s *service) Delete(ctx context.Context, id int32) error {
	return s.storage.Delete(ctx, id)
}
