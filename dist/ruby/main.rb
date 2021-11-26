$LOAD_PATH.unshift(File.dirname(__FILE__))

require 'json'
require_relative 'event_stream_filter_services_pb'

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
    ],
    desired_result: true
  )
)

# resp = stub.register_condition(r)
# p "- found #{resp.inspect}"

r = EventStreamFilter::FilterRequest.new(
  event: EventStreamFilter::Event.new(
    id: 'test',
    tenant_id: '1',
    kind: 'CREATED',
    payload:  Google::Protobuf::Struct.new(
      fields: {
        name: Google::Protobuf::Value.new(string_value: 'test'),
      }
    ),
  )
)

resp = stub.filter(r)
hash = JSON.parse(resp.to_json)
pp hash
