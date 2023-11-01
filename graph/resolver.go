package graph

import (
	"sync"

	"github.com/nexentra/aesir/ent"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	EvalObservers map[string]chan []*ent.Eval
	mu            sync.Mutex
	Client        *ent.Client
}
