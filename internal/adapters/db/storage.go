package db

import (
	"context"
	"errors"
	"github.com/Elementary1092/crud/internal/domain"
	"github.com/Elementary1092/crud/pkg/client/postgre"
	"github.com/Elementary1092/crud/pkg/logging"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

type storage struct {
	db     postgre.Client
	logger *logging.Logger
}

func NewPostgreSQLStorage(client postgre.Client, logger *logging.Logger) domain.Storage {
	return &storage{
		db:     client,
		logger: logger,
	}
}

func (s *storage) GetOne(ctx context.Context, id int32) (*domain.Post, error) {
	query := `SELECT id, user_id, title, body FROM posts WHERE id = $1;`
	s.logger.Traceln(query)

	post := domain.Post{}
	err := s.db.QueryRow(ctx, query, id).Scan(&post.Id, &post.UserId, &post.Title, &post.Body)
	if err = s.checkAndLogError(err); err != nil {
		return nil, err
	}

	return &post, err
}

func (s *storage) GetAll(ctx context.Context) ([]*domain.Post, error) {
	query := `SELECT id, user_id, title, body FROM posts;`
	s.logger.Traceln(query)

	rows, err := s.db.Query(ctx, query)
	if err = s.checkAndLogError(err); err != nil {
		return nil, err
	}

	return s.scanRows(rows)
}

func (s *storage) GetAllByUserId(ctx context.Context, id int32) ([]*domain.Post, error) {
	query := `SELECT id, user_id, title, body FROM posts WHERE user_id = $1;`
	s.logger.Traceln(query)

	rows, err := s.db.Query(ctx, query, id)
	if err = s.checkAndLogError(err); err != nil {
		return nil, err
	}
	defer rows.Close()

	return s.scanRows(rows)
}

func (s *storage) Save(ctx context.Context, post *domain.Post) error {
	query := `INSERT INTO posts (user_id, title, body) VALUES ($1, $2, $3) RETURNING id;`
	s.logger.Traceln(query)

	err := s.db.QueryRow(ctx, query, post.UserId, post.Title, post.Body).Scan(&post.Id)
	if err = s.checkAndLogError(err); err != nil {
		return err
	}
	s.logger.Infof("Successfully inserted post: %v", post)

	return nil
}

func (s *storage) Update(ctx context.Context, post domain.Post) error {
	query := `UPDATE posts SET title = $2, body = $3 WHERE id = $1;`
	s.logger.Traceln(query)

	_, err := s.db.Exec(ctx, query, post.Id, post.Title, post.Body)
	if err = s.checkAndLogError(err); err != nil {
		return err
	}

	return nil
}

func (s *storage) Delete(ctx context.Context, id int32) error {
	query := `DELETE FROM posts WHERE id = $1;`
	s.logger.Traceln(query)

	_, err := s.db.Exec(ctx, query, id)
	if err = s.checkAndLogError(err); err != nil {
		return err
	}

	return nil
}

func (s *storage) checkAndLogError(err error) error {
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.Is(err, pgErr) {
			pgErr = err.(*pgconn.PgError)
			s.logger.Errorf(
				"Got SQL error: %s; %s; where: %s",
				pgErr.Message, pgErr.Detail, pgErr.Where,
			)

			return pgErr
		}

		s.logger.Errorf("Caught: %v", err)
		return err
	}

	return nil
}

func (s *storage) scanRows(rows pgx.Rows) ([]*domain.Post, error) {
	posts := make([]*domain.Post, 0)

	for rows.Next() {
		post := &domain.Post{}

		err := rows.Scan(&post.Id, &post.UserId, &post.Title, &post.Body)
		if err != nil {
			s.logger.Errorf("Ocurred while scanning the row : %v", err)
			return nil, err
		}

		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		s.logger.Errorf("Ocurred on rows: %v", err)
		return nil, err
	}

	return posts, nil
}

func (s *storage) Close() {
	s.db.Close()
}
