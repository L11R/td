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

// Sessions represents TL type `sessions#1da11c6e`.
type Sessions struct {
	// List of sessions
	Sessions []Session
}

// SessionsTypeID is TL type id of Sessions.
const SessionsTypeID = 0x1da11c6e

// Ensuring interfaces in compile-time for Sessions.
var (
	_ bin.Encoder     = &Sessions{}
	_ bin.Decoder     = &Sessions{}
	_ bin.BareEncoder = &Sessions{}
	_ bin.BareDecoder = &Sessions{}
)

func (s *Sessions) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Sessions == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *Sessions) String() string {
	if s == nil {
		return "Sessions(nil)"
	}
	type Alias Sessions
	return fmt.Sprintf("Sessions%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Sessions) TypeID() uint32 {
	return SessionsTypeID
}

// TypeName returns name of type in TL schema.
func (*Sessions) TypeName() string {
	return "sessions"
}

// TypeInfo returns info about TL type.
func (s *Sessions) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "sessions",
		ID:   SessionsTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Sessions",
			SchemaName: "sessions",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *Sessions) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sessions#1da11c6e as nil")
	}
	b.PutID(SessionsTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *Sessions) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sessions#1da11c6e as nil")
	}
	b.PutInt(len(s.Sessions))
	for idx, v := range s.Sessions {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare sessions#1da11c6e: field sessions element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *Sessions) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sessions#1da11c6e to nil")
	}
	if err := b.ConsumeID(SessionsTypeID); err != nil {
		return fmt.Errorf("unable to decode sessions#1da11c6e: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *Sessions) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sessions#1da11c6e to nil")
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode sessions#1da11c6e: field sessions: %w", err)
		}

		if headerLen > 0 {
			s.Sessions = make([]Session, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value Session
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare sessions#1da11c6e: field sessions: %w", err)
			}
			s.Sessions = append(s.Sessions, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *Sessions) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode sessions#1da11c6e as nil")
	}
	b.ObjStart()
	b.PutID("sessions")
	b.FieldStart("sessions")
	b.ArrStart()
	for idx, v := range s.Sessions {
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode sessions#1da11c6e: field sessions element with index %d: %w", idx, err)
		}
	}
	b.ArrEnd()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *Sessions) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode sessions#1da11c6e to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("sessions"); err != nil {
				return fmt.Errorf("unable to decode sessions#1da11c6e: %w", err)
			}
		case "sessions":
			if err := b.Arr(func(b tdjson.Decoder) error {
				var value Session
				if err := value.DecodeTDLibJSON(b); err != nil {
					return fmt.Errorf("unable to decode sessions#1da11c6e: field sessions: %w", err)
				}
				s.Sessions = append(s.Sessions, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode sessions#1da11c6e: field sessions: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetSessions returns value of Sessions field.
func (s *Sessions) GetSessions() (value []Session) {
	return s.Sessions
}
