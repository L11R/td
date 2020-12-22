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

// ChannelsEditCreatorRequest represents TL type `channels.editCreator#8f38cd1f`.
// Transfer channel ownership
//
// See https://core.telegram.org/method/channels.editCreator for reference.
type ChannelsEditCreatorRequest struct {
	// Channel
	Channel InputChannelClass
	// New channel owner
	UserID InputUserClass
	// 2FA password¹ of account
	//
	// Links:
	//  1) https://core.telegram.org/api/srp
	Password InputCheckPasswordSRPClass
}

// ChannelsEditCreatorRequestTypeID is TL type id of ChannelsEditCreatorRequest.
const ChannelsEditCreatorRequestTypeID = 0x8f38cd1f

// String implements fmt.Stringer.
func (e *ChannelsEditCreatorRequest) String() string {
	if e == nil {
		return "ChannelsEditCreatorRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("ChannelsEditCreatorRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tChannel: ")
	sb.WriteString(e.Channel.String())
	sb.WriteString(",\n")
	sb.WriteString("\tUserID: ")
	sb.WriteString(e.UserID.String())
	sb.WriteString(",\n")
	sb.WriteString("\tPassword: ")
	sb.WriteString(e.Password.String())
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (e *ChannelsEditCreatorRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode channels.editCreator#8f38cd1f as nil")
	}
	b.PutID(ChannelsEditCreatorRequestTypeID)
	if e.Channel == nil {
		return fmt.Errorf("unable to encode channels.editCreator#8f38cd1f: field channel is nil")
	}
	if err := e.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.editCreator#8f38cd1f: field channel: %w", err)
	}
	if e.UserID == nil {
		return fmt.Errorf("unable to encode channels.editCreator#8f38cd1f: field user_id is nil")
	}
	if err := e.UserID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.editCreator#8f38cd1f: field user_id: %w", err)
	}
	if e.Password == nil {
		return fmt.Errorf("unable to encode channels.editCreator#8f38cd1f: field password is nil")
	}
	if err := e.Password.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.editCreator#8f38cd1f: field password: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (e *ChannelsEditCreatorRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode channels.editCreator#8f38cd1f to nil")
	}
	if err := b.ConsumeID(ChannelsEditCreatorRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.editCreator#8f38cd1f: %w", err)
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.editCreator#8f38cd1f: field channel: %w", err)
		}
		e.Channel = value
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.editCreator#8f38cd1f: field user_id: %w", err)
		}
		e.UserID = value
	}
	{
		value, err := DecodeInputCheckPasswordSRP(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.editCreator#8f38cd1f: field password: %w", err)
		}
		e.Password = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ChannelsEditCreatorRequest.
var (
	_ bin.Encoder = &ChannelsEditCreatorRequest{}
	_ bin.Decoder = &ChannelsEditCreatorRequest{}
)

// ChannelsEditCreator invokes method channels.editCreator#8f38cd1f returning error if any.
// Transfer channel ownership
//
// See https://core.telegram.org/method/channels.editCreator for reference.
func (c *Client) ChannelsEditCreator(ctx context.Context, request *ChannelsEditCreatorRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
