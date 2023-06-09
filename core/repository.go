package core

import "context"

type Repository interface {
	Store(ctx context.Context, e *Event) error
	FindByIdPrefix(ctx context.Context, prefixes []string) ([]Event, error)
	FindByAuthors(ctx context.Context, authors []string) ([]Event, error)
	Find(ctx context.Context, id string) (*Event, error)
}
