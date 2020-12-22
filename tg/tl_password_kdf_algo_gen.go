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

// PasswordKdfAlgoUnknown represents TL type `passwordKdfAlgoUnknown#d45ab096`.
// Unknown KDF (most likely, the client is outdated and does not support the specified KDF algorithm)
//
// See https://core.telegram.org/constructor/passwordKdfAlgoUnknown for reference.
type PasswordKdfAlgoUnknown struct {
}

// PasswordKdfAlgoUnknownTypeID is TL type id of PasswordKdfAlgoUnknown.
const PasswordKdfAlgoUnknownTypeID = 0xd45ab096

// String implements fmt.Stringer.
func (p *PasswordKdfAlgoUnknown) String() string {
	if p == nil {
		return "PasswordKdfAlgoUnknown(nil)"
	}
	var sb strings.Builder
	sb.WriteString("PasswordKdfAlgoUnknown")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (p *PasswordKdfAlgoUnknown) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode passwordKdfAlgoUnknown#d45ab096 as nil")
	}
	b.PutID(PasswordKdfAlgoUnknownTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (p *PasswordKdfAlgoUnknown) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode passwordKdfAlgoUnknown#d45ab096 to nil")
	}
	if err := b.ConsumeID(PasswordKdfAlgoUnknownTypeID); err != nil {
		return fmt.Errorf("unable to decode passwordKdfAlgoUnknown#d45ab096: %w", err)
	}
	return nil
}

// construct implements constructor of PasswordKdfAlgoClass.
func (p PasswordKdfAlgoUnknown) construct() PasswordKdfAlgoClass { return &p }

// Ensuring interfaces in compile-time for PasswordKdfAlgoUnknown.
var (
	_ bin.Encoder = &PasswordKdfAlgoUnknown{}
	_ bin.Decoder = &PasswordKdfAlgoUnknown{}

	_ PasswordKdfAlgoClass = &PasswordKdfAlgoUnknown{}
)

// PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow represents TL type `passwordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow#3a912d4a`.
// This key derivation algorithm defines that SRP 2FA login¹ must be used
//
// Links:
//  1) https://core.telegram.org/api/srp
//
// See https://core.telegram.org/constructor/passwordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow for reference.
type PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow struct {
	// One of two salts used by the derivation function (see SRP 2FA login¹)
	//
	// Links:
	//  1) https://core.telegram.org/api/srp
	Salt1 []byte
	// One of two salts used by the derivation function (see SRP 2FA login¹)
	//
	// Links:
	//  1) https://core.telegram.org/api/srp
	Salt2 []byte
	// Base (see SRP 2FA login¹)
	//
	// Links:
	//  1) https://core.telegram.org/api/srp
	G int
	// 2048-bit modulus (see SRP 2FA login¹)
	//
	// Links:
	//  1) https://core.telegram.org/api/srp
	P []byte
}

// PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPowTypeID is TL type id of PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow.
const PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPowTypeID = 0x3a912d4a

// String implements fmt.Stringer.
func (p *PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow) String() string {
	if p == nil {
		return "PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow(nil)"
	}
	var sb strings.Builder
	sb.WriteString("PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow")
	sb.WriteString("{\n")
	sb.WriteString("\tSalt1: ")
	sb.WriteString(fmt.Sprint(p.Salt1))
	sb.WriteString(",\n")
	sb.WriteString("\tSalt2: ")
	sb.WriteString(fmt.Sprint(p.Salt2))
	sb.WriteString(",\n")
	sb.WriteString("\tG: ")
	sb.WriteString(fmt.Sprint(p.G))
	sb.WriteString(",\n")
	sb.WriteString("\tP: ")
	sb.WriteString(fmt.Sprint(p.P))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (p *PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode passwordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow#3a912d4a as nil")
	}
	b.PutID(PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPowTypeID)
	b.PutBytes(p.Salt1)
	b.PutBytes(p.Salt2)
	b.PutInt(p.G)
	b.PutBytes(p.P)
	return nil
}

// Decode implements bin.Decoder.
func (p *PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode passwordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow#3a912d4a to nil")
	}
	if err := b.ConsumeID(PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPowTypeID); err != nil {
		return fmt.Errorf("unable to decode passwordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow#3a912d4a: %w", err)
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode passwordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow#3a912d4a: field salt1: %w", err)
		}
		p.Salt1 = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode passwordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow#3a912d4a: field salt2: %w", err)
		}
		p.Salt2 = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode passwordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow#3a912d4a: field g: %w", err)
		}
		p.G = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode passwordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow#3a912d4a: field p: %w", err)
		}
		p.P = value
	}
	return nil
}

// construct implements constructor of PasswordKdfAlgoClass.
func (p PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow) construct() PasswordKdfAlgoClass {
	return &p
}

// Ensuring interfaces in compile-time for PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow.
var (
	_ bin.Encoder = &PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow{}
	_ bin.Decoder = &PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow{}

	_ PasswordKdfAlgoClass = &PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow{}
)

// PasswordKdfAlgoClass represents PasswordKdfAlgo generic type.
//
// See https://core.telegram.org/type/PasswordKdfAlgo for reference.
//
// Example:
//  g, err := DecodePasswordKdfAlgo(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *PasswordKdfAlgoUnknown: // passwordKdfAlgoUnknown#d45ab096
//  case *PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow: // passwordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow#3a912d4a
//  default: panic(v)
//  }
type PasswordKdfAlgoClass interface {
	bin.Encoder
	bin.Decoder
	construct() PasswordKdfAlgoClass
	fmt.Stringer
}

// DecodePasswordKdfAlgo implements binary de-serialization for PasswordKdfAlgoClass.
func DecodePasswordKdfAlgo(buf *bin.Buffer) (PasswordKdfAlgoClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case PasswordKdfAlgoUnknownTypeID:
		// Decoding passwordKdfAlgoUnknown#d45ab096.
		v := PasswordKdfAlgoUnknown{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PasswordKdfAlgoClass: %w", err)
		}
		return &v, nil
	case PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPowTypeID:
		// Decoding passwordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow#3a912d4a.
		v := PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PasswordKdfAlgoClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode PasswordKdfAlgoClass: %w", bin.NewUnexpectedID(id))
	}
}

// PasswordKdfAlgo boxes the PasswordKdfAlgoClass providing a helper.
type PasswordKdfAlgoBox struct {
	PasswordKdfAlgo PasswordKdfAlgoClass
}

// Decode implements bin.Decoder for PasswordKdfAlgoBox.
func (b *PasswordKdfAlgoBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode PasswordKdfAlgoBox to nil")
	}
	v, err := DecodePasswordKdfAlgo(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.PasswordKdfAlgo = v
	return nil
}

// Encode implements bin.Encode for PasswordKdfAlgoBox.
func (b *PasswordKdfAlgoBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.PasswordKdfAlgo == nil {
		return fmt.Errorf("unable to encode PasswordKdfAlgoClass as nil")
	}
	return b.PasswordKdfAlgo.Encode(buf)
}
