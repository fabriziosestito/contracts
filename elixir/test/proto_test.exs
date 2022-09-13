defmodule ProtoTest do
  use ExUnit.Case

  alias CloudEvents.CloudEvent

  test "should decode to the right struct" do
    %Test.Event{id: event_id} = event = Test.Event.new(id: UUID.uuid4())

    cloudevent = %CloudEvent{
      data:
        {:proto_data,
         %Google.Protobuf.Any{
           __unknown_fields__: [],
           type_url: "test.Event",
           value: Test.Event.encode(event)
         }},
      id: UUID.uuid4(),
      source: "wandalorian",
      spec_version: "1.0",
      type: "test.Event"
    }

    encoded_cloudevent = CloudEvent.encode(cloudevent)

    assert {:ok, %Test.Event{id: ^event_id}} = Proto.from_event(encoded_cloudevent)
  end

  test "should encode to the right struct" do
    event = Test.Event.new(id: UUID.uuid4())
    cloudevent_id = UUID.uuid4()

    cloudevent = %CloudEvent{
      data:
        {:proto_data,
         %Google.Protobuf.Any{
           __unknown_fields__: [],
           type_url: "Test.Event",
           value: Test.Event.encode(event)
         }},
      id: cloudevent_id,
      source: "wandalorian",
      spec_version: "1.0",
      type: "Test.Event"
    }

    encoded_cloudevent = CloudEvent.encode(cloudevent)

    assert encoded_cloudevent == Proto.to_event(event, id: cloudevent_id, source: "wandalorian")
  end

  test "should return error if the event is not wrapped in a CloudEvent" do
    event = [id: UUID.uuid4()] |> Test.Event.new() |> Test.Event.encode()

    assert {:error, :invalid_envelope} = Proto.from_event(event)
  end

  test "should return error if the could not be decoded" do
    event = <<0, 0>>

    assert {:error, :decoding_error} = Proto.from_event(event)
  end

  test "should return error if the event type is unknown" do
    cloudevent = %CloudEvent{
      data:
        {:proto_data,
         %Google.Protobuf.Any{
           __unknown_fields__: [],
           type_url: "Unknown.Event",
           value: <<0, 0, 0, 0, 0, 0, 0, 0>>
         }},
      id: UUID.uuid4(),
      source: "wandalorian",
      spec_version: "1.0",
      type: "Unknown.Event"
    }

    assert {:error, :unknown_event} = cloudevent |> CloudEvent.encode() |> Proto.from_event()
  end
end
