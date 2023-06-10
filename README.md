# complex-event-processor
Match an event stream against a predefined set of access patterns to identify relevant events.


```sh
curl -X POST localhost:8080/expression \
    -H 'Content-Type: application/json ' \
    -d '{
	  "id": "complex-event-1",
	  "logical_expression": {
		"connector": "and",
		"operands": [
		  {
			"predicate": {
			  "id": "A",
			  "event_type": "EMAIL_OPENED",
			  "conditions": [
				{
				  "field": "age",
				  "operator": "less_than",
				  "value": 18
				},
				{
				  "field": "provider",
				  "operator": "not_equal",
				  "value": "gmail"
				}
			  ],
			  "immutable": true
			}
		  },
		  {
			"predicate": {
			  "id": "B",
			  "event_type": "EMAIL_CLICKED",
			  "conditions": [
				{
				  "field": "link",
				  "operator": "equal",
				  "value": "http://google.com"
				}
			  ],
			  "immutable": true
			}
		  }
		]
	  }
	}'

curl -X POST localhost:8080/event \
    -H 'Content-Type: application/json ' \
    -d '{
      "id": "1",
      "timestamp": 1,
      "type": "EMAIL_OPENED",
      "payload": {
        "email": "john@example.org"
      }
    }'

curl -X POST localhost:8080/event \
    -H 'Content-Type: application/json ' \
    -d '{
      "id": "1",
      "timestamp": 1,
      "type": "EMAIL_OPENED",
      "payload": {
        "age": 17,
        "provider": "hotmail"
      }
    }'

curl -X POST localhost:8080/event \
    -H 'Content-Type: application/json ' \
    -d '{
      "id": "2",
      "timestamp": 1,
      "type": "EMAIL_CLICKED",
      "payload": {
        "link": "http://google.com"
      }
    }'
```
