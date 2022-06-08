package db

import (
	"context"
	"github.com/elem1092/crud/internal/domain"
)

type Storage struct {
}

func (s *Storage) GetOne(ctx context.Context, id int64) (*domain.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Storage) GetAll(ctx context.Context, limit, offset int) ([]*domain.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Storage) GetAllByUserId(ctx context.Context, id int64) ([]*domain.Post, error) {
	return nil, nil
}

func (s *Storage) Save(ctx context.Context, user domain.SavePostDTO) (*domain.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Storage) Update(ctx context.Context, user domain.UpdatePostDTO) (*domain.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Storage) Delete(ctx context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}
