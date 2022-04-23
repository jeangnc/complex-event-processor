# frozen_string_literal: true

require_relative "lib/complex-event-processor/version"

Gem::Specification.new do |spec|
  spec.name          = "complex-event-processor"
  spec.version       = ComplexEventProcessor::VERSION
  spec.authors       = ["Jean Carlos Gonçalves"]
  spec.email         = ["jean.gnc@gmail.com"]

  spec.summary       = "API client library for Event Stream Filter API"
  spec.description   = "complex-event-processor is the official client for Event Stream Filter API"
  spec.homepage      = "https://github.com/jeangnc/complex-event-processor/tree/main/dist/ruby/"
  spec.license       = "MIT"
  spec.required_ruby_version = ">= 2.5.0"

  spec.metadata["homepage_uri"] = spec.homepage
  spec.metadata["source_code_uri"] = "https://github.com/jeangnc/complex-event-processor/tree/main/dist/ruby/"
  spec.metadata["changelog_uri"] = "https://github.com/jeangnc/complex-event-processor/tree/main/dist/ruby/CHANGELOG.md"

  # Specify which files should be added to the gem when it is released.
  # The `git ls-files -z` loads the files in the RubyGem that have been added into git.
  spec.files = Dir.chdir(File.expand_path(__dir__)) do
    `git ls-files -z`.split("\x0").reject { |f| f.match(%r{\A(?:test|spec|features)/}) }
  end
  spec.bindir        = "exe"
  spec.executables   = spec.files.grep(%r{\Aexe/}) { |f| File.basename(f) }
  spec.require_paths = ["lib", "lib/complex-event-processor/proto"]

  spec.add_dependency "grpc", "~> 1.41"
  spec.add_development_dependency "rake", "~> 13.0"
  spec.add_development_dependency "rspec", "~> 3.0"
end
