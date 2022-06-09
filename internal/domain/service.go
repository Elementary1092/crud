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

func (s *service) GetPostById(ctx context.Context, id int32) (*Post, error) {
    return s.storage.GetOne(ctx, id)
}

func (s *service) GetPostsByUserId(ctx context.Context, id int32) ([]*Post, error) {
    return s.storage.GetAllByUserId(ctx, id)
}

func (s *service) GetAllPosts(ctx context.Context) ([]*Post, error) {
    return s.storage.GetAll(ctx)
}

func (s *service) Save(ctx context.Context, dto SavePostDTO) (*Post, error) {
    var post = Post{
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

func (s *service) Update(ctx context.Context, dto UpdatePostDTO) error {
    post := Post{
        Id:    dto.Id,
        Title: dto.Title,
        Body:  dto.Body,
    }

    return s.storage.Update(ctx, post)
}

func (s *service) Delete(ctx context.Context, id int32) error {
    return s.storage.Delete(ctx, id)
}
