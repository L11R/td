// Code generated by gotdgen, DO NOT EDIT.

package tdapi

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// MessageSendOptions represents TL type `messageSendOptions#3682d6ba`.
type MessageSendOptions struct {
	// Pass true to disable notification for the message
	DisableNotification bool
	// Pass true if the message is sent from the background
	FromBackground bool
	// Message scheduling state; pass null to send message immediately. Messages sent to a
	// secret chat, live location messages and self-destructing messages can't be scheduled
	SchedulingState MessageSchedulingStateClass
}

// MessageSendOptionsTypeID is TL type id of MessageSendOptions.
const MessageSendOptionsTypeID = 0x3682d6ba

// Ensuring interfaces in compile-time for MessageSendOptions.
var (
	_ bin.Encoder     = &MessageSendOptions{}
	_ bin.Decoder     = &MessageSendOptions{}
	_ bin.BareEncoder = &MessageSendOptions{}
	_ bin.BareDecoder = &MessageSendOptions{}
)

func (m *MessageSendOptions) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.DisableNotification == false) {
		return false
	}
	if !(m.FromBackground == false) {
		return false
	}
	if !(m.SchedulingState == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessageSendOptions) String() string {
	if m == nil {
		return "MessageSendOptions(nil)"
	}
	type Alias MessageSendOptions
	return fmt.Sprintf("MessageSendOptions%+v", Alias(*m))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessageSendOptions) TypeID() uint32 {
	return MessageSendOptionsTypeID
}

// TypeName returns name of type in TL schema.
func (*MessageSendOptions) TypeName() string {
	return "messageSendOptions"
}

// TypeInfo returns info about TL type.
func (m *MessageSendOptions) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messageSendOptions",
		ID:   MessageSendOptionsTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "DisableNotification",
			SchemaName: "disable_notification",
		},
		{
			Name:       "FromBackground",
			SchemaName: "from_background",
		},
		{
			Name:       "SchedulingState",
			SchemaName: "scheduling_state",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (m *MessageSendOptions) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageSendOptions#3682d6ba as nil")
	}
	b.PutID(MessageSendOptionsTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MessageSendOptions) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageSendOptions#3682d6ba as nil")
	}
	b.PutBool(m.DisableNotification)
	b.PutBool(m.FromBackground)
	if m.SchedulingState == nil {
		return fmt.Errorf("unable to encode messageSendOptions#3682d6ba: field scheduling_state is nil")
	}
	if err := m.SchedulingState.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messageSendOptions#3682d6ba: field scheduling_state: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (m *MessageSendOptions) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageSendOptions#3682d6ba to nil")
	}
	if err := b.ConsumeID(MessageSendOptionsTypeID); err != nil {
		return fmt.Errorf("unable to decode messageSendOptions#3682d6ba: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MessageSendOptions) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageSendOptions#3682d6ba to nil")
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode messageSendOptions#3682d6ba: field disable_notification: %w", err)
		}
		m.DisableNotification = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode messageSendOptions#3682d6ba: field from_background: %w", err)
		}
		m.FromBackground = value
	}
	{
		value, err := DecodeMessageSchedulingState(b)
		if err != nil {
			return fmt.Errorf("unable to decode messageSendOptions#3682d6ba: field scheduling_state: %w", err)
		}
		m.SchedulingState = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (m *MessageSendOptions) EncodeTDLibJSON(b tdjson.Encoder) error {
	if m == nil {
		return fmt.Errorf("can't encode messageSendOptions#3682d6ba as nil")
	}
	b.ObjStart()
	b.PutID("messageSendOptions")
	b.FieldStart("disable_notification")
	b.PutBool(m.DisableNotification)
	b.FieldStart("from_background")
	b.PutBool(m.FromBackground)
	b.FieldStart("scheduling_state")
	if m.SchedulingState == nil {
		return fmt.Errorf("unable to encode messageSendOptions#3682d6ba: field scheduling_state is nil")
	}
	if err := m.SchedulingState.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode messageSendOptions#3682d6ba: field scheduling_state: %w", err)
	}
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (m *MessageSendOptions) DecodeTDLibJSON(b tdjson.Decoder) error {
	if m == nil {
		return fmt.Errorf("can't decode messageSendOptions#3682d6ba to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("messageSendOptions"); err != nil {
				return fmt.Errorf("unable to decode messageSendOptions#3682d6ba: %w", err)
			}
		case "disable_notification":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode messageSendOptions#3682d6ba: field disable_notification: %w", err)
			}
			m.DisableNotification = value
		case "from_background":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode messageSendOptions#3682d6ba: field from_background: %w", err)
			}
			m.FromBackground = value
		case "scheduling_state":
			value, err := DecodeTDLibJSONMessageSchedulingState(b)
			if err != nil {
				return fmt.Errorf("unable to decode messageSendOptions#3682d6ba: field scheduling_state: %w", err)
			}
			m.SchedulingState = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetDisableNotification returns value of DisableNotification field.
func (m *MessageSendOptions) GetDisableNotification() (value bool) {
	return m.DisableNotification
}

// GetFromBackground returns value of FromBackground field.
func (m *MessageSendOptions) GetFromBackground() (value bool) {
	return m.FromBackground
}

// GetSchedulingState returns value of SchedulingState field.
func (m *MessageSendOptions) GetSchedulingState() (value MessageSchedulingStateClass) {
	return m.SchedulingState
}
