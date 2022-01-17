# frozen_string_literal: true

require_relative 'predicate'

module ComplexEventProcessor
  class Condition
    attr_reader :id, :tenant_id, :event_type, :predicates, :desired_result

    def initialize(id:, tenant_id:, event_type:, predicates:, desired_result:)
      @id = id
      @tenant_id = tenant_id
      @event_type = event_type
      @predicates = predicates
      @desired_result = desired_result
    end
  end
end
