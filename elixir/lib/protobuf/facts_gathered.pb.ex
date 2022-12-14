defmodule Trento.Checks.V1.FactError do
  @moduledoc false
  use Protobuf, protoc_gen_elixir_version: "0.11.0", syntax: :proto3

  field :message, 1, type: :string
  field :type, 2, type: :string
end

defmodule Trento.Checks.V1.Fact do
  @moduledoc false
  use Protobuf, protoc_gen_elixir_version: "0.11.0", syntax: :proto3

  oneof :value, 0

  field :check_id, 1, type: :string, json_name: "checkId"
  field :name, 2, type: :string
  field :text_value, 3, type: :string, json_name: "textValue", oneof: 0
  field :numeric_value, 4, type: :float, json_name: "numericValue", oneof: 0
  field :error_value, 5, type: Trento.Checks.V1.FactError, json_name: "errorValue", oneof: 0
end

defmodule Trento.Checks.V1.FactsGathered do
  @moduledoc false
  use Protobuf, protoc_gen_elixir_version: "0.11.0", syntax: :proto3

  field :execution_id, 1, type: :string, json_name: "executionId"
  field :agent_id, 2, type: :string, json_name: "agentId"

  field :facts_gathered, 3,
    repeated: true,
    type: Trento.Checks.V1.Fact,
    json_name: "factsGathered"
end