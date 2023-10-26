package graph

import (
	"sync"

	"github.com/nexentra/aesir/ent"
	"github.com/nexentra/aesir/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Eval          []*model.Eval
	EvalObservers map[string]chan []*model.Eval
	mu            sync.Mutex
	Client        *ent.Client
}
