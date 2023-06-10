package state

import (
	"context"
	"strconv"

	"github.com/jeangnc/complex-event-processor/pkg/types"
	"github.com/redis/go-redis/v9"
)

type RedisRepository struct {
	client *redis.Client
}

func NewRedisRepository(addr string, password string) RedisRepository {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
	})

	return RedisRepository{
		client: client,
	}
}

func (r RedisRepository) Save(ctx context.Context, event types.Event, impact types.Impact) error {
	_, err := r.client.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		for k := range impact.Predicates {
			member := redis.Z{
				Score:  float64(event.Timestamp),
				Member: event.Id,
			}

			pipe.ZAdd(ctx, k, member)
		}

		return nil
	})

	return err
}

func (r RedisRepository) Load(ctx context.Context, event types.Event, expressions []*types.Expression) (map[*types.Expression]types.State, error) {
	promises := make(map[*types.Expression]map[string]*redis.ZSliceCmd)

	_, err := r.client.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		for _, e := range expressions {
			keys := extractKeys(&e.LogicalExpression)
			promises[e] = make(map[string]*redis.ZSliceCmd, 0)

			for _, k := range keys {
				min := "-inf"
				max := "+inf"

				if e.Window > 0 {
					min = strconv.FormatInt(event.Timestamp-e.Window, 10)
					max = strconv.FormatInt(event.Timestamp+e.Window, 10)
				}

				opts := &redis.ZRangeBy{
					Min:   min,
					Max:   max,
					Count: 1,
				}

				promises[e][k] = pipe.ZRangeByScoreWithScores(ctx, k, opts)
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	states := make(map[*types.Expression]types.State)

	for _, e := range expressions {
		values := make(map[string]bool)

		for predicateId, promise := range promises[e] {
			response := promise.Val()
			values[predicateId] = len(response) > 0
		}

		states[e] = types.State{Predicates: values}
	}

	return states, nil
}

func extractKeys(l *types.LogicalExpression) []string {
	values := make([]string, 0, 0)

	for _, o := range l.Operands {
		if o.LogicalExpression != nil {
			values = append(values, extractKeys(o.LogicalExpression)...)
			continue
		}

		values = append(values, o.Predicate.Id)
	}

	return values
}
