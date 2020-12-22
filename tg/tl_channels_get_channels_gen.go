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

// ChannelsGetChannelsRequest represents TL type `channels.getChannels#a7f6bbb`.
// Get info about channels/supergroups¹
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// See https://core.telegram.org/method/channels.getChannels for reference.
type ChannelsGetChannelsRequest struct {
	// IDs of channels/supergroups to get info about
	ID []InputChannelClass
}

// ChannelsGetChannelsRequestTypeID is TL type id of ChannelsGetChannelsRequest.
const ChannelsGetChannelsRequestTypeID = 0xa7f6bbb

// String implements fmt.Stringer.
func (g *ChannelsGetChannelsRequest) String() string {
	if g == nil {
		return "ChannelsGetChannelsRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("ChannelsGetChannelsRequest")
	sb.WriteString("{\n")
	sb.WriteByte('[')
	for _, v := range g.ID {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (g *ChannelsGetChannelsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode channels.getChannels#a7f6bbb as nil")
	}
	b.PutID(ChannelsGetChannelsRequestTypeID)
	b.PutVectorHeader(len(g.ID))
	for idx, v := range g.ID {
		if v == nil {
			return fmt.Errorf("unable to encode channels.getChannels#a7f6bbb: field id element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode channels.getChannels#a7f6bbb: field id element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *ChannelsGetChannelsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode channels.getChannels#a7f6bbb to nil")
	}
	if err := b.ConsumeID(ChannelsGetChannelsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.getChannels#a7f6bbb: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode channels.getChannels#a7f6bbb: field id: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInputChannel(b)
			if err != nil {
				return fmt.Errorf("unable to decode channels.getChannels#a7f6bbb: field id: %w", err)
			}
			g.ID = append(g.ID, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for ChannelsGetChannelsRequest.
var (
	_ bin.Encoder = &ChannelsGetChannelsRequest{}
	_ bin.Decoder = &ChannelsGetChannelsRequest{}
)

// ChannelsGetChannels invokes method channels.getChannels#a7f6bbb returning error if any.
// Get info about channels/supergroups¹
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// See https://core.telegram.org/method/channels.getChannels for reference.
func (c *Client) ChannelsGetChannels(ctx context.Context, id []InputChannelClass) (MessagesChatsClass, error) {
	var result MessagesChatsBox

	request := &ChannelsGetChannelsRequest{
		ID: id,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Chats, nil
}
