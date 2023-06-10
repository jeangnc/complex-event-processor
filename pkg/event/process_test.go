package event

import (
	"reflect"
	"testing"

	"github.com/jeangnc/complex-event-processor/pkg/expression"
	"github.com/jeangnc/complex-event-processor/pkg/state"
	"github.com/jeangnc/complex-event-processor/pkg/types"
)

//
func TestEventProcessing(t *testing.T) {
	type testCase struct {
		description    string
		event          types.Event
		predicate      types.Predicate
		expectedResult map[string]bool
	}

	p := types.Predicate{
		Id:         "my-predicate",
		EventType:  "MY_TYPE",
		Conditions: []types.Condition{},
	}

	e := types.Event{
		Type:    "MY_TYPE",
		Payload: map[string]interface{}{},
	}

	testCases := []testCase{
		testCase{
			description: "Basic event processing",
			event:       e,
			predicate:   p,
			expectedResult: map[string]bool{
				"1": true,
			},
		},
	}

	for _, s := range testCases {
		t.Run(s.description, func(t *testing.T) {
			ex := types.Expression{
				Id: "1",
				LogicalExpression: types.LogicalExpression{
					Connector: types.CONNECTOR_AND,
					Operands: []types.Operand{
						types.Operand{Predicate: &s.predicate},
					},
				},
			}

			index := expression.NewTemporaryIndex()
			index.Append(ex)

			repository := state.NewRedisRepository("localhost:6379", "")

			result := Process(&index, &repository, s.event)

			if !reflect.DeepEqual(result, s.expectedResult) {
				t.Fatalf(`Failed to search impacted expressions: %v %v`, result, s.expectedResult)
			}
		})
	}
}
