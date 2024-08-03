package repository

import (
	"context"

	"yatter-backend-go/app/domain/object"
)

type Status interface {
	Post(ctx context.Context, status string, mediaIds []int)
}
