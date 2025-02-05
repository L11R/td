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

// SuggestedActionEnableArchiveAndMuteNewChats represents TL type `suggestedActionEnableArchiveAndMuteNewChats#7841ec4f`.
type SuggestedActionEnableArchiveAndMuteNewChats struct {
}

// SuggestedActionEnableArchiveAndMuteNewChatsTypeID is TL type id of SuggestedActionEnableArchiveAndMuteNewChats.
const SuggestedActionEnableArchiveAndMuteNewChatsTypeID = 0x7841ec4f

// construct implements constructor of SuggestedActionClass.
func (s SuggestedActionEnableArchiveAndMuteNewChats) construct() SuggestedActionClass { return &s }

// Ensuring interfaces in compile-time for SuggestedActionEnableArchiveAndMuteNewChats.
var (
	_ bin.Encoder     = &SuggestedActionEnableArchiveAndMuteNewChats{}
	_ bin.Decoder     = &SuggestedActionEnableArchiveAndMuteNewChats{}
	_ bin.BareEncoder = &SuggestedActionEnableArchiveAndMuteNewChats{}
	_ bin.BareDecoder = &SuggestedActionEnableArchiveAndMuteNewChats{}

	_ SuggestedActionClass = &SuggestedActionEnableArchiveAndMuteNewChats{}
)

func (s *SuggestedActionEnableArchiveAndMuteNewChats) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *SuggestedActionEnableArchiveAndMuteNewChats) String() string {
	if s == nil {
		return "SuggestedActionEnableArchiveAndMuteNewChats(nil)"
	}
	type Alias SuggestedActionEnableArchiveAndMuteNewChats
	return fmt.Sprintf("SuggestedActionEnableArchiveAndMuteNewChats%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SuggestedActionEnableArchiveAndMuteNewChats) TypeID() uint32 {
	return SuggestedActionEnableArchiveAndMuteNewChatsTypeID
}

// TypeName returns name of type in TL schema.
func (*SuggestedActionEnableArchiveAndMuteNewChats) TypeName() string {
	return "suggestedActionEnableArchiveAndMuteNewChats"
}

// TypeInfo returns info about TL type.
func (s *SuggestedActionEnableArchiveAndMuteNewChats) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "suggestedActionEnableArchiveAndMuteNewChats",
		ID:   SuggestedActionEnableArchiveAndMuteNewChatsTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (s *SuggestedActionEnableArchiveAndMuteNewChats) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode suggestedActionEnableArchiveAndMuteNewChats#7841ec4f as nil")
	}
	b.PutID(SuggestedActionEnableArchiveAndMuteNewChatsTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SuggestedActionEnableArchiveAndMuteNewChats) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode suggestedActionEnableArchiveAndMuteNewChats#7841ec4f as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *SuggestedActionEnableArchiveAndMuteNewChats) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode suggestedActionEnableArchiveAndMuteNewChats#7841ec4f to nil")
	}
	if err := b.ConsumeID(SuggestedActionEnableArchiveAndMuteNewChatsTypeID); err != nil {
		return fmt.Errorf("unable to decode suggestedActionEnableArchiveAndMuteNewChats#7841ec4f: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SuggestedActionEnableArchiveAndMuteNewChats) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode suggestedActionEnableArchiveAndMuteNewChats#7841ec4f to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SuggestedActionEnableArchiveAndMuteNewChats) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode suggestedActionEnableArchiveAndMuteNewChats#7841ec4f as nil")
	}
	b.ObjStart()
	b.PutID("suggestedActionEnableArchiveAndMuteNewChats")
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SuggestedActionEnableArchiveAndMuteNewChats) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode suggestedActionEnableArchiveAndMuteNewChats#7841ec4f to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("suggestedActionEnableArchiveAndMuteNewChats"); err != nil {
				return fmt.Errorf("unable to decode suggestedActionEnableArchiveAndMuteNewChats#7841ec4f: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// SuggestedActionCheckPassword represents TL type `suggestedActionCheckPassword#71e072b7`.
type SuggestedActionCheckPassword struct {
}

// SuggestedActionCheckPasswordTypeID is TL type id of SuggestedActionCheckPassword.
const SuggestedActionCheckPasswordTypeID = 0x71e072b7

// construct implements constructor of SuggestedActionClass.
func (s SuggestedActionCheckPassword) construct() SuggestedActionClass { return &s }

// Ensuring interfaces in compile-time for SuggestedActionCheckPassword.
var (
	_ bin.Encoder     = &SuggestedActionCheckPassword{}
	_ bin.Decoder     = &SuggestedActionCheckPassword{}
	_ bin.BareEncoder = &SuggestedActionCheckPassword{}
	_ bin.BareDecoder = &SuggestedActionCheckPassword{}

	_ SuggestedActionClass = &SuggestedActionCheckPassword{}
)

func (s *SuggestedActionCheckPassword) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *SuggestedActionCheckPassword) String() string {
	if s == nil {
		return "SuggestedActionCheckPassword(nil)"
	}
	type Alias SuggestedActionCheckPassword
	return fmt.Sprintf("SuggestedActionCheckPassword%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SuggestedActionCheckPassword) TypeID() uint32 {
	return SuggestedActionCheckPasswordTypeID
}

// TypeName returns name of type in TL schema.
func (*SuggestedActionCheckPassword) TypeName() string {
	return "suggestedActionCheckPassword"
}

// TypeInfo returns info about TL type.
func (s *SuggestedActionCheckPassword) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "suggestedActionCheckPassword",
		ID:   SuggestedActionCheckPasswordTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (s *SuggestedActionCheckPassword) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode suggestedActionCheckPassword#71e072b7 as nil")
	}
	b.PutID(SuggestedActionCheckPasswordTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SuggestedActionCheckPassword) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode suggestedActionCheckPassword#71e072b7 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *SuggestedActionCheckPassword) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode suggestedActionCheckPassword#71e072b7 to nil")
	}
	if err := b.ConsumeID(SuggestedActionCheckPasswordTypeID); err != nil {
		return fmt.Errorf("unable to decode suggestedActionCheckPassword#71e072b7: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SuggestedActionCheckPassword) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode suggestedActionCheckPassword#71e072b7 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SuggestedActionCheckPassword) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode suggestedActionCheckPassword#71e072b7 as nil")
	}
	b.ObjStart()
	b.PutID("suggestedActionCheckPassword")
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SuggestedActionCheckPassword) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode suggestedActionCheckPassword#71e072b7 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("suggestedActionCheckPassword"); err != nil {
				return fmt.Errorf("unable to decode suggestedActionCheckPassword#71e072b7: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// SuggestedActionCheckPhoneNumber represents TL type `suggestedActionCheckPhoneNumber#26ab77eb`.
type SuggestedActionCheckPhoneNumber struct {
}

// SuggestedActionCheckPhoneNumberTypeID is TL type id of SuggestedActionCheckPhoneNumber.
const SuggestedActionCheckPhoneNumberTypeID = 0x26ab77eb

// construct implements constructor of SuggestedActionClass.
func (s SuggestedActionCheckPhoneNumber) construct() SuggestedActionClass { return &s }

// Ensuring interfaces in compile-time for SuggestedActionCheckPhoneNumber.
var (
	_ bin.Encoder     = &SuggestedActionCheckPhoneNumber{}
	_ bin.Decoder     = &SuggestedActionCheckPhoneNumber{}
	_ bin.BareEncoder = &SuggestedActionCheckPhoneNumber{}
	_ bin.BareDecoder = &SuggestedActionCheckPhoneNumber{}

	_ SuggestedActionClass = &SuggestedActionCheckPhoneNumber{}
)

func (s *SuggestedActionCheckPhoneNumber) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *SuggestedActionCheckPhoneNumber) String() string {
	if s == nil {
		return "SuggestedActionCheckPhoneNumber(nil)"
	}
	type Alias SuggestedActionCheckPhoneNumber
	return fmt.Sprintf("SuggestedActionCheckPhoneNumber%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SuggestedActionCheckPhoneNumber) TypeID() uint32 {
	return SuggestedActionCheckPhoneNumberTypeID
}

// TypeName returns name of type in TL schema.
func (*SuggestedActionCheckPhoneNumber) TypeName() string {
	return "suggestedActionCheckPhoneNumber"
}

// TypeInfo returns info about TL type.
func (s *SuggestedActionCheckPhoneNumber) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "suggestedActionCheckPhoneNumber",
		ID:   SuggestedActionCheckPhoneNumberTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (s *SuggestedActionCheckPhoneNumber) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode suggestedActionCheckPhoneNumber#26ab77eb as nil")
	}
	b.PutID(SuggestedActionCheckPhoneNumberTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SuggestedActionCheckPhoneNumber) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode suggestedActionCheckPhoneNumber#26ab77eb as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *SuggestedActionCheckPhoneNumber) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode suggestedActionCheckPhoneNumber#26ab77eb to nil")
	}
	if err := b.ConsumeID(SuggestedActionCheckPhoneNumberTypeID); err != nil {
		return fmt.Errorf("unable to decode suggestedActionCheckPhoneNumber#26ab77eb: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SuggestedActionCheckPhoneNumber) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode suggestedActionCheckPhoneNumber#26ab77eb to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SuggestedActionCheckPhoneNumber) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode suggestedActionCheckPhoneNumber#26ab77eb as nil")
	}
	b.ObjStart()
	b.PutID("suggestedActionCheckPhoneNumber")
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SuggestedActionCheckPhoneNumber) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode suggestedActionCheckPhoneNumber#26ab77eb to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("suggestedActionCheckPhoneNumber"); err != nil {
				return fmt.Errorf("unable to decode suggestedActionCheckPhoneNumber#26ab77eb: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// SuggestedActionSeeTicksHint represents TL type `suggestedActionSeeTicksHint#3f4ae062`.
type SuggestedActionSeeTicksHint struct {
}

// SuggestedActionSeeTicksHintTypeID is TL type id of SuggestedActionSeeTicksHint.
const SuggestedActionSeeTicksHintTypeID = 0x3f4ae062

// construct implements constructor of SuggestedActionClass.
func (s SuggestedActionSeeTicksHint) construct() SuggestedActionClass { return &s }

// Ensuring interfaces in compile-time for SuggestedActionSeeTicksHint.
var (
	_ bin.Encoder     = &SuggestedActionSeeTicksHint{}
	_ bin.Decoder     = &SuggestedActionSeeTicksHint{}
	_ bin.BareEncoder = &SuggestedActionSeeTicksHint{}
	_ bin.BareDecoder = &SuggestedActionSeeTicksHint{}

	_ SuggestedActionClass = &SuggestedActionSeeTicksHint{}
)

func (s *SuggestedActionSeeTicksHint) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *SuggestedActionSeeTicksHint) String() string {
	if s == nil {
		return "SuggestedActionSeeTicksHint(nil)"
	}
	type Alias SuggestedActionSeeTicksHint
	return fmt.Sprintf("SuggestedActionSeeTicksHint%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SuggestedActionSeeTicksHint) TypeID() uint32 {
	return SuggestedActionSeeTicksHintTypeID
}

// TypeName returns name of type in TL schema.
func (*SuggestedActionSeeTicksHint) TypeName() string {
	return "suggestedActionSeeTicksHint"
}

// TypeInfo returns info about TL type.
func (s *SuggestedActionSeeTicksHint) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "suggestedActionSeeTicksHint",
		ID:   SuggestedActionSeeTicksHintTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (s *SuggestedActionSeeTicksHint) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode suggestedActionSeeTicksHint#3f4ae062 as nil")
	}
	b.PutID(SuggestedActionSeeTicksHintTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SuggestedActionSeeTicksHint) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode suggestedActionSeeTicksHint#3f4ae062 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *SuggestedActionSeeTicksHint) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode suggestedActionSeeTicksHint#3f4ae062 to nil")
	}
	if err := b.ConsumeID(SuggestedActionSeeTicksHintTypeID); err != nil {
		return fmt.Errorf("unable to decode suggestedActionSeeTicksHint#3f4ae062: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SuggestedActionSeeTicksHint) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode suggestedActionSeeTicksHint#3f4ae062 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SuggestedActionSeeTicksHint) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode suggestedActionSeeTicksHint#3f4ae062 as nil")
	}
	b.ObjStart()
	b.PutID("suggestedActionSeeTicksHint")
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SuggestedActionSeeTicksHint) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode suggestedActionSeeTicksHint#3f4ae062 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("suggestedActionSeeTicksHint"); err != nil {
				return fmt.Errorf("unable to decode suggestedActionSeeTicksHint#3f4ae062: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// SuggestedActionConvertToBroadcastGroup represents TL type `suggestedActionConvertToBroadcastGroup#c67a2e38`.
type SuggestedActionConvertToBroadcastGroup struct {
	// Supergroup identifier
	SupergroupID int64
}

// SuggestedActionConvertToBroadcastGroupTypeID is TL type id of SuggestedActionConvertToBroadcastGroup.
const SuggestedActionConvertToBroadcastGroupTypeID = 0xc67a2e38

// construct implements constructor of SuggestedActionClass.
func (s SuggestedActionConvertToBroadcastGroup) construct() SuggestedActionClass { return &s }

// Ensuring interfaces in compile-time for SuggestedActionConvertToBroadcastGroup.
var (
	_ bin.Encoder     = &SuggestedActionConvertToBroadcastGroup{}
	_ bin.Decoder     = &SuggestedActionConvertToBroadcastGroup{}
	_ bin.BareEncoder = &SuggestedActionConvertToBroadcastGroup{}
	_ bin.BareDecoder = &SuggestedActionConvertToBroadcastGroup{}

	_ SuggestedActionClass = &SuggestedActionConvertToBroadcastGroup{}
)

func (s *SuggestedActionConvertToBroadcastGroup) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.SupergroupID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SuggestedActionConvertToBroadcastGroup) String() string {
	if s == nil {
		return "SuggestedActionConvertToBroadcastGroup(nil)"
	}
	type Alias SuggestedActionConvertToBroadcastGroup
	return fmt.Sprintf("SuggestedActionConvertToBroadcastGroup%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SuggestedActionConvertToBroadcastGroup) TypeID() uint32 {
	return SuggestedActionConvertToBroadcastGroupTypeID
}

// TypeName returns name of type in TL schema.
func (*SuggestedActionConvertToBroadcastGroup) TypeName() string {
	return "suggestedActionConvertToBroadcastGroup"
}

// TypeInfo returns info about TL type.
func (s *SuggestedActionConvertToBroadcastGroup) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "suggestedActionConvertToBroadcastGroup",
		ID:   SuggestedActionConvertToBroadcastGroupTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "SupergroupID",
			SchemaName: "supergroup_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SuggestedActionConvertToBroadcastGroup) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode suggestedActionConvertToBroadcastGroup#c67a2e38 as nil")
	}
	b.PutID(SuggestedActionConvertToBroadcastGroupTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SuggestedActionConvertToBroadcastGroup) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode suggestedActionConvertToBroadcastGroup#c67a2e38 as nil")
	}
	b.PutInt53(s.SupergroupID)
	return nil
}

// Decode implements bin.Decoder.
func (s *SuggestedActionConvertToBroadcastGroup) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode suggestedActionConvertToBroadcastGroup#c67a2e38 to nil")
	}
	if err := b.ConsumeID(SuggestedActionConvertToBroadcastGroupTypeID); err != nil {
		return fmt.Errorf("unable to decode suggestedActionConvertToBroadcastGroup#c67a2e38: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SuggestedActionConvertToBroadcastGroup) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode suggestedActionConvertToBroadcastGroup#c67a2e38 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode suggestedActionConvertToBroadcastGroup#c67a2e38: field supergroup_id: %w", err)
		}
		s.SupergroupID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SuggestedActionConvertToBroadcastGroup) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode suggestedActionConvertToBroadcastGroup#c67a2e38 as nil")
	}
	b.ObjStart()
	b.PutID("suggestedActionConvertToBroadcastGroup")
	b.FieldStart("supergroup_id")
	b.PutInt53(s.SupergroupID)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SuggestedActionConvertToBroadcastGroup) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode suggestedActionConvertToBroadcastGroup#c67a2e38 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("suggestedActionConvertToBroadcastGroup"); err != nil {
				return fmt.Errorf("unable to decode suggestedActionConvertToBroadcastGroup#c67a2e38: %w", err)
			}
		case "supergroup_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode suggestedActionConvertToBroadcastGroup#c67a2e38: field supergroup_id: %w", err)
			}
			s.SupergroupID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetSupergroupID returns value of SupergroupID field.
func (s *SuggestedActionConvertToBroadcastGroup) GetSupergroupID() (value int64) {
	return s.SupergroupID
}

// SuggestedActionClassName is schema name of SuggestedActionClass.
const SuggestedActionClassName = "SuggestedAction"

// SuggestedActionClass represents SuggestedAction generic type.
//
// Example:
//  g, err := tdapi.DecodeSuggestedAction(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tdapi.SuggestedActionEnableArchiveAndMuteNewChats: // suggestedActionEnableArchiveAndMuteNewChats#7841ec4f
//  case *tdapi.SuggestedActionCheckPassword: // suggestedActionCheckPassword#71e072b7
//  case *tdapi.SuggestedActionCheckPhoneNumber: // suggestedActionCheckPhoneNumber#26ab77eb
//  case *tdapi.SuggestedActionSeeTicksHint: // suggestedActionSeeTicksHint#3f4ae062
//  case *tdapi.SuggestedActionConvertToBroadcastGroup: // suggestedActionConvertToBroadcastGroup#c67a2e38
//  default: panic(v)
//  }
type SuggestedActionClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() SuggestedActionClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	EncodeTDLibJSON(b tdjson.Encoder) error
	DecodeTDLibJSON(b tdjson.Decoder) error
}

// DecodeSuggestedAction implements binary de-serialization for SuggestedActionClass.
func DecodeSuggestedAction(buf *bin.Buffer) (SuggestedActionClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case SuggestedActionEnableArchiveAndMuteNewChatsTypeID:
		// Decoding suggestedActionEnableArchiveAndMuteNewChats#7841ec4f.
		v := SuggestedActionEnableArchiveAndMuteNewChats{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SuggestedActionClass: %w", err)
		}
		return &v, nil
	case SuggestedActionCheckPasswordTypeID:
		// Decoding suggestedActionCheckPassword#71e072b7.
		v := SuggestedActionCheckPassword{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SuggestedActionClass: %w", err)
		}
		return &v, nil
	case SuggestedActionCheckPhoneNumberTypeID:
		// Decoding suggestedActionCheckPhoneNumber#26ab77eb.
		v := SuggestedActionCheckPhoneNumber{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SuggestedActionClass: %w", err)
		}
		return &v, nil
	case SuggestedActionSeeTicksHintTypeID:
		// Decoding suggestedActionSeeTicksHint#3f4ae062.
		v := SuggestedActionSeeTicksHint{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SuggestedActionClass: %w", err)
		}
		return &v, nil
	case SuggestedActionConvertToBroadcastGroupTypeID:
		// Decoding suggestedActionConvertToBroadcastGroup#c67a2e38.
		v := SuggestedActionConvertToBroadcastGroup{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SuggestedActionClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode SuggestedActionClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONSuggestedAction implements binary de-serialization for SuggestedActionClass.
func DecodeTDLibJSONSuggestedAction(buf tdjson.Decoder) (SuggestedActionClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "suggestedActionEnableArchiveAndMuteNewChats":
		// Decoding suggestedActionEnableArchiveAndMuteNewChats#7841ec4f.
		v := SuggestedActionEnableArchiveAndMuteNewChats{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SuggestedActionClass: %w", err)
		}
		return &v, nil
	case "suggestedActionCheckPassword":
		// Decoding suggestedActionCheckPassword#71e072b7.
		v := SuggestedActionCheckPassword{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SuggestedActionClass: %w", err)
		}
		return &v, nil
	case "suggestedActionCheckPhoneNumber":
		// Decoding suggestedActionCheckPhoneNumber#26ab77eb.
		v := SuggestedActionCheckPhoneNumber{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SuggestedActionClass: %w", err)
		}
		return &v, nil
	case "suggestedActionSeeTicksHint":
		// Decoding suggestedActionSeeTicksHint#3f4ae062.
		v := SuggestedActionSeeTicksHint{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SuggestedActionClass: %w", err)
		}
		return &v, nil
	case "suggestedActionConvertToBroadcastGroup":
		// Decoding suggestedActionConvertToBroadcastGroup#c67a2e38.
		v := SuggestedActionConvertToBroadcastGroup{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SuggestedActionClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode SuggestedActionClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// SuggestedAction boxes the SuggestedActionClass providing a helper.
type SuggestedActionBox struct {
	SuggestedAction SuggestedActionClass
}

// Decode implements bin.Decoder for SuggestedActionBox.
func (b *SuggestedActionBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode SuggestedActionBox to nil")
	}
	v, err := DecodeSuggestedAction(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.SuggestedAction = v
	return nil
}

// Encode implements bin.Encode for SuggestedActionBox.
func (b *SuggestedActionBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.SuggestedAction == nil {
		return fmt.Errorf("unable to encode SuggestedActionClass as nil")
	}
	return b.SuggestedAction.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for SuggestedActionBox.
func (b *SuggestedActionBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode SuggestedActionBox to nil")
	}
	v, err := DecodeTDLibJSONSuggestedAction(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.SuggestedAction = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for SuggestedActionBox.
func (b *SuggestedActionBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.SuggestedAction == nil {
		return fmt.Errorf("unable to encode SuggestedActionClass as nil")
	}
	return b.SuggestedAction.EncodeTDLibJSON(buf)
}
