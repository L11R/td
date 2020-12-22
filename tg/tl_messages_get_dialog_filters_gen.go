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

// MessagesGetDialogFiltersRequest represents TL type `messages.getDialogFilters#f19ed96d`.
// Get folders¹
//
// Links:
//  1) https://core.telegram.org/api/folders
//
// See https://core.telegram.org/method/messages.getDialogFilters for reference.
type MessagesGetDialogFiltersRequest struct {
}

// MessagesGetDialogFiltersRequestTypeID is TL type id of MessagesGetDialogFiltersRequest.
const MessagesGetDialogFiltersRequestTypeID = 0xf19ed96d

// String implements fmt.Stringer.
func (g *MessagesGetDialogFiltersRequest) String() string {
	if g == nil {
		return "MessagesGetDialogFiltersRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesGetDialogFiltersRequest")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (g *MessagesGetDialogFiltersRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getDialogFilters#f19ed96d as nil")
	}
	b.PutID(MessagesGetDialogFiltersRequestTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetDialogFiltersRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getDialogFilters#f19ed96d to nil")
	}
	if err := b.ConsumeID(MessagesGetDialogFiltersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getDialogFilters#f19ed96d: %w", err)
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetDialogFiltersRequest.
var (
	_ bin.Encoder = &MessagesGetDialogFiltersRequest{}
	_ bin.Decoder = &MessagesGetDialogFiltersRequest{}
)

// MessagesGetDialogFilters invokes method messages.getDialogFilters#f19ed96d returning error if any.
// Get folders¹
//
// Links:
//  1) https://core.telegram.org/api/folders
//
// See https://core.telegram.org/method/messages.getDialogFilters for reference.
func (c *Client) MessagesGetDialogFilters(ctx context.Context) ([]DialogFilter, error) {
	var result DialogFilterVector

	request := &MessagesGetDialogFiltersRequest{}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Elems, nil
}
