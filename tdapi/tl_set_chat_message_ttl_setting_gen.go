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

// SetChatMessageTTLSettingRequest represents TL type `setChatMessageTtlSetting#12d6f2f7`.
type SetChatMessageTTLSettingRequest struct {
	// Chat identifier
	ChatID int64
	// New TTL value, in seconds; must be one of 0, 86400, 7 * 86400, or 31 * 86400 unless
	// the chat is secret
	TTL int32
}

// SetChatMessageTTLSettingRequestTypeID is TL type id of SetChatMessageTTLSettingRequest.
const SetChatMessageTTLSettingRequestTypeID = 0x12d6f2f7

// Ensuring interfaces in compile-time for SetChatMessageTTLSettingRequest.
var (
	_ bin.Encoder     = &SetChatMessageTTLSettingRequest{}
	_ bin.Decoder     = &SetChatMessageTTLSettingRequest{}
	_ bin.BareEncoder = &SetChatMessageTTLSettingRequest{}
	_ bin.BareDecoder = &SetChatMessageTTLSettingRequest{}
)

func (s *SetChatMessageTTLSettingRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.ChatID == 0) {
		return false
	}
	if !(s.TTL == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SetChatMessageTTLSettingRequest) String() string {
	if s == nil {
		return "SetChatMessageTTLSettingRequest(nil)"
	}
	type Alias SetChatMessageTTLSettingRequest
	return fmt.Sprintf("SetChatMessageTTLSettingRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SetChatMessageTTLSettingRequest) TypeID() uint32 {
	return SetChatMessageTTLSettingRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SetChatMessageTTLSettingRequest) TypeName() string {
	return "setChatMessageTtlSetting"
}

// TypeInfo returns info about TL type.
func (s *SetChatMessageTTLSettingRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "setChatMessageTtlSetting",
		ID:   SetChatMessageTTLSettingRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "TTL",
			SchemaName: "ttl",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SetChatMessageTTLSettingRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setChatMessageTtlSetting#12d6f2f7 as nil")
	}
	b.PutID(SetChatMessageTTLSettingRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SetChatMessageTTLSettingRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setChatMessageTtlSetting#12d6f2f7 as nil")
	}
	b.PutInt53(s.ChatID)
	b.PutInt32(s.TTL)
	return nil
}

// Decode implements bin.Decoder.
func (s *SetChatMessageTTLSettingRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setChatMessageTtlSetting#12d6f2f7 to nil")
	}
	if err := b.ConsumeID(SetChatMessageTTLSettingRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode setChatMessageTtlSetting#12d6f2f7: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SetChatMessageTTLSettingRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setChatMessageTtlSetting#12d6f2f7 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode setChatMessageTtlSetting#12d6f2f7: field chat_id: %w", err)
		}
		s.ChatID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode setChatMessageTtlSetting#12d6f2f7: field ttl: %w", err)
		}
		s.TTL = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SetChatMessageTTLSettingRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode setChatMessageTtlSetting#12d6f2f7 as nil")
	}
	b.ObjStart()
	b.PutID("setChatMessageTtlSetting")
	b.FieldStart("chat_id")
	b.PutInt53(s.ChatID)
	b.FieldStart("ttl")
	b.PutInt32(s.TTL)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SetChatMessageTTLSettingRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode setChatMessageTtlSetting#12d6f2f7 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("setChatMessageTtlSetting"); err != nil {
				return fmt.Errorf("unable to decode setChatMessageTtlSetting#12d6f2f7: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode setChatMessageTtlSetting#12d6f2f7: field chat_id: %w", err)
			}
			s.ChatID = value
		case "ttl":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode setChatMessageTtlSetting#12d6f2f7: field ttl: %w", err)
			}
			s.TTL = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (s *SetChatMessageTTLSettingRequest) GetChatID() (value int64) {
	return s.ChatID
}

// GetTTL returns value of TTL field.
func (s *SetChatMessageTTLSettingRequest) GetTTL() (value int32) {
	return s.TTL
}

// SetChatMessageTTLSetting invokes method setChatMessageTtlSetting#12d6f2f7 returning error if any.
func (c *Client) SetChatMessageTTLSetting(ctx context.Context, request *SetChatMessageTTLSettingRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
