package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	"fmt"

	"github.com/Jonathan-Zollinger/Tuirl/graph/model"
	"github.com/charmbracelet/log"
	"github.com/sony/sonyflake"
)

var sf *sonyflake.Sonyflake

// CreateNotecard is the resolver for the createNotecard field.
func (r *mutationResolver) CreateNotecard(ctx context.Context, input model.NewNotecard) (*model.Notecard, error) {
	newID, err := sf.NextID()
	if err != nil {
		log.Fatal("failed to create a unique ID when calling sonyflake.Sonyflake for CreateNotecard")
	}
	newNotecard := &model.Notecard{
		ID:    fmt.Sprintf("T%d", newID),
		Title: input.Title,
		Entry: input.Entry,
	}
	r.notecards = append(r.notecards, newNotecard)
	return newNotecard, nil
}

// CreateSection is the resolver for the createSection field.
func (r *mutationResolver) CreateSection(ctx context.Context, input model.NewSection) (*model.Section, error) {
	newID, err := sf.NextID()
	if err != nil {
		log.Fatal("failed to create a unique ID when calling sonyflake.Sonyflake for CreateNotecard")
	}
	newSection := &model.Section{
		ID:       fmt.Sprintf("T%d", newID),
		Title:    input.Title,
		SubTitle: input.Subtitle,
	}
	r.sections = append(r.sections, newSection)
	return newSection, nil
}

// Notecards is the resolver for the notecards field.
func (r *queryResolver) Notecards(ctx context.Context) ([]*model.Notecard, error) {
	return r.notecards, nil
}

// Sections is the resolver for the sections field.
func (r *queryResolver) Sections(ctx context.Context) ([]*model.Section, error) {
	return r.sections, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
