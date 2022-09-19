package events

import (
	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

type CloudEventOptions struct {
	Id     string
	Source string
	Time   *timestamppb.Timestamp
}

func ToEvent(event proto.Message, opts CloudEventOptions) ([]byte, error) {
	if opts.Id == "" {
		opts.Id = uuid.New().String()
	}

	if opts.Source == "" {
		opts.Source = "trento"
	}

	if opts.Time == nil {
		opts.Time = timestamppb.Now()
	}

	data, err := anypb.New(event)
	if err != nil {
		return nil, err
	}

	attr := CloudEventAttributeValue{
		Attr: &CloudEventAttributeValue_CeTimestamp{
			CeTimestamp: opts.Time,
		},
	}

	ce := CloudEvent{
		Id:          opts.Id,
		Source:      opts.Source,
		SpecVersion: "1.0",
		Type:        eventTypeFromProto(event),
		Data: &CloudEvent_ProtoData{
			ProtoData: data,
		},
		Attributes: map[string]*CloudEventAttributeValue{
			"time": &attr,
		},
	}

	rawCe, err := proto.Marshal(&ce)
	if err != nil {
		return nil, err
	}

	return rawCe, nil
}

func FromEvent(src []byte, to proto.Message) error {
	var decodedCe CloudEvent
	err := proto.Unmarshal(src, &decodedCe)
	if err != nil {
		return err
	}

	err = anypb.UnmarshalTo(decodedCe.GetProtoData(), to, proto.UnmarshalOptions{})
	if err != nil {
		return err
	}

	return nil
}

func eventTypeFromProto(message proto.Message) string {
	return string(message.ProtoReflect().Descriptor().FullName())
}

func EventType(src []byte) (string, error) {
	var decodedCe CloudEvent
	err := proto.Unmarshal(src, &decodedCe)
	if err != nil {
		return "", err
	}

	any, err := anypb.UnmarshalNew(decodedCe.GetProtoData(), proto.UnmarshalOptions{})
	if err != nil {
		return "", err
	}

	name := any.ProtoReflect().Type().Descriptor().FullName()
	return string(name), nil
}
