package state

import (
	"context"

	"github.com/jeangnc/complex-event-processor/pkg/types"
)

type Repository interface {
	Save(ctx context.Context, event types.Event, impact types.Impact) error
	Load(ctx context.Context, event types.Event, expressions []*types.Expression) (map[*types.Expression]types.State, error)
}
