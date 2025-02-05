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

// GetChatFilterRequest represents TL type `getChatFilter#6cdb65eb`.
type GetChatFilterRequest struct {
	// Chat filter identifier
	ChatFilterID int32
}

// GetChatFilterRequestTypeID is TL type id of GetChatFilterRequest.
const GetChatFilterRequestTypeID = 0x6cdb65eb

// Ensuring interfaces in compile-time for GetChatFilterRequest.
var (
	_ bin.Encoder     = &GetChatFilterRequest{}
	_ bin.Decoder     = &GetChatFilterRequest{}
	_ bin.BareEncoder = &GetChatFilterRequest{}
	_ bin.BareDecoder = &GetChatFilterRequest{}
)

func (g *GetChatFilterRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ChatFilterID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetChatFilterRequest) String() string {
	if g == nil {
		return "GetChatFilterRequest(nil)"
	}
	type Alias GetChatFilterRequest
	return fmt.Sprintf("GetChatFilterRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetChatFilterRequest) TypeID() uint32 {
	return GetChatFilterRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetChatFilterRequest) TypeName() string {
	return "getChatFilter"
}

// TypeInfo returns info about TL type.
func (g *GetChatFilterRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getChatFilter",
		ID:   GetChatFilterRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatFilterID",
			SchemaName: "chat_filter_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetChatFilterRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatFilter#6cdb65eb as nil")
	}
	b.PutID(GetChatFilterRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetChatFilterRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatFilter#6cdb65eb as nil")
	}
	b.PutInt32(g.ChatFilterID)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetChatFilterRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatFilter#6cdb65eb to nil")
	}
	if err := b.ConsumeID(GetChatFilterRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getChatFilter#6cdb65eb: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetChatFilterRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatFilter#6cdb65eb to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getChatFilter#6cdb65eb: field chat_filter_id: %w", err)
		}
		g.ChatFilterID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetChatFilterRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatFilter#6cdb65eb as nil")
	}
	b.ObjStart()
	b.PutID("getChatFilter")
	b.FieldStart("chat_filter_id")
	b.PutInt32(g.ChatFilterID)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetChatFilterRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatFilter#6cdb65eb to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getChatFilter"); err != nil {
				return fmt.Errorf("unable to decode getChatFilter#6cdb65eb: %w", err)
			}
		case "chat_filter_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode getChatFilter#6cdb65eb: field chat_filter_id: %w", err)
			}
			g.ChatFilterID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatFilterID returns value of ChatFilterID field.
func (g *GetChatFilterRequest) GetChatFilterID() (value int32) {
	return g.ChatFilterID
}

// GetChatFilter invokes method getChatFilter#6cdb65eb returning error if any.
func (c *Client) GetChatFilter(ctx context.Context, chatfilterid int32) (*ChatFilter, error) {
	var result ChatFilter

	request := &GetChatFilterRequest{
		ChatFilterID: chatfilterid,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
