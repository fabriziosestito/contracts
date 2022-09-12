defmodule Proto do
  @moduledoc """
  Documentation for `Proto`.
  """

  @doc """
  Encode and wrap an event struct to a protobuf CloudEvent.
  """
  @spec to_event(struct(), Keyword.t()) :: binary()
  def to_event(%mod{} = struct, opts \\ []) do
    id = Keyword.get(opts, :id, UUID.uuid4())
    source = Keyword.get(opts, :source, "trento")
    data = Protobuf.Encoder.encode(struct)

    cloud_event =
      CloudEvents.CloudEvent.new(
        data: {:proto_data, Google.Protobuf.Any.new!(value: data, type_url: get_type(mod))},
        spec_version: "1.0",
        type: get_type(mod),
        id: id,
        source: source
      )

    cloud_event
    # Keeping the unknonw fields causes dialyzer to complain
    |> Map.drop([:__unknown_fields__])
    |> CloudEvents.CloudEvent.encode()
  end

  @doc """
  Decode and unwrap a protobuf CloudEvent to an event struct.
  """
  @spec from_event(binary()) ::
          {:ok, struct()}
          | {:error, :invalid_cloud_event}
          | {:error, :decoding_error}
          | {:error, :event_not_found}
  def from_event(value) do
    try do
      case CloudEvents.CloudEvent.decode(value) do
        %{type: type, data: {:proto_data, %Google.Protobuf.Any{value: data}}} ->
          decode(type, data)

        _ ->
          {:error, :invalid_cloud_event}
      end
    rescue
      error -> {:error, :decoding_error}
    end
  end

  defp decode(type, data) do
    try do
      module_name = Macro.camelize(type)
      module = Module.safe_concat([module_name])

      {:ok, module.decode(data)}
    rescue
      ArgumentError -> {:error, :event_not_found}
    end
  end

  defp get_type(mod) do
    mod
    |> Atom.to_string()
    |> String.replace("Elixir.", "")
  end
end
