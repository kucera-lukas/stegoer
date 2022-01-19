package repository

import (
	"Stegoer/ent"
	"Stegoer/ent/user"
	"Stegoer/graph/generated"
	"Stegoer/pkg/adapter/controller"
	"Stegoer/pkg/entity/model"
	"Stegoer/pkg/util"
	"context"
)

// NewUserRepository returns a specific implementation of the controller.User interface
func NewUserRepository(client *ent.Client) controller.User {
	return &userRepository{client: client}
}

type userRepository struct {
	client *ent.Client
}

func (r *userRepository) Get(ctx context.Context, name string) (*model.User, error) {
	entUser, err := r.client.User.Query().Where(user.NameEQ(name)).Only(ctx)

	if err != nil {
		return nil, model.NewDBError(ctx, err.Error())
	}

	return entUser, nil
}

func (r *userRepository) Create(ctx context.Context, input generated.NewUser) (*model.User, error) {
	hashedPassword, err := util.HashPassword(input.Password)

	if err != nil {
		return nil, model.NewValidationError(ctx, err.Error(), "password")
	}

	entUser, err := r.client.User.Create().SetName(input.Username).SetPassword(hashedPassword).Save(ctx)

	if err != nil {
		return nil, model.NewDBError(ctx, err.Error())
	}

	return entUser, nil
}

func (r *userRepository) Update(ctx context.Context, entUser model.User, input generated.UpdateUser) (*model.User, error) {
	update := entUser.Update()

	if input.Name != nil {
		update = update.SetName(*input.Name)
	}

	if input.Password != nil {
		hashedPassword, err := util.HashPassword(*input.Password)
		if err != nil {
			return nil, model.NewValidationError(ctx, err.Error(), "password")
		}

		update = update.SetPassword(hashedPassword)
	}

	updatedEntUser, err := update.Save(ctx)

	if err != nil {
		return nil, model.NewDBError(ctx, err.Error())
	}

	return updatedEntUser, nil
}
