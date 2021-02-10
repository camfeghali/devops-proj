package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"hackernews-clone/database"
	"hackernews-clone/graph/model"
	"context"
	"fmt"
	"hackernews-clone/graph/generated"

)

func (r *mutationResolver) CreateLink(ctx context.Context, input model.NewLink) (*model.Link, error) {
	link := model.Link{Title: input.Title, Address: input.Address, UserID: 1}

	var user model.User

	database.Db.Create(&link)
	database.Db.First(&user)

	return &model.Link{
		ID:      link.ID,
		Title:   link.Title,
		Address: link.Address,
		User:    &user,
	}, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := model.User{Username: input.Username, Password: input.Password}

	database.Db.Create(&user)
	return &model.User{ID: user.ID, Username: user.Username, Password: user.Password}, nil
}


func (r *queryResolver) Links(ctx context.Context) ([]*model.Link, error) {
	var resultLinks []*model.Link
	var dbLinks []model.Link

	database.Db.Preload("User").Find(&dbLinks)

	for _, link := range dbLinks {

		resultLinks = append(
			resultLinks,
			&model.Link{
				ID:      link.ID,
				Title:   link.Title,
				Address: link.Address,
				User:    link.User,
			})
	}

	return resultLinks, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	var resultUsers []*model.User
	var dbUsers []model.User

	database.Db.Find(&dbUsers)

	for _, user := range dbUsers {
		resultUsers = append(resultUsers, &model.User{ID: user.ID, Username: user.Username, Password: user.Password})
	}

	return resultUsers, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
