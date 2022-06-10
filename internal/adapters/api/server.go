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

func (h *server) GetPosts(in *services.GetPostsRequest, stream services.CRUDService_GetPostsServer) error {
	h.logger.Info("Handling get posts request")
	requestType := *(in.GetNeeded().Enum())

	switch requestType {
	case 0:
		if in.GetId() <= 0 {
			h.logger.Errorf("Got illegal id value: %d", in.GetId())
			return ErrIllegalOperation
		}
		h.logger.Infof("Getting post by its id: %d", in.GetId())
		post, err := h.service.GetPostById(context.Background(), in.GetId())
		if err != nil {
			return err
		}

		if err := stream.Send(domainPostToPostDTO(post)); err != nil {
			h.logger.Errorf("Error while sending data to the client: %v", err)
			return err
		}
	case 1:
		if in.GetId() <= 0 {
			h.logger.Errorf("Got illegal id value: %d", in.GetId())
			return ErrIllegalOperation
		}
		h.logger.Infof("Getting post by user's id: %d", in.GetId())
		postsFromDB, err := h.service.GetPostsByUserId(context.Background(), in.GetId())
		if err != nil {
			return err
		}

		for _, post := range convertDomainPostsToPostsDTO(postsFromDB) {
			if err := stream.Send(post); err != nil {
				h.logger.Errorf("Error while sending data to the client: %v", err)
				return err
			}
		}
	case 2:
		h.logger.Info("Getting all posts")
		postsFromDB, err := h.service.GetAllPosts(context.Background())
		if err != nil {
			return err
		}

		for _, post := range convertDomainPostsToPostsDTO(postsFromDB) {
			if err := stream.Send(post); err != nil {
				h.logger.Errorf("Error while sending data to the client: %v", err)
				return err
			}
		}
	default:
		h.logger.Warnf("Got unknown number from enum: %d", requestType)
		return ErrUnknownOperation
	}

	return nil
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
