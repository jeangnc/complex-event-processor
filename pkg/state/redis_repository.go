package state

import (
	"context"
	"strconv"
	"strings"

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
	expressionCommands := make(map[*types.Expression]map[string]*redis.Cmd)

	_, err := r.client.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		for _, e := range expressions {
			expressionCommands[e] = make(map[string]*redis.Cmd, 0)

			min := "-inf"
			max := "+inf"

			if e.Window > 0 {
				min = strconv.FormatInt(event.Timestamp-e.Window, 10)
				max = strconv.FormatInt(event.Timestamp+e.Window, 10)
			}

			for _, keys := range gambiKeysToLoad(e.LogicalExpression) {
				id := strings.Join(keys, ";")

				var fname string

				if len(keys) > 1 {
					fname = "zsequence"
				} else {
					fname = "zvalue"
				}

				expressionCommands[e][id] = pipe.FCall(ctx, fname, keys, min, max)
			}
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

	states := make(map[*types.Expression]types.State)

	for e, commands := range expressionCommands {
		values := make(map[string]bool)

		for id, command := range commands {
			keys := strings.Split(id, ";")

			if len(keys) > 1 {
				prefix := ""
				response, _ := command.Slice()

				for i, k := range keys {
					values[prefix+k] = response[i] != nil
					prefix += k + ";"
				}
			} else {
				response, _ := command.Slice()
				values[id] = len(response) > 0
			}
		}

		states[e] = types.State{Predicates: values}
	}

	return states, nil
}

func gambiKeysToLoad(l types.LogicalExpression) [][]string {
	r := make([][]string, 0, 0)

	if l.Connector == types.CONNECTOR_SEQUENCE {
		keys := make([]string, 0, 0)

		for _, o := range l.Operands {
			keys = append(keys, o.Predicate.Id)
		}

		r = append(r, keys)
		return r
	}

	for _, o := range l.Operands {
		if o.LogicalExpression != nil {
			r = append(r, gambiKeysToLoad(*o.LogicalExpression)...)
			continue
		}

		r = append(r, []string{o.Predicate.Id})
	}

	return r
}
