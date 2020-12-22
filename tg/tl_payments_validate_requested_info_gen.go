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

// PaymentsValidateRequestedInfoRequest represents TL type `payments.validateRequestedInfo#770a8e74`.
// Submit requested order information for validation
//
// See https://core.telegram.org/method/payments.validateRequestedInfo for reference.
type PaymentsValidateRequestedInfoRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Save order information to re-use it for future orders
	Save bool
	// Message ID of payment form
	MsgID int
	// Requested order information
	Info PaymentRequestedInfo
}

// PaymentsValidateRequestedInfoRequestTypeID is TL type id of PaymentsValidateRequestedInfoRequest.
const PaymentsValidateRequestedInfoRequestTypeID = 0x770a8e74

// String implements fmt.Stringer.
func (v *PaymentsValidateRequestedInfoRequest) String() string {
	if v == nil {
		return "PaymentsValidateRequestedInfoRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("PaymentsValidateRequestedInfoRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(v.Flags.String())
	sb.WriteString(",\n")
	sb.WriteString("\tMsgID: ")
	sb.WriteString(fmt.Sprint(v.MsgID))
	sb.WriteString(",\n")
	sb.WriteString("\tInfo: ")
	sb.WriteString(v.Info.String())
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (v *PaymentsValidateRequestedInfoRequest) Encode(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't encode payments.validateRequestedInfo#770a8e74 as nil")
	}
	b.PutID(PaymentsValidateRequestedInfoRequestTypeID)
	if err := v.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode payments.validateRequestedInfo#770a8e74: field flags: %w", err)
	}
	b.PutInt(v.MsgID)
	if err := v.Info.Encode(b); err != nil {
		return fmt.Errorf("unable to encode payments.validateRequestedInfo#770a8e74: field info: %w", err)
	}
	return nil
}

// SetSave sets value of Save conditional field.
func (v *PaymentsValidateRequestedInfoRequest) SetSave(value bool) {
	if value {
		v.Flags.Set(0)
	} else {
		v.Flags.Unset(0)
	}
}

// Decode implements bin.Decoder.
func (v *PaymentsValidateRequestedInfoRequest) Decode(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't decode payments.validateRequestedInfo#770a8e74 to nil")
	}
	if err := b.ConsumeID(PaymentsValidateRequestedInfoRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode payments.validateRequestedInfo#770a8e74: %w", err)
	}
	{
		if err := v.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode payments.validateRequestedInfo#770a8e74: field flags: %w", err)
		}
	}
	v.Save = v.Flags.Has(0)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode payments.validateRequestedInfo#770a8e74: field msg_id: %w", err)
		}
		v.MsgID = value
	}
	{
		if err := v.Info.Decode(b); err != nil {
			return fmt.Errorf("unable to decode payments.validateRequestedInfo#770a8e74: field info: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for PaymentsValidateRequestedInfoRequest.
var (
	_ bin.Encoder = &PaymentsValidateRequestedInfoRequest{}
	_ bin.Decoder = &PaymentsValidateRequestedInfoRequest{}
)

// PaymentsValidateRequestedInfo invokes method payments.validateRequestedInfo#770a8e74 returning error if any.
// Submit requested order information for validation
//
// See https://core.telegram.org/method/payments.validateRequestedInfo for reference.
func (c *Client) PaymentsValidateRequestedInfo(ctx context.Context, request *PaymentsValidateRequestedInfoRequest) (*PaymentsValidatedRequestedInfo, error) {
	var result PaymentsValidatedRequestedInfo

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
