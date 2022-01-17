# frozen_string_literal: true

module ComplexEventProcessor
  class Event
    attr_reader :id, :tenant_id, :kind, :payload

    def initialize(id:, tenant_id:, kind:, payload:)
      @id = id
      @tenant_id = tenant_id
      @kind = kind
      @payload = payload
    end
  end
end
