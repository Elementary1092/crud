package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/elem1092/crud/internal/domain"
	services "github.com/elem1092/crud/pkg/client/grpc"
	"github.com/elem1092/crud/pkg/logging"
)

var (
	ErrEmptyContent     = errors.New("did not expect empty content of a post")
	ErrUnknownOperation = errors.New("unknown operation")
	ErrIllegalOperation = errors.New("illegal operation")
)

type server struct {
	service domain.Service
	logger  *logging.Logger
	services.UnimplementedCRUDServiceServer
}

func NewServer(service domain.Service, logger *logging.Logger) services.CRUDServiceServer {
	return &server{service, logger, services.UnimplementedCRUDServiceServer{}}
}

func (h *server) SavePost(ctx context.Context,
	in *services.SavePostDTO) (*services.PostDTO, error) {
	h.logger.Info("Handling save request")
	if in.GetContent() == nil {
		h.logger.Error("Got empty content of a post")
		return nil, ErrEmptyContent
	}

	newPost := domain.SavePostDTO{
		UserId: in.GetUserId(),
		Title:  in.GetContent().GetTitle(),
		Body:   in.GetContent().GetBody(),
	}

	saved, err := h.service.Save(ctx, newPost)
	if err != nil {
		return nil, err
	}

	h.logger.Info("Successfully saved a record")

	return domainPostToPostDTO(saved), nil
}

func (h *server) GetPosts(ctx context.Context, in *services.GetPostsRequest) (*services.GetPostsResponse, error) {
	h.logger.Info("Handling get posts request")
	requestType := *(in.GetNeeded().Enum())

	posts := make([]*services.PostDTO, 1)
	switch requestType {
	case 0:
		if in.GetId() <= 0 {

		}
		h.logger.Info("Getting post by its id: %d", in.GetId())
		post, err := h.service.GetPostById(ctx, in.GetId())
		if err != nil {
			return nil, err
		}

		posts[0] = domainPostToPostDTO(post)
	case 1:
		postsFromDB, err := h.service.GetPostsByUserId(ctx, in.GetId())
		if err != nil {
			return nil, err
		}

		posts = convertDomainPostsToPostsDTO(postsFromDB)
	case 2:
		postsFromDB, err := h.service.GetAllPosts(ctx)
		if err != nil {
			return nil, err
		}

		posts = convertDomainPostsToPostsDTO(postsFromDB)
	default:
		h.logger.Warnf("Got unknown number from enum: %d", requestType)
		return nil, ErrUnknownOperation
	}

	return &services.GetPostsResponse{Posts: posts}, nil
}

func (h *server) DeletePost(ctx context.Context,
	in *services.DeleteRequest) (*services.ErrorResponse, error) {
	h.logger.Info("Handling delete request")
	id := in.GetId()
	if id <= 0 {
		h.logger.Warnf("Got invalid id to delete: %d", id)
		return &services.ErrorResponse{Error: fmt.Sprintf("invalid id: %d", id)}, ErrIllegalOperation
	}

	err := h.service.Delete(ctx, id)
	if err != nil {
		return &services.ErrorResponse{Error: err.Error()}, err
	}

	return nil, nil
}

func (h *server) UpdatePost(ctx context.Context,
	in *services.UpdatePostDTO) (*services.ErrorResponse, error) {
	h.logger.Info("Handling delete request")
	id := in.GetId()
	if id <= 0 {
		h.logger.Warnf("Got invalid id to delete: %d", id)
		return &services.ErrorResponse{Error: fmt.Sprintf("invalid id: %d", id)}, ErrIllegalOperation
	}

	content := in.GetContent()
	if content == nil {
		h.logger.Warn("Got empty content")
		return &services.ErrorResponse{Error: ErrEmptyContent.Error()}, ErrEmptyContent
	}

	postToUpdate := domain.UpdatePostDTO{
		Id:    id,
		Title: content.GetTitle(),
		Body:  content.GetBody(),
	}

	err := h.service.Update(ctx, postToUpdate)
	if err != nil {
		return &services.ErrorResponse{Error: err.Error()}, err
	}

	return nil, nil
}

func domainPostToPostDTO(post *domain.Post) *services.PostDTO {
	return &services.PostDTO{
		Id:     post.Id,
		UserId: post.UserId,
		Content: &services.ContentDTO{
			Title: post.Title,
			Body:  post.Body,
		},
	}
}

func convertDomainPostsToPostsDTO(posts []*domain.Post) []*services.PostDTO {
	postsDTO := make([]*services.PostDTO, len(posts))
	for i := range posts {
		postsDTO[i] = domainPostToPostDTO(posts[i])
	}

	return postsDTO
}
