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

// TopPeer represents TL type `topPeer#edcdc05b`.
// Top peer
//
// See https://core.telegram.org/constructor/topPeer for reference.
type TopPeer struct {
	// Peer
	Peer PeerClass
	// Rating as computed in top peer rating »¹
	//
	// Links:
	//  1) https://core.telegram.org/api/top-rating
	Rating float64
}

// TopPeerTypeID is TL type id of TopPeer.
const TopPeerTypeID = 0xedcdc05b

// String implements fmt.Stringer.
func (t *TopPeer) String() string {
	if t == nil {
		return "TopPeer(nil)"
	}
	var sb strings.Builder
	sb.WriteString("TopPeer")
	sb.WriteString("{\n")
	sb.WriteString("\tPeer: ")
	sb.WriteString(t.Peer.String())
	sb.WriteString(",\n")
	sb.WriteString("\tRating: ")
	sb.WriteString(fmt.Sprint(t.Rating))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (t *TopPeer) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode topPeer#edcdc05b as nil")
	}
	b.PutID(TopPeerTypeID)
	if t.Peer == nil {
		return fmt.Errorf("unable to encode topPeer#edcdc05b: field peer is nil")
	}
	if err := t.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode topPeer#edcdc05b: field peer: %w", err)
	}
	b.PutDouble(t.Rating)
	return nil
}

// Decode implements bin.Decoder.
func (t *TopPeer) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode topPeer#edcdc05b to nil")
	}
	if err := b.ConsumeID(TopPeerTypeID); err != nil {
		return fmt.Errorf("unable to decode topPeer#edcdc05b: %w", err)
	}
	{
		value, err := DecodePeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode topPeer#edcdc05b: field peer: %w", err)
		}
		t.Peer = value
	}
	{
		value, err := b.Double()
		if err != nil {
			return fmt.Errorf("unable to decode topPeer#edcdc05b: field rating: %w", err)
		}
		t.Rating = value
	}
	return nil
}

// Ensuring interfaces in compile-time for TopPeer.
var (
	_ bin.Encoder = &TopPeer{}
	_ bin.Decoder = &TopPeer{}
)
