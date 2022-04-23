require 'complex-event-processor/client'

c = ComplexEventProcessor::Client.new

condition = ComplexEventProcessor::Condition.new(
  id: 'test',
  tenant_id: '1',
  event_type: 'CREATED',
  desired_result: true,
  predicates: [
    ComplexEventProcessor::Predicate.new(
      name: 'name',
      operator: 'eq',
      value: 'test',
    )
  ],
)

c.register_condition(condition)

event = ComplexEventProcessor::Event.new(
  id: 'test',
  tenant_id: '1',
  kind: 'CREATED',
  payload:  {
    name: 'test',
  }
)

response = c.filter(event)
hash = JSON.parse(response.to_json)
pp hash
