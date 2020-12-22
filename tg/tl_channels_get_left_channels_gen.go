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

// ChannelsGetLeftChannelsRequest represents TL type `channels.getLeftChannels#8341ecc0`.
// Get a list of channels/supergroups¹ we left
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// See https://core.telegram.org/method/channels.getLeftChannels for reference.
type ChannelsGetLeftChannelsRequest struct {
	// Offset for pagination¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	Offset int
}

// ChannelsGetLeftChannelsRequestTypeID is TL type id of ChannelsGetLeftChannelsRequest.
const ChannelsGetLeftChannelsRequestTypeID = 0x8341ecc0

// String implements fmt.Stringer.
func (g *ChannelsGetLeftChannelsRequest) String() string {
	if g == nil {
		return "ChannelsGetLeftChannelsRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("ChannelsGetLeftChannelsRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tOffset: ")
	sb.WriteString(fmt.Sprint(g.Offset))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (g *ChannelsGetLeftChannelsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode channels.getLeftChannels#8341ecc0 as nil")
	}
	b.PutID(ChannelsGetLeftChannelsRequestTypeID)
	b.PutInt(g.Offset)
	return nil
}

// Decode implements bin.Decoder.
func (g *ChannelsGetLeftChannelsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode channels.getLeftChannels#8341ecc0 to nil")
	}
	if err := b.ConsumeID(ChannelsGetLeftChannelsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.getLeftChannels#8341ecc0: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode channels.getLeftChannels#8341ecc0: field offset: %w", err)
		}
		g.Offset = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ChannelsGetLeftChannelsRequest.
var (
	_ bin.Encoder = &ChannelsGetLeftChannelsRequest{}
	_ bin.Decoder = &ChannelsGetLeftChannelsRequest{}
)

// ChannelsGetLeftChannels invokes method channels.getLeftChannels#8341ecc0 returning error if any.
// Get a list of channels/supergroups¹ we left
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// See https://core.telegram.org/method/channels.getLeftChannels for reference.
func (c *Client) ChannelsGetLeftChannels(ctx context.Context, offset int) (MessagesChatsClass, error) {
	var result MessagesChatsBox

	request := &ChannelsGetLeftChannelsRequest{
		Offset: offset,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Chats, nil
}
