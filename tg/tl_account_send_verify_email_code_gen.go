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

// AccountSendVerifyEmailCodeRequest represents TL type `account.sendVerifyEmailCode#7011509f`.
// Send the verification email code for telegram passport¹.
//
// Links:
//  1) https://core.telegram.org/passport
//
// See https://core.telegram.org/method/account.sendVerifyEmailCode for reference.
type AccountSendVerifyEmailCodeRequest struct {
	// The email where to send the code
	Email string
}

// AccountSendVerifyEmailCodeRequestTypeID is TL type id of AccountSendVerifyEmailCodeRequest.
const AccountSendVerifyEmailCodeRequestTypeID = 0x7011509f

// String implements fmt.Stringer.
func (s *AccountSendVerifyEmailCodeRequest) String() string {
	if s == nil {
		return "AccountSendVerifyEmailCodeRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("AccountSendVerifyEmailCodeRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tEmail: ")
	sb.WriteString(fmt.Sprint(s.Email))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (s *AccountSendVerifyEmailCodeRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode account.sendVerifyEmailCode#7011509f as nil")
	}
	b.PutID(AccountSendVerifyEmailCodeRequestTypeID)
	b.PutString(s.Email)
	return nil
}

// Decode implements bin.Decoder.
func (s *AccountSendVerifyEmailCodeRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode account.sendVerifyEmailCode#7011509f to nil")
	}
	if err := b.ConsumeID(AccountSendVerifyEmailCodeRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.sendVerifyEmailCode#7011509f: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.sendVerifyEmailCode#7011509f: field email: %w", err)
		}
		s.Email = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountSendVerifyEmailCodeRequest.
var (
	_ bin.Encoder = &AccountSendVerifyEmailCodeRequest{}
	_ bin.Decoder = &AccountSendVerifyEmailCodeRequest{}
)

// AccountSendVerifyEmailCode invokes method account.sendVerifyEmailCode#7011509f returning error if any.
// Send the verification email code for telegram passport¹.
//
// Links:
//  1) https://core.telegram.org/passport
//
// See https://core.telegram.org/method/account.sendVerifyEmailCode for reference.
func (c *Client) AccountSendVerifyEmailCode(ctx context.Context, email string) (*AccountSentEmailCode, error) {
	var result AccountSentEmailCode

	request := &AccountSendVerifyEmailCodeRequest{
		Email: email,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
