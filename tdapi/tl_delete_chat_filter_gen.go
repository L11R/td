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

// DeleteChatFilterRequest represents TL type `deleteChatFilter#e0d04873`.
type DeleteChatFilterRequest struct {
	// Chat filter identifier
	ChatFilterID int32
}

// DeleteChatFilterRequestTypeID is TL type id of DeleteChatFilterRequest.
const DeleteChatFilterRequestTypeID = 0xe0d04873

// Ensuring interfaces in compile-time for DeleteChatFilterRequest.
var (
	_ bin.Encoder     = &DeleteChatFilterRequest{}
	_ bin.Decoder     = &DeleteChatFilterRequest{}
	_ bin.BareEncoder = &DeleteChatFilterRequest{}
	_ bin.BareDecoder = &DeleteChatFilterRequest{}
)

func (d *DeleteChatFilterRequest) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.ChatFilterID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *DeleteChatFilterRequest) String() string {
	if d == nil {
		return "DeleteChatFilterRequest(nil)"
	}
	type Alias DeleteChatFilterRequest
	return fmt.Sprintf("DeleteChatFilterRequest%+v", Alias(*d))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*DeleteChatFilterRequest) TypeID() uint32 {
	return DeleteChatFilterRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*DeleteChatFilterRequest) TypeName() string {
	return "deleteChatFilter"
}

// TypeInfo returns info about TL type.
func (d *DeleteChatFilterRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "deleteChatFilter",
		ID:   DeleteChatFilterRequestTypeID,
	}
	if d == nil {
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
func (d *DeleteChatFilterRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode deleteChatFilter#e0d04873 as nil")
	}
	b.PutID(DeleteChatFilterRequestTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *DeleteChatFilterRequest) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode deleteChatFilter#e0d04873 as nil")
	}
	b.PutInt32(d.ChatFilterID)
	return nil
}

// Decode implements bin.Decoder.
func (d *DeleteChatFilterRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode deleteChatFilter#e0d04873 to nil")
	}
	if err := b.ConsumeID(DeleteChatFilterRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode deleteChatFilter#e0d04873: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *DeleteChatFilterRequest) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode deleteChatFilter#e0d04873 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode deleteChatFilter#e0d04873: field chat_filter_id: %w", err)
		}
		d.ChatFilterID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (d *DeleteChatFilterRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if d == nil {
		return fmt.Errorf("can't encode deleteChatFilter#e0d04873 as nil")
	}
	b.ObjStart()
	b.PutID("deleteChatFilter")
	b.FieldStart("chat_filter_id")
	b.PutInt32(d.ChatFilterID)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (d *DeleteChatFilterRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if d == nil {
		return fmt.Errorf("can't decode deleteChatFilter#e0d04873 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("deleteChatFilter"); err != nil {
				return fmt.Errorf("unable to decode deleteChatFilter#e0d04873: %w", err)
			}
		case "chat_filter_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode deleteChatFilter#e0d04873: field chat_filter_id: %w", err)
			}
			d.ChatFilterID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatFilterID returns value of ChatFilterID field.
func (d *DeleteChatFilterRequest) GetChatFilterID() (value int32) {
	return d.ChatFilterID
}

// DeleteChatFilter invokes method deleteChatFilter#e0d04873 returning error if any.
func (c *Client) DeleteChatFilter(ctx context.Context, chatfilterid int32) error {
	var ok Ok

	request := &DeleteChatFilterRequest{
		ChatFilterID: chatfilterid,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
