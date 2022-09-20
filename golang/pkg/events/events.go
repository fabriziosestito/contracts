package events

import (
	"time"

	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

type cloudEventOptions struct {
	id     *string
	source *string
	time   *time.Time
}

type Option = func(fields *cloudEventOptions)

func WithId(id *string) Option {
	return func(fields *cloudEventOptions) {
		fields.id = id
	}
}

func WithSource(source *string) Option {
	return func(fields *cloudEventOptions) {
		fields.source = source
	}
}

func WithTime(time *time.Time) Option {
	return func(fields *cloudEventOptions) {
		fields.time = time
	}
}

func ToEvent(event proto.Message, opts ...Option) ([]byte, error) {
	fields := &cloudEventOptions{}
	for _, opt := range opts {
		opt(fields)
	}

	data, err := anypb.New(event)
	if err != nil {
		return nil, err
	}

	var defaultId string
	if fields.id == nil {
		defaultId = uuid.New().String()
		fields.id = &defaultId
	}

	var defaultTime time.Time
	if fields.time == nil {
		defaultTime = time.Now()
		fields.time = &defaultTime
	}

	var defaultSource string
	if fields.source == nil {
		defaultSource = "trento"
		fields.source = &defaultSource
	}

	attr := CloudEventAttributeValue{
		Attr: &CloudEventAttributeValue_CeTimestamp{
			CeTimestamp: timestamppb.New(*fields.time),
		},
	}

	ce := CloudEvent{
		Id:          *fields.id,
		Source:      *fields.source,
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
