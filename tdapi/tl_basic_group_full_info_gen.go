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

// BasicGroupFullInfo represents TL type `basicGroupFullInfo#f191cfca`.
type BasicGroupFullInfo struct {
	// Chat photo; may be null
	Photo ChatPhoto
	// Contains full information about a basic group
	Description string
	// User identifier of the creator of the group; 0 if unknown
	CreatorUserID int64
	// Group members
	Members []ChatMember
	// Primary invite link for this group; may be null. For chat administrators with
	// can_invite_users right only. Updated only after the basic group is opened
	InviteLink ChatInviteLink
	// List of commands of bots in the group
	BotCommands []BotCommands
}

// BasicGroupFullInfoTypeID is TL type id of BasicGroupFullInfo.
const BasicGroupFullInfoTypeID = 0xf191cfca

// Ensuring interfaces in compile-time for BasicGroupFullInfo.
var (
	_ bin.Encoder     = &BasicGroupFullInfo{}
	_ bin.Decoder     = &BasicGroupFullInfo{}
	_ bin.BareEncoder = &BasicGroupFullInfo{}
	_ bin.BareDecoder = &BasicGroupFullInfo{}
)

func (b *BasicGroupFullInfo) Zero() bool {
	if b == nil {
		return true
	}
	if !(b.Photo.Zero()) {
		return false
	}
	if !(b.Description == "") {
		return false
	}
	if !(b.CreatorUserID == 0) {
		return false
	}
	if !(b.Members == nil) {
		return false
	}
	if !(b.InviteLink.Zero()) {
		return false
	}
	if !(b.BotCommands == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (b *BasicGroupFullInfo) String() string {
	if b == nil {
		return "BasicGroupFullInfo(nil)"
	}
	type Alias BasicGroupFullInfo
	return fmt.Sprintf("BasicGroupFullInfo%+v", Alias(*b))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BasicGroupFullInfo) TypeID() uint32 {
	return BasicGroupFullInfoTypeID
}

// TypeName returns name of type in TL schema.
func (*BasicGroupFullInfo) TypeName() string {
	return "basicGroupFullInfo"
}

// TypeInfo returns info about TL type.
func (b *BasicGroupFullInfo) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "basicGroupFullInfo",
		ID:   BasicGroupFullInfoTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Photo",
			SchemaName: "photo",
		},
		{
			Name:       "Description",
			SchemaName: "description",
		},
		{
			Name:       "CreatorUserID",
			SchemaName: "creator_user_id",
		},
		{
			Name:       "Members",
			SchemaName: "members",
		},
		{
			Name:       "InviteLink",
			SchemaName: "invite_link",
		},
		{
			Name:       "BotCommands",
			SchemaName: "bot_commands",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (b *BasicGroupFullInfo) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode basicGroupFullInfo#f191cfca as nil")
	}
	buf.PutID(BasicGroupFullInfoTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BasicGroupFullInfo) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode basicGroupFullInfo#f191cfca as nil")
	}
	if err := b.Photo.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode basicGroupFullInfo#f191cfca: field photo: %w", err)
	}
	buf.PutString(b.Description)
	buf.PutInt53(b.CreatorUserID)
	buf.PutInt(len(b.Members))
	for idx, v := range b.Members {
		if err := v.EncodeBare(buf); err != nil {
			return fmt.Errorf("unable to encode bare basicGroupFullInfo#f191cfca: field members element with index %d: %w", idx, err)
		}
	}
	if err := b.InviteLink.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode basicGroupFullInfo#f191cfca: field invite_link: %w", err)
	}
	buf.PutInt(len(b.BotCommands))
	for idx, v := range b.BotCommands {
		if err := v.EncodeBare(buf); err != nil {
			return fmt.Errorf("unable to encode bare basicGroupFullInfo#f191cfca: field bot_commands element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (b *BasicGroupFullInfo) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode basicGroupFullInfo#f191cfca to nil")
	}
	if err := buf.ConsumeID(BasicGroupFullInfoTypeID); err != nil {
		return fmt.Errorf("unable to decode basicGroupFullInfo#f191cfca: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BasicGroupFullInfo) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode basicGroupFullInfo#f191cfca to nil")
	}
	{
		if err := b.Photo.Decode(buf); err != nil {
			return fmt.Errorf("unable to decode basicGroupFullInfo#f191cfca: field photo: %w", err)
		}
	}
	{
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode basicGroupFullInfo#f191cfca: field description: %w", err)
		}
		b.Description = value
	}
	{
		value, err := buf.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode basicGroupFullInfo#f191cfca: field creator_user_id: %w", err)
		}
		b.CreatorUserID = value
	}
	{
		headerLen, err := buf.Int()
		if err != nil {
			return fmt.Errorf("unable to decode basicGroupFullInfo#f191cfca: field members: %w", err)
		}

		if headerLen > 0 {
			b.Members = make([]ChatMember, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value ChatMember
			if err := value.DecodeBare(buf); err != nil {
				return fmt.Errorf("unable to decode bare basicGroupFullInfo#f191cfca: field members: %w", err)
			}
			b.Members = append(b.Members, value)
		}
	}
	{
		if err := b.InviteLink.Decode(buf); err != nil {
			return fmt.Errorf("unable to decode basicGroupFullInfo#f191cfca: field invite_link: %w", err)
		}
	}
	{
		headerLen, err := buf.Int()
		if err != nil {
			return fmt.Errorf("unable to decode basicGroupFullInfo#f191cfca: field bot_commands: %w", err)
		}

		if headerLen > 0 {
			b.BotCommands = make([]BotCommands, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value BotCommands
			if err := value.DecodeBare(buf); err != nil {
				return fmt.Errorf("unable to decode bare basicGroupFullInfo#f191cfca: field bot_commands: %w", err)
			}
			b.BotCommands = append(b.BotCommands, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (b *BasicGroupFullInfo) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil {
		return fmt.Errorf("can't encode basicGroupFullInfo#f191cfca as nil")
	}
	buf.ObjStart()
	buf.PutID("basicGroupFullInfo")
	buf.FieldStart("photo")
	if err := b.Photo.EncodeTDLibJSON(buf); err != nil {
		return fmt.Errorf("unable to encode basicGroupFullInfo#f191cfca: field photo: %w", err)
	}
	buf.FieldStart("description")
	buf.PutString(b.Description)
	buf.FieldStart("creator_user_id")
	buf.PutInt53(b.CreatorUserID)
	buf.FieldStart("members")
	buf.ArrStart()
	for idx, v := range b.Members {
		if err := v.EncodeTDLibJSON(buf); err != nil {
			return fmt.Errorf("unable to encode basicGroupFullInfo#f191cfca: field members element with index %d: %w", idx, err)
		}
	}
	buf.ArrEnd()
	buf.FieldStart("invite_link")
	if err := b.InviteLink.EncodeTDLibJSON(buf); err != nil {
		return fmt.Errorf("unable to encode basicGroupFullInfo#f191cfca: field invite_link: %w", err)
	}
	buf.FieldStart("bot_commands")
	buf.ArrStart()
	for idx, v := range b.BotCommands {
		if err := v.EncodeTDLibJSON(buf); err != nil {
			return fmt.Errorf("unable to encode basicGroupFullInfo#f191cfca: field bot_commands element with index %d: %w", idx, err)
		}
	}
	buf.ArrEnd()
	buf.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (b *BasicGroupFullInfo) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("can't decode basicGroupFullInfo#f191cfca to nil")
	}

	return buf.Obj(func(buf tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := buf.ConsumeID("basicGroupFullInfo"); err != nil {
				return fmt.Errorf("unable to decode basicGroupFullInfo#f191cfca: %w", err)
			}
		case "photo":
			if err := b.Photo.DecodeTDLibJSON(buf); err != nil {
				return fmt.Errorf("unable to decode basicGroupFullInfo#f191cfca: field photo: %w", err)
			}
		case "description":
			value, err := buf.String()
			if err != nil {
				return fmt.Errorf("unable to decode basicGroupFullInfo#f191cfca: field description: %w", err)
			}
			b.Description = value
		case "creator_user_id":
			value, err := buf.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode basicGroupFullInfo#f191cfca: field creator_user_id: %w", err)
			}
			b.CreatorUserID = value
		case "members":
			if err := buf.Arr(func(buf tdjson.Decoder) error {
				var value ChatMember
				if err := value.DecodeTDLibJSON(buf); err != nil {
					return fmt.Errorf("unable to decode basicGroupFullInfo#f191cfca: field members: %w", err)
				}
				b.Members = append(b.Members, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode basicGroupFullInfo#f191cfca: field members: %w", err)
			}
		case "invite_link":
			if err := b.InviteLink.DecodeTDLibJSON(buf); err != nil {
				return fmt.Errorf("unable to decode basicGroupFullInfo#f191cfca: field invite_link: %w", err)
			}
		case "bot_commands":
			if err := buf.Arr(func(buf tdjson.Decoder) error {
				var value BotCommands
				if err := value.DecodeTDLibJSON(buf); err != nil {
					return fmt.Errorf("unable to decode basicGroupFullInfo#f191cfca: field bot_commands: %w", err)
				}
				b.BotCommands = append(b.BotCommands, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode basicGroupFullInfo#f191cfca: field bot_commands: %w", err)
			}
		default:
			return buf.Skip()
		}
		return nil
	})
}

// GetPhoto returns value of Photo field.
func (b *BasicGroupFullInfo) GetPhoto() (value ChatPhoto) {
	return b.Photo
}

// GetDescription returns value of Description field.
func (b *BasicGroupFullInfo) GetDescription() (value string) {
	return b.Description
}

// GetCreatorUserID returns value of CreatorUserID field.
func (b *BasicGroupFullInfo) GetCreatorUserID() (value int64) {
	return b.CreatorUserID
}

// GetMembers returns value of Members field.
func (b *BasicGroupFullInfo) GetMembers() (value []ChatMember) {
	return b.Members
}

// GetInviteLink returns value of InviteLink field.
func (b *BasicGroupFullInfo) GetInviteLink() (value ChatInviteLink) {
	return b.InviteLink
}

// GetBotCommands returns value of BotCommands field.
func (b *BasicGroupFullInfo) GetBotCommands() (value []BotCommands) {
	return b.BotCommands
}
