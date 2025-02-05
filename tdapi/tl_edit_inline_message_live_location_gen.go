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

// EditInlineMessageLiveLocationRequest represents TL type `editInlineMessageLiveLocation#f6a5da00`.
type EditInlineMessageLiveLocationRequest struct {
	// Inline message identifier
	InlineMessageID string
	// The new message reply markup; pass null if none
	ReplyMarkup ReplyMarkupClass
	// New location content of the message; pass null to stop sharing the live location
	Location Location
	// The new direction in which the location moves, in degrees; 1-360. Pass 0 if unknown
	Heading int32
	// The new maximum distance for proximity alerts, in meters (0-100000). Pass 0 if the
	// notification is disabled
	ProximityAlertRadius int32
}

// EditInlineMessageLiveLocationRequestTypeID is TL type id of EditInlineMessageLiveLocationRequest.
const EditInlineMessageLiveLocationRequestTypeID = 0xf6a5da00

// Ensuring interfaces in compile-time for EditInlineMessageLiveLocationRequest.
var (
	_ bin.Encoder     = &EditInlineMessageLiveLocationRequest{}
	_ bin.Decoder     = &EditInlineMessageLiveLocationRequest{}
	_ bin.BareEncoder = &EditInlineMessageLiveLocationRequest{}
	_ bin.BareDecoder = &EditInlineMessageLiveLocationRequest{}
)

func (e *EditInlineMessageLiveLocationRequest) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.InlineMessageID == "") {
		return false
	}
	if !(e.ReplyMarkup == nil) {
		return false
	}
	if !(e.Location.Zero()) {
		return false
	}
	if !(e.Heading == 0) {
		return false
	}
	if !(e.ProximityAlertRadius == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *EditInlineMessageLiveLocationRequest) String() string {
	if e == nil {
		return "EditInlineMessageLiveLocationRequest(nil)"
	}
	type Alias EditInlineMessageLiveLocationRequest
	return fmt.Sprintf("EditInlineMessageLiveLocationRequest%+v", Alias(*e))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*EditInlineMessageLiveLocationRequest) TypeID() uint32 {
	return EditInlineMessageLiveLocationRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*EditInlineMessageLiveLocationRequest) TypeName() string {
	return "editInlineMessageLiveLocation"
}

// TypeInfo returns info about TL type.
func (e *EditInlineMessageLiveLocationRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "editInlineMessageLiveLocation",
		ID:   EditInlineMessageLiveLocationRequestTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "InlineMessageID",
			SchemaName: "inline_message_id",
		},
		{
			Name:       "ReplyMarkup",
			SchemaName: "reply_markup",
		},
		{
			Name:       "Location",
			SchemaName: "location",
		},
		{
			Name:       "Heading",
			SchemaName: "heading",
		},
		{
			Name:       "ProximityAlertRadius",
			SchemaName: "proximity_alert_radius",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *EditInlineMessageLiveLocationRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode editInlineMessageLiveLocation#f6a5da00 as nil")
	}
	b.PutID(EditInlineMessageLiveLocationRequestTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *EditInlineMessageLiveLocationRequest) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode editInlineMessageLiveLocation#f6a5da00 as nil")
	}
	b.PutString(e.InlineMessageID)
	if e.ReplyMarkup == nil {
		return fmt.Errorf("unable to encode editInlineMessageLiveLocation#f6a5da00: field reply_markup is nil")
	}
	if err := e.ReplyMarkup.Encode(b); err != nil {
		return fmt.Errorf("unable to encode editInlineMessageLiveLocation#f6a5da00: field reply_markup: %w", err)
	}
	if err := e.Location.Encode(b); err != nil {
		return fmt.Errorf("unable to encode editInlineMessageLiveLocation#f6a5da00: field location: %w", err)
	}
	b.PutInt32(e.Heading)
	b.PutInt32(e.ProximityAlertRadius)
	return nil
}

// Decode implements bin.Decoder.
func (e *EditInlineMessageLiveLocationRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode editInlineMessageLiveLocation#f6a5da00 to nil")
	}
	if err := b.ConsumeID(EditInlineMessageLiveLocationRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode editInlineMessageLiveLocation#f6a5da00: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *EditInlineMessageLiveLocationRequest) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode editInlineMessageLiveLocation#f6a5da00 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode editInlineMessageLiveLocation#f6a5da00: field inline_message_id: %w", err)
		}
		e.InlineMessageID = value
	}
	{
		value, err := DecodeReplyMarkup(b)
		if err != nil {
			return fmt.Errorf("unable to decode editInlineMessageLiveLocation#f6a5da00: field reply_markup: %w", err)
		}
		e.ReplyMarkup = value
	}
	{
		if err := e.Location.Decode(b); err != nil {
			return fmt.Errorf("unable to decode editInlineMessageLiveLocation#f6a5da00: field location: %w", err)
		}
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode editInlineMessageLiveLocation#f6a5da00: field heading: %w", err)
		}
		e.Heading = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode editInlineMessageLiveLocation#f6a5da00: field proximity_alert_radius: %w", err)
		}
		e.ProximityAlertRadius = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (e *EditInlineMessageLiveLocationRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if e == nil {
		return fmt.Errorf("can't encode editInlineMessageLiveLocation#f6a5da00 as nil")
	}
	b.ObjStart()
	b.PutID("editInlineMessageLiveLocation")
	b.FieldStart("inline_message_id")
	b.PutString(e.InlineMessageID)
	b.FieldStart("reply_markup")
	if e.ReplyMarkup == nil {
		return fmt.Errorf("unable to encode editInlineMessageLiveLocation#f6a5da00: field reply_markup is nil")
	}
	if err := e.ReplyMarkup.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode editInlineMessageLiveLocation#f6a5da00: field reply_markup: %w", err)
	}
	b.FieldStart("location")
	if err := e.Location.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode editInlineMessageLiveLocation#f6a5da00: field location: %w", err)
	}
	b.FieldStart("heading")
	b.PutInt32(e.Heading)
	b.FieldStart("proximity_alert_radius")
	b.PutInt32(e.ProximityAlertRadius)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (e *EditInlineMessageLiveLocationRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if e == nil {
		return fmt.Errorf("can't decode editInlineMessageLiveLocation#f6a5da00 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("editInlineMessageLiveLocation"); err != nil {
				return fmt.Errorf("unable to decode editInlineMessageLiveLocation#f6a5da00: %w", err)
			}
		case "inline_message_id":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode editInlineMessageLiveLocation#f6a5da00: field inline_message_id: %w", err)
			}
			e.InlineMessageID = value
		case "reply_markup":
			value, err := DecodeTDLibJSONReplyMarkup(b)
			if err != nil {
				return fmt.Errorf("unable to decode editInlineMessageLiveLocation#f6a5da00: field reply_markup: %w", err)
			}
			e.ReplyMarkup = value
		case "location":
			if err := e.Location.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode editInlineMessageLiveLocation#f6a5da00: field location: %w", err)
			}
		case "heading":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode editInlineMessageLiveLocation#f6a5da00: field heading: %w", err)
			}
			e.Heading = value
		case "proximity_alert_radius":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode editInlineMessageLiveLocation#f6a5da00: field proximity_alert_radius: %w", err)
			}
			e.ProximityAlertRadius = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetInlineMessageID returns value of InlineMessageID field.
func (e *EditInlineMessageLiveLocationRequest) GetInlineMessageID() (value string) {
	return e.InlineMessageID
}

// GetReplyMarkup returns value of ReplyMarkup field.
func (e *EditInlineMessageLiveLocationRequest) GetReplyMarkup() (value ReplyMarkupClass) {
	return e.ReplyMarkup
}

// GetLocation returns value of Location field.
func (e *EditInlineMessageLiveLocationRequest) GetLocation() (value Location) {
	return e.Location
}

// GetHeading returns value of Heading field.
func (e *EditInlineMessageLiveLocationRequest) GetHeading() (value int32) {
	return e.Heading
}

// GetProximityAlertRadius returns value of ProximityAlertRadius field.
func (e *EditInlineMessageLiveLocationRequest) GetProximityAlertRadius() (value int32) {
	return e.ProximityAlertRadius
}

// EditInlineMessageLiveLocation invokes method editInlineMessageLiveLocation#f6a5da00 returning error if any.
func (c *Client) EditInlineMessageLiveLocation(ctx context.Context, request *EditInlineMessageLiveLocationRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
