$LOAD_PATH.unshift(File.dirname(__FILE__))

require_relative 'event-stream-filter_services_pb'

stub = EventStreamFilter::EventStream::Stub.new('localhost:8080', :this_channel_is_insecure)

r = EventStreamFilter::FilterRequest.new(

)

resp = stub.filter(r)
p "- found #{resp.inspect}"
