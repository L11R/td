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

// ChatPhotoInfo represents TL type `chatPhotoInfo#9f51bb6`.
type ChatPhotoInfo struct {
	// A small (160x160) chat photo variant in JPEG format. The file can be downloaded only
	// before the photo is changed
	Small File
	// A big (640x640) chat photo variant in JPEG format. The file can be downloaded only
	// before the photo is changed
	Big File
	// Chat photo minithumbnail; may be null
	Minithumbnail Minithumbnail
	// True, if the photo has animated variant
	HasAnimation bool
}

// ChatPhotoInfoTypeID is TL type id of ChatPhotoInfo.
const ChatPhotoInfoTypeID = 0x9f51bb6

// Ensuring interfaces in compile-time for ChatPhotoInfo.
var (
	_ bin.Encoder     = &ChatPhotoInfo{}
	_ bin.Decoder     = &ChatPhotoInfo{}
	_ bin.BareEncoder = &ChatPhotoInfo{}
	_ bin.BareDecoder = &ChatPhotoInfo{}
)

func (c *ChatPhotoInfo) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Small.Zero()) {
		return false
	}
	if !(c.Big.Zero()) {
		return false
	}
	if !(c.Minithumbnail.Zero()) {
		return false
	}
	if !(c.HasAnimation == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatPhotoInfo) String() string {
	if c == nil {
		return "ChatPhotoInfo(nil)"
	}
	type Alias ChatPhotoInfo
	return fmt.Sprintf("ChatPhotoInfo%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatPhotoInfo) TypeID() uint32 {
	return ChatPhotoInfoTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatPhotoInfo) TypeName() string {
	return "chatPhotoInfo"
}

// TypeInfo returns info about TL type.
func (c *ChatPhotoInfo) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatPhotoInfo",
		ID:   ChatPhotoInfoTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Small",
			SchemaName: "small",
		},
		{
			Name:       "Big",
			SchemaName: "big",
		},
		{
			Name:       "Minithumbnail",
			SchemaName: "minithumbnail",
		},
		{
			Name:       "HasAnimation",
			SchemaName: "has_animation",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatPhotoInfo) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatPhotoInfo#9f51bb6 as nil")
	}
	b.PutID(ChatPhotoInfoTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatPhotoInfo) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatPhotoInfo#9f51bb6 as nil")
	}
	if err := c.Small.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chatPhotoInfo#9f51bb6: field small: %w", err)
	}
	if err := c.Big.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chatPhotoInfo#9f51bb6: field big: %w", err)
	}
	if err := c.Minithumbnail.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chatPhotoInfo#9f51bb6: field minithumbnail: %w", err)
	}
	b.PutBool(c.HasAnimation)
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatPhotoInfo) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatPhotoInfo#9f51bb6 to nil")
	}
	if err := b.ConsumeID(ChatPhotoInfoTypeID); err != nil {
		return fmt.Errorf("unable to decode chatPhotoInfo#9f51bb6: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatPhotoInfo) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatPhotoInfo#9f51bb6 to nil")
	}
	{
		if err := c.Small.Decode(b); err != nil {
			return fmt.Errorf("unable to decode chatPhotoInfo#9f51bb6: field small: %w", err)
		}
	}
	{
		if err := c.Big.Decode(b); err != nil {
			return fmt.Errorf("unable to decode chatPhotoInfo#9f51bb6: field big: %w", err)
		}
	}
	{
		if err := c.Minithumbnail.Decode(b); err != nil {
			return fmt.Errorf("unable to decode chatPhotoInfo#9f51bb6: field minithumbnail: %w", err)
		}
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatPhotoInfo#9f51bb6: field has_animation: %w", err)
		}
		c.HasAnimation = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *ChatPhotoInfo) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode chatPhotoInfo#9f51bb6 as nil")
	}
	b.ObjStart()
	b.PutID("chatPhotoInfo")
	b.FieldStart("small")
	if err := c.Small.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode chatPhotoInfo#9f51bb6: field small: %w", err)
	}
	b.FieldStart("big")
	if err := c.Big.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode chatPhotoInfo#9f51bb6: field big: %w", err)
	}
	b.FieldStart("minithumbnail")
	if err := c.Minithumbnail.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode chatPhotoInfo#9f51bb6: field minithumbnail: %w", err)
	}
	b.FieldStart("has_animation")
	b.PutBool(c.HasAnimation)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *ChatPhotoInfo) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode chatPhotoInfo#9f51bb6 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("chatPhotoInfo"); err != nil {
				return fmt.Errorf("unable to decode chatPhotoInfo#9f51bb6: %w", err)
			}
		case "small":
			if err := c.Small.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode chatPhotoInfo#9f51bb6: field small: %w", err)
			}
		case "big":
			if err := c.Big.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode chatPhotoInfo#9f51bb6: field big: %w", err)
			}
		case "minithumbnail":
			if err := c.Minithumbnail.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode chatPhotoInfo#9f51bb6: field minithumbnail: %w", err)
			}
		case "has_animation":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatPhotoInfo#9f51bb6: field has_animation: %w", err)
			}
			c.HasAnimation = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetSmall returns value of Small field.
func (c *ChatPhotoInfo) GetSmall() (value File) {
	return c.Small
}

// GetBig returns value of Big field.
func (c *ChatPhotoInfo) GetBig() (value File) {
	return c.Big
}

// GetMinithumbnail returns value of Minithumbnail field.
func (c *ChatPhotoInfo) GetMinithumbnail() (value Minithumbnail) {
	return c.Minithumbnail
}

// GetHasAnimation returns value of HasAnimation field.
func (c *ChatPhotoInfo) GetHasAnimation() (value bool) {
	return c.HasAnimation
}
