package schema

import (
	"context"
	"errors"

	"github.com/firstcontributions/backend/internal/gateway/session"
	"github.com/firstcontributions/backend/internal/storemanager"
)

func (m *Resolver) CreateUser(
	ctx context.Context,
	args struct {
		User *CreateUserInput
	},
) (*User, error) {
	session := session.FromContext(ctx)
	if session == nil {
		return nil, errors.New("unauthorized")
	}

	userModelInput, err := args.User.ToModel()
	if err != nil {
		return nil, err
	}

	user, err := storemanager.FromContext(ctx).UsersStore.CreateUser(ctx, userModelInput)
	if err != nil {
		return nil, err
	}
	return NewUser(user), nil
}
func (m *Resolver) UpdateUser(
	ctx context.Context,
	args struct {
		User *UpdateUserInput
	},
) (*User, error) {
	store := storemanager.FromContext(ctx)

	id, err := ParseGraphqlID(args.User.ID)
	if err != nil {
		return nil, err
	}
	if err := store.UsersStore.UpdateUser(ctx, id.ID, args.User.ToModel()); err != nil {
		return nil, err
	}
	user, err := store.UsersStore.GetUserByID(ctx, id.ID)
	if err != nil {
		return nil, err
	}
	return NewUser(user), nil
}
