package commment

import (
	"context"
	"fmt"
)

// Comment - a rep of the comment struct for our service.
type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

// Store - Hold all of our method need to interact with our data repository.
type Store interface {
	GetComment(ctx context.Context, id string) (Comment, error)
}

type Service struct {
	Store Store
}

// NewService - Returns a pointer to a new service
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

// GetComment - Get a comment ny id stored within the database.
func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("Getting latest comments")
	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Comment{}, err
	}
	return cmt, nil
}
