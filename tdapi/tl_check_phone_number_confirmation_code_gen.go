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

// CheckPhoneNumberConfirmationCodeRequest represents TL type `checkPhoneNumberConfirmationCode#afa638da`.
type CheckPhoneNumberConfirmationCodeRequest struct {
	// The phone number confirmation code
	Code string
}

// CheckPhoneNumberConfirmationCodeRequestTypeID is TL type id of CheckPhoneNumberConfirmationCodeRequest.
const CheckPhoneNumberConfirmationCodeRequestTypeID = 0xafa638da

// Ensuring interfaces in compile-time for CheckPhoneNumberConfirmationCodeRequest.
var (
	_ bin.Encoder     = &CheckPhoneNumberConfirmationCodeRequest{}
	_ bin.Decoder     = &CheckPhoneNumberConfirmationCodeRequest{}
	_ bin.BareEncoder = &CheckPhoneNumberConfirmationCodeRequest{}
	_ bin.BareDecoder = &CheckPhoneNumberConfirmationCodeRequest{}
)

func (c *CheckPhoneNumberConfirmationCodeRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Code == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *CheckPhoneNumberConfirmationCodeRequest) String() string {
	if c == nil {
		return "CheckPhoneNumberConfirmationCodeRequest(nil)"
	}
	type Alias CheckPhoneNumberConfirmationCodeRequest
	return fmt.Sprintf("CheckPhoneNumberConfirmationCodeRequest%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*CheckPhoneNumberConfirmationCodeRequest) TypeID() uint32 {
	return CheckPhoneNumberConfirmationCodeRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*CheckPhoneNumberConfirmationCodeRequest) TypeName() string {
	return "checkPhoneNumberConfirmationCode"
}

// TypeInfo returns info about TL type.
func (c *CheckPhoneNumberConfirmationCodeRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "checkPhoneNumberConfirmationCode",
		ID:   CheckPhoneNumberConfirmationCodeRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Code",
			SchemaName: "code",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *CheckPhoneNumberConfirmationCodeRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode checkPhoneNumberConfirmationCode#afa638da as nil")
	}
	b.PutID(CheckPhoneNumberConfirmationCodeRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *CheckPhoneNumberConfirmationCodeRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode checkPhoneNumberConfirmationCode#afa638da as nil")
	}
	b.PutString(c.Code)
	return nil
}

// Decode implements bin.Decoder.
func (c *CheckPhoneNumberConfirmationCodeRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode checkPhoneNumberConfirmationCode#afa638da to nil")
	}
	if err := b.ConsumeID(CheckPhoneNumberConfirmationCodeRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode checkPhoneNumberConfirmationCode#afa638da: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *CheckPhoneNumberConfirmationCodeRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode checkPhoneNumberConfirmationCode#afa638da to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode checkPhoneNumberConfirmationCode#afa638da: field code: %w", err)
		}
		c.Code = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *CheckPhoneNumberConfirmationCodeRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode checkPhoneNumberConfirmationCode#afa638da as nil")
	}
	b.ObjStart()
	b.PutID("checkPhoneNumberConfirmationCode")
	b.FieldStart("code")
	b.PutString(c.Code)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *CheckPhoneNumberConfirmationCodeRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode checkPhoneNumberConfirmationCode#afa638da to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("checkPhoneNumberConfirmationCode"); err != nil {
				return fmt.Errorf("unable to decode checkPhoneNumberConfirmationCode#afa638da: %w", err)
			}
		case "code":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode checkPhoneNumberConfirmationCode#afa638da: field code: %w", err)
			}
			c.Code = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetCode returns value of Code field.
func (c *CheckPhoneNumberConfirmationCodeRequest) GetCode() (value string) {
	return c.Code
}

// CheckPhoneNumberConfirmationCode invokes method checkPhoneNumberConfirmationCode#afa638da returning error if any.
func (c *Client) CheckPhoneNumberConfirmationCode(ctx context.Context, code string) error {
	var ok Ok

	request := &CheckPhoneNumberConfirmationCodeRequest{
		Code: code,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
