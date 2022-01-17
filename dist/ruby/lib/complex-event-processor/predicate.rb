# frozen_string_literal: true

module EventStreamFilter
  class Predicate
    attr_reader :name, :operator, :value

    def initialize(name:, operator:, value:)
      @name = name
      @operator = operator
      @value = value
    end
  end
end
