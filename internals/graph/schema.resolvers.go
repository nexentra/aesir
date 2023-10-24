package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"

	"github.com/nexentra/aesir/internals/graph/model"
)

// EvaluateSnippet is the resolver for the EvaluateSnippet field.
func (r *mutationResolver) EvaluateSnippet(ctx context.Context, input model.EvalInput) (*model.Eval, error) {
	randNumber := input.Snippet
	eval := &model.Eval{
		Snippet: randNumber,
		Result:  randNumber,
	}
	r.eval = append(r.eval, eval)
	return eval, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
