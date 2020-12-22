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

// InputGeoPointEmpty represents TL type `inputGeoPointEmpty#e4c123d6`.
// Empty GeoPoint constructor.
//
// See https://core.telegram.org/constructor/inputGeoPointEmpty for reference.
type InputGeoPointEmpty struct {
}

// InputGeoPointEmptyTypeID is TL type id of InputGeoPointEmpty.
const InputGeoPointEmptyTypeID = 0xe4c123d6

// String implements fmt.Stringer.
func (i *InputGeoPointEmpty) String() string {
	if i == nil {
		return "InputGeoPointEmpty(nil)"
	}
	var sb strings.Builder
	sb.WriteString("InputGeoPointEmpty")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (i *InputGeoPointEmpty) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputGeoPointEmpty#e4c123d6 as nil")
	}
	b.PutID(InputGeoPointEmptyTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputGeoPointEmpty) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputGeoPointEmpty#e4c123d6 to nil")
	}
	if err := b.ConsumeID(InputGeoPointEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode inputGeoPointEmpty#e4c123d6: %w", err)
	}
	return nil
}

// construct implements constructor of InputGeoPointClass.
func (i InputGeoPointEmpty) construct() InputGeoPointClass { return &i }

// Ensuring interfaces in compile-time for InputGeoPointEmpty.
var (
	_ bin.Encoder = &InputGeoPointEmpty{}
	_ bin.Decoder = &InputGeoPointEmpty{}

	_ InputGeoPointClass = &InputGeoPointEmpty{}
)

// InputGeoPoint represents TL type `inputGeoPoint#48222faf`.
// Defines a GeoPoint by its coordinates.
//
// See https://core.telegram.org/constructor/inputGeoPoint for reference.
type InputGeoPoint struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Latitide
	Lat float64
	// Longtitude
	Long float64
	// The estimated horizontal accuracy of the location, in meters; as defined by the sender.
	//
	// Use SetAccuracyRadius and GetAccuracyRadius helpers.
	AccuracyRadius int
}

// InputGeoPointTypeID is TL type id of InputGeoPoint.
const InputGeoPointTypeID = 0x48222faf

// String implements fmt.Stringer.
func (i *InputGeoPoint) String() string {
	if i == nil {
		return "InputGeoPoint(nil)"
	}
	var sb strings.Builder
	sb.WriteString("InputGeoPoint")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(i.Flags.String())
	sb.WriteString(",\n")
	sb.WriteString("\tLat: ")
	sb.WriteString(fmt.Sprint(i.Lat))
	sb.WriteString(",\n")
	sb.WriteString("\tLong: ")
	sb.WriteString(fmt.Sprint(i.Long))
	sb.WriteString(",\n")
	if i.Flags.Has(0) {
		sb.WriteString("\tAccuracyRadius: ")
		sb.WriteString(fmt.Sprint(i.AccuracyRadius))
		sb.WriteString(",\n")
	}
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (i *InputGeoPoint) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputGeoPoint#48222faf as nil")
	}
	b.PutID(InputGeoPointTypeID)
	if err := i.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputGeoPoint#48222faf: field flags: %w", err)
	}
	b.PutDouble(i.Lat)
	b.PutDouble(i.Long)
	if i.Flags.Has(0) {
		b.PutInt(i.AccuracyRadius)
	}
	return nil
}

// SetAccuracyRadius sets value of AccuracyRadius conditional field.
func (i *InputGeoPoint) SetAccuracyRadius(value int) {
	i.Flags.Set(0)
	i.AccuracyRadius = value
}

// GetAccuracyRadius returns value of AccuracyRadius conditional field and
// boolean which is true if field was set.
func (i *InputGeoPoint) GetAccuracyRadius() (value int, ok bool) {
	if !i.Flags.Has(0) {
		return value, false
	}
	return i.AccuracyRadius, true
}

// Decode implements bin.Decoder.
func (i *InputGeoPoint) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputGeoPoint#48222faf to nil")
	}
	if err := b.ConsumeID(InputGeoPointTypeID); err != nil {
		return fmt.Errorf("unable to decode inputGeoPoint#48222faf: %w", err)
	}
	{
		if err := i.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode inputGeoPoint#48222faf: field flags: %w", err)
		}
	}
	{
		value, err := b.Double()
		if err != nil {
			return fmt.Errorf("unable to decode inputGeoPoint#48222faf: field lat: %w", err)
		}
		i.Lat = value
	}
	{
		value, err := b.Double()
		if err != nil {
			return fmt.Errorf("unable to decode inputGeoPoint#48222faf: field long: %w", err)
		}
		i.Long = value
	}
	if i.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputGeoPoint#48222faf: field accuracy_radius: %w", err)
		}
		i.AccuracyRadius = value
	}
	return nil
}

// construct implements constructor of InputGeoPointClass.
func (i InputGeoPoint) construct() InputGeoPointClass { return &i }

// Ensuring interfaces in compile-time for InputGeoPoint.
var (
	_ bin.Encoder = &InputGeoPoint{}
	_ bin.Decoder = &InputGeoPoint{}

	_ InputGeoPointClass = &InputGeoPoint{}
)

// InputGeoPointClass represents InputGeoPoint generic type.
//
// See https://core.telegram.org/type/InputGeoPoint for reference.
//
// Example:
//  g, err := DecodeInputGeoPoint(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *InputGeoPointEmpty: // inputGeoPointEmpty#e4c123d6
//  case *InputGeoPoint: // inputGeoPoint#48222faf
//  default: panic(v)
//  }
type InputGeoPointClass interface {
	bin.Encoder
	bin.Decoder
	construct() InputGeoPointClass
	fmt.Stringer
}

// DecodeInputGeoPoint implements binary de-serialization for InputGeoPointClass.
func DecodeInputGeoPoint(buf *bin.Buffer) (InputGeoPointClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputGeoPointEmptyTypeID:
		// Decoding inputGeoPointEmpty#e4c123d6.
		v := InputGeoPointEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputGeoPointClass: %w", err)
		}
		return &v, nil
	case InputGeoPointTypeID:
		// Decoding inputGeoPoint#48222faf.
		v := InputGeoPoint{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputGeoPointClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputGeoPointClass: %w", bin.NewUnexpectedID(id))
	}
}

// InputGeoPoint boxes the InputGeoPointClass providing a helper.
type InputGeoPointBox struct {
	InputGeoPoint InputGeoPointClass
}

// Decode implements bin.Decoder for InputGeoPointBox.
func (b *InputGeoPointBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputGeoPointBox to nil")
	}
	v, err := DecodeInputGeoPoint(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputGeoPoint = v
	return nil
}

// Encode implements bin.Encode for InputGeoPointBox.
func (b *InputGeoPointBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputGeoPoint == nil {
		return fmt.Errorf("unable to encode InputGeoPointClass as nil")
	}
	return b.InputGeoPoint.Encode(buf)
}
