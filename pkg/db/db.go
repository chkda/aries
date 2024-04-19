package db

import "context"

type Writer interface {
	Write(ctx context.Context, query string, args ...any) error
}
