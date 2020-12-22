// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"fmt"
	"strings"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)
var _ = strings.Builder{}

// UsersSetSecureValueErrorsRequest represents TL type `users.setSecureValueErrors#90c894b5`.
// Notify the user that the sent passport¹ data contains some errors The user will not be able to re-submit their Passport data to you until the errors are fixed (the contents of the field for which you returned the error must change).
// Use this if the data submitted by the user doesn't satisfy the standards your service requires for any reason. For example, if a birthday date seems invalid, a submitted document is blurry, a scan shows evidence of tampering, etc. Supply some details in the error message to make sure the user knows how to correct the issues.
//
// Links:
//  1) https://core.telegram.org/passport
//
// See https://core.telegram.org/method/users.setSecureValueErrors for reference.
type UsersSetSecureValueErrorsRequest struct {
	// The user
	ID InputUserClass
	// Errors
	Errors []SecureValueErrorClass
}

// UsersSetSecureValueErrorsRequestTypeID is TL type id of UsersSetSecureValueErrorsRequest.
const UsersSetSecureValueErrorsRequestTypeID = 0x90c894b5

// String implements fmt.Stringer.
func (s *UsersSetSecureValueErrorsRequest) String() string {
	if s == nil {
		return "UsersSetSecureValueErrorsRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("UsersSetSecureValueErrorsRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tID: ")
	sb.WriteString(s.ID.String())
	sb.WriteString(",\n")
	sb.WriteByte('[')
	for _, v := range s.Errors {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (s *UsersSetSecureValueErrorsRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode users.setSecureValueErrors#90c894b5 as nil")
	}
	b.PutID(UsersSetSecureValueErrorsRequestTypeID)
	if s.ID == nil {
		return fmt.Errorf("unable to encode users.setSecureValueErrors#90c894b5: field id is nil")
	}
	if err := s.ID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode users.setSecureValueErrors#90c894b5: field id: %w", err)
	}
	b.PutVectorHeader(len(s.Errors))
	for idx, v := range s.Errors {
		if v == nil {
			return fmt.Errorf("unable to encode users.setSecureValueErrors#90c894b5: field errors element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode users.setSecureValueErrors#90c894b5: field errors element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *UsersSetSecureValueErrorsRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode users.setSecureValueErrors#90c894b5 to nil")
	}
	if err := b.ConsumeID(UsersSetSecureValueErrorsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode users.setSecureValueErrors#90c894b5: %w", err)
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode users.setSecureValueErrors#90c894b5: field id: %w", err)
		}
		s.ID = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode users.setSecureValueErrors#90c894b5: field errors: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeSecureValueError(b)
			if err != nil {
				return fmt.Errorf("unable to decode users.setSecureValueErrors#90c894b5: field errors: %w", err)
			}
			s.Errors = append(s.Errors, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for UsersSetSecureValueErrorsRequest.
var (
	_ bin.Encoder = &UsersSetSecureValueErrorsRequest{}
	_ bin.Decoder = &UsersSetSecureValueErrorsRequest{}
)

// UsersSetSecureValueErrors invokes method users.setSecureValueErrors#90c894b5 returning error if any.
// Notify the user that the sent passport¹ data contains some errors The user will not be able to re-submit their Passport data to you until the errors are fixed (the contents of the field for which you returned the error must change).
// Use this if the data submitted by the user doesn't satisfy the standards your service requires for any reason. For example, if a birthday date seems invalid, a submitted document is blurry, a scan shows evidence of tampering, etc. Supply some details in the error message to make sure the user knows how to correct the issues.
//
// Links:
//  1) https://core.telegram.org/passport
//
// See https://core.telegram.org/method/users.setSecureValueErrors for reference.
func (c *Client) UsersSetSecureValueErrors(ctx context.Context, request *UsersSetSecureValueErrorsRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
