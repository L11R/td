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

// UploadGetFileRequest represents TL type `upload.getFile#b15a9afc`.
// Returns content of a whole file or its part.
//
// See https://core.telegram.org/method/upload.getFile for reference.
type UploadGetFileRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Disable some checks on limit and offset values, useful for example to stream videos by keyframes
	Precise bool
	// Whether the current client supports CDN downloads¹
	//
	// Links:
	//  1) https://core.telegram.org/cdn
	CDNSupported bool
	// File location
	Location InputFileLocationClass
	// Number of bytes to be skipped
	Offset int
	// Number of bytes to be returned
	Limit int
}

// UploadGetFileRequestTypeID is TL type id of UploadGetFileRequest.
const UploadGetFileRequestTypeID = 0xb15a9afc

// String implements fmt.Stringer.
func (g *UploadGetFileRequest) String() string {
	if g == nil {
		return "UploadGetFileRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("UploadGetFileRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(g.Flags.String())
	sb.WriteString(",\n")
	sb.WriteString("\tLocation: ")
	sb.WriteString(g.Location.String())
	sb.WriteString(",\n")
	sb.WriteString("\tOffset: ")
	sb.WriteString(fmt.Sprint(g.Offset))
	sb.WriteString(",\n")
	sb.WriteString("\tLimit: ")
	sb.WriteString(fmt.Sprint(g.Limit))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (g *UploadGetFileRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode upload.getFile#b15a9afc as nil")
	}
	b.PutID(UploadGetFileRequestTypeID)
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode upload.getFile#b15a9afc: field flags: %w", err)
	}
	if g.Location == nil {
		return fmt.Errorf("unable to encode upload.getFile#b15a9afc: field location is nil")
	}
	if err := g.Location.Encode(b); err != nil {
		return fmt.Errorf("unable to encode upload.getFile#b15a9afc: field location: %w", err)
	}
	b.PutInt(g.Offset)
	b.PutInt(g.Limit)
	return nil
}

// SetPrecise sets value of Precise conditional field.
func (g *UploadGetFileRequest) SetPrecise(value bool) {
	if value {
		g.Flags.Set(0)
	} else {
		g.Flags.Unset(0)
	}
}

// SetCDNSupported sets value of CDNSupported conditional field.
func (g *UploadGetFileRequest) SetCDNSupported(value bool) {
	if value {
		g.Flags.Set(1)
	} else {
		g.Flags.Unset(1)
	}
}

// Decode implements bin.Decoder.
func (g *UploadGetFileRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode upload.getFile#b15a9afc to nil")
	}
	if err := b.ConsumeID(UploadGetFileRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode upload.getFile#b15a9afc: %w", err)
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode upload.getFile#b15a9afc: field flags: %w", err)
		}
	}
	g.Precise = g.Flags.Has(0)
	g.CDNSupported = g.Flags.Has(1)
	{
		value, err := DecodeInputFileLocation(b)
		if err != nil {
			return fmt.Errorf("unable to decode upload.getFile#b15a9afc: field location: %w", err)
		}
		g.Location = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode upload.getFile#b15a9afc: field offset: %w", err)
		}
		g.Offset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode upload.getFile#b15a9afc: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// Ensuring interfaces in compile-time for UploadGetFileRequest.
var (
	_ bin.Encoder = &UploadGetFileRequest{}
	_ bin.Decoder = &UploadGetFileRequest{}
)

// UploadGetFile invokes method upload.getFile#b15a9afc returning error if any.
// Returns content of a whole file or its part.
//
// See https://core.telegram.org/method/upload.getFile for reference.
func (c *Client) UploadGetFile(ctx context.Context, request *UploadGetFileRequest) (UploadFileClass, error) {
	var result UploadFileBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.File, nil
}
