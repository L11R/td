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

// ApproveChatJoinRequestRequest represents TL type `approveChatJoinRequest#bcb04b05`.
type ApproveChatJoinRequestRequest struct {
	// Chat identifier
	ChatID int64
	// Identifier of the user, which request will be approved
	UserID int64
}

// ApproveChatJoinRequestRequestTypeID is TL type id of ApproveChatJoinRequestRequest.
const ApproveChatJoinRequestRequestTypeID = 0xbcb04b05

// Ensuring interfaces in compile-time for ApproveChatJoinRequestRequest.
var (
	_ bin.Encoder     = &ApproveChatJoinRequestRequest{}
	_ bin.Decoder     = &ApproveChatJoinRequestRequest{}
	_ bin.BareEncoder = &ApproveChatJoinRequestRequest{}
	_ bin.BareDecoder = &ApproveChatJoinRequestRequest{}
)

func (a *ApproveChatJoinRequestRequest) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.ChatID == 0) {
		return false
	}
	if !(a.UserID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *ApproveChatJoinRequestRequest) String() string {
	if a == nil {
		return "ApproveChatJoinRequestRequest(nil)"
	}
	type Alias ApproveChatJoinRequestRequest
	return fmt.Sprintf("ApproveChatJoinRequestRequest%+v", Alias(*a))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ApproveChatJoinRequestRequest) TypeID() uint32 {
	return ApproveChatJoinRequestRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ApproveChatJoinRequestRequest) TypeName() string {
	return "approveChatJoinRequest"
}

// TypeInfo returns info about TL type.
func (a *ApproveChatJoinRequestRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "approveChatJoinRequest",
		ID:   ApproveChatJoinRequestRequestTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *ApproveChatJoinRequestRequest) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode approveChatJoinRequest#bcb04b05 as nil")
	}
	b.PutID(ApproveChatJoinRequestRequestTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *ApproveChatJoinRequestRequest) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode approveChatJoinRequest#bcb04b05 as nil")
	}
	b.PutInt53(a.ChatID)
	b.PutInt53(a.UserID)
	return nil
}

// Decode implements bin.Decoder.
func (a *ApproveChatJoinRequestRequest) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode approveChatJoinRequest#bcb04b05 to nil")
	}
	if err := b.ConsumeID(ApproveChatJoinRequestRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode approveChatJoinRequest#bcb04b05: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *ApproveChatJoinRequestRequest) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode approveChatJoinRequest#bcb04b05 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode approveChatJoinRequest#bcb04b05: field chat_id: %w", err)
		}
		a.ChatID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode approveChatJoinRequest#bcb04b05: field user_id: %w", err)
		}
		a.UserID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (a *ApproveChatJoinRequestRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if a == nil {
		return fmt.Errorf("can't encode approveChatJoinRequest#bcb04b05 as nil")
	}
	b.ObjStart()
	b.PutID("approveChatJoinRequest")
	b.FieldStart("chat_id")
	b.PutInt53(a.ChatID)
	b.FieldStart("user_id")
	b.PutInt53(a.UserID)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (a *ApproveChatJoinRequestRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if a == nil {
		return fmt.Errorf("can't decode approveChatJoinRequest#bcb04b05 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("approveChatJoinRequest"); err != nil {
				return fmt.Errorf("unable to decode approveChatJoinRequest#bcb04b05: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode approveChatJoinRequest#bcb04b05: field chat_id: %w", err)
			}
			a.ChatID = value
		case "user_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode approveChatJoinRequest#bcb04b05: field user_id: %w", err)
			}
			a.UserID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (a *ApproveChatJoinRequestRequest) GetChatID() (value int64) {
	return a.ChatID
}

// GetUserID returns value of UserID field.
func (a *ApproveChatJoinRequestRequest) GetUserID() (value int64) {
	return a.UserID
}

// ApproveChatJoinRequest invokes method approveChatJoinRequest#bcb04b05 returning error if any.
func (c *Client) ApproveChatJoinRequest(ctx context.Context, request *ApproveChatJoinRequestRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
