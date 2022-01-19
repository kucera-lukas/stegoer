package controller

import (
	"Stegoer/graph/generated"
	"Stegoer/pkg/entity/model"
	"context"
)

// User controller interface
type User interface {
	Get(ctx context.Context, name string) (*model.User, error)
	Create(ctx context.Context, input generated.NewUser) (*model.User, error)
	Update(ctx context.Context, entUser model.User, input generated.UpdateUser) (*model.User, error)
}
