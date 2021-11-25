$LOAD_PATH.unshift(File.dirname(__FILE__))

require_relative 'event-stream-filter_services_pb'

stub = EventStreamFilter::EventStream::Stub.new('localhost:8080', :this_channel_is_insecure)

r = EventStreamFilter::RegisterRequest.new(
  condition: EventStreamFilter::Condition.new(
    id: 'test',
    tenant_id: '1',
    event_type: 'CREATED',
    predicates: [
      EventStreamFilter::Predicate.new(
        name: 'name',
        operator: 'eq',
        value: Google::Protobuf::Value.new(string_value: 'test'),
      )
    ]
  )
)

resp = stub.register_condition(r)
p "- found #{resp.inspect}"
