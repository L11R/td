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

// AuthAuthorization represents TL type `auth.authorization#cd050916`.
// Contains user authorization info.
//
// See https://core.telegram.org/constructor/auth.authorization for reference.
type AuthAuthorization struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Temporary passport¹ sessions
	//
	// Links:
	//  1) https://core.telegram.org/passport
	//
	// Use SetTmpSessions and GetTmpSessions helpers.
	TmpSessions int
	// Info on authorized user
	User UserClass
}

// AuthAuthorizationTypeID is TL type id of AuthAuthorization.
const AuthAuthorizationTypeID = 0xcd050916

// String implements fmt.Stringer.
func (a *AuthAuthorization) String() string {
	if a == nil {
		return "AuthAuthorization(nil)"
	}
	var sb strings.Builder
	sb.WriteString("AuthAuthorization")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(a.Flags.String())
	sb.WriteString(",\n")
	if a.Flags.Has(0) {
		sb.WriteString("\tTmpSessions: ")
		sb.WriteString(fmt.Sprint(a.TmpSessions))
		sb.WriteString(",\n")
	}
	sb.WriteString("\tUser: ")
	sb.WriteString(a.User.String())
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (a *AuthAuthorization) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode auth.authorization#cd050916 as nil")
	}
	b.PutID(AuthAuthorizationTypeID)
	if err := a.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode auth.authorization#cd050916: field flags: %w", err)
	}
	if a.Flags.Has(0) {
		b.PutInt(a.TmpSessions)
	}
	if a.User == nil {
		return fmt.Errorf("unable to encode auth.authorization#cd050916: field user is nil")
	}
	if err := a.User.Encode(b); err != nil {
		return fmt.Errorf("unable to encode auth.authorization#cd050916: field user: %w", err)
	}
	return nil
}

// SetTmpSessions sets value of TmpSessions conditional field.
func (a *AuthAuthorization) SetTmpSessions(value int) {
	a.Flags.Set(0)
	a.TmpSessions = value
}

// GetTmpSessions returns value of TmpSessions conditional field and
// boolean which is true if field was set.
func (a *AuthAuthorization) GetTmpSessions() (value int, ok bool) {
	if !a.Flags.Has(0) {
		return value, false
	}
	return a.TmpSessions, true
}

// Decode implements bin.Decoder.
func (a *AuthAuthorization) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode auth.authorization#cd050916 to nil")
	}
	if err := b.ConsumeID(AuthAuthorizationTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.authorization#cd050916: %w", err)
	}
	{
		if err := a.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode auth.authorization#cd050916: field flags: %w", err)
		}
	}
	if a.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode auth.authorization#cd050916: field tmp_sessions: %w", err)
		}
		a.TmpSessions = value
	}
	{
		value, err := DecodeUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode auth.authorization#cd050916: field user: %w", err)
		}
		a.User = value
	}
	return nil
}

// construct implements constructor of AuthAuthorizationClass.
func (a AuthAuthorization) construct() AuthAuthorizationClass { return &a }

// Ensuring interfaces in compile-time for AuthAuthorization.
var (
	_ bin.Encoder = &AuthAuthorization{}
	_ bin.Decoder = &AuthAuthorization{}

	_ AuthAuthorizationClass = &AuthAuthorization{}
)

// AuthAuthorizationSignUpRequired represents TL type `auth.authorizationSignUpRequired#44747e9a`.
// An account with this phone number doesn't exist on telegram: the user has to enter basic information and sign up¹
//
// Links:
//  1) https://core.telegram.org/api/auth
//
// See https://core.telegram.org/constructor/auth.authorizationSignUpRequired for reference.
type AuthAuthorizationSignUpRequired struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Telegram's terms of service: the user must read and accept the terms of service before signing up to telegram
	//
	// Use SetTermsOfService and GetTermsOfService helpers.
	TermsOfService HelpTermsOfService
}

// AuthAuthorizationSignUpRequiredTypeID is TL type id of AuthAuthorizationSignUpRequired.
const AuthAuthorizationSignUpRequiredTypeID = 0x44747e9a

// String implements fmt.Stringer.
func (a *AuthAuthorizationSignUpRequired) String() string {
	if a == nil {
		return "AuthAuthorizationSignUpRequired(nil)"
	}
	var sb strings.Builder
	sb.WriteString("AuthAuthorizationSignUpRequired")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(a.Flags.String())
	sb.WriteString(",\n")
	if a.Flags.Has(0) {
		sb.WriteString("\tTermsOfService: ")
		sb.WriteString(a.TermsOfService.String())
		sb.WriteString(",\n")
	}
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (a *AuthAuthorizationSignUpRequired) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode auth.authorizationSignUpRequired#44747e9a as nil")
	}
	b.PutID(AuthAuthorizationSignUpRequiredTypeID)
	if err := a.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode auth.authorizationSignUpRequired#44747e9a: field flags: %w", err)
	}
	if a.Flags.Has(0) {
		if err := a.TermsOfService.Encode(b); err != nil {
			return fmt.Errorf("unable to encode auth.authorizationSignUpRequired#44747e9a: field terms_of_service: %w", err)
		}
	}
	return nil
}

// SetTermsOfService sets value of TermsOfService conditional field.
func (a *AuthAuthorizationSignUpRequired) SetTermsOfService(value HelpTermsOfService) {
	a.Flags.Set(0)
	a.TermsOfService = value
}

// GetTermsOfService returns value of TermsOfService conditional field and
// boolean which is true if field was set.
func (a *AuthAuthorizationSignUpRequired) GetTermsOfService() (value HelpTermsOfService, ok bool) {
	if !a.Flags.Has(0) {
		return value, false
	}
	return a.TermsOfService, true
}

// Decode implements bin.Decoder.
func (a *AuthAuthorizationSignUpRequired) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode auth.authorizationSignUpRequired#44747e9a to nil")
	}
	if err := b.ConsumeID(AuthAuthorizationSignUpRequiredTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.authorizationSignUpRequired#44747e9a: %w", err)
	}
	{
		if err := a.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode auth.authorizationSignUpRequired#44747e9a: field flags: %w", err)
		}
	}
	if a.Flags.Has(0) {
		if err := a.TermsOfService.Decode(b); err != nil {
			return fmt.Errorf("unable to decode auth.authorizationSignUpRequired#44747e9a: field terms_of_service: %w", err)
		}
	}
	return nil
}

// construct implements constructor of AuthAuthorizationClass.
func (a AuthAuthorizationSignUpRequired) construct() AuthAuthorizationClass { return &a }

// Ensuring interfaces in compile-time for AuthAuthorizationSignUpRequired.
var (
	_ bin.Encoder = &AuthAuthorizationSignUpRequired{}
	_ bin.Decoder = &AuthAuthorizationSignUpRequired{}

	_ AuthAuthorizationClass = &AuthAuthorizationSignUpRequired{}
)

// AuthAuthorizationClass represents auth.Authorization generic type.
//
// See https://core.telegram.org/type/auth.Authorization for reference.
//
// Example:
//  g, err := DecodeAuthAuthorization(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *AuthAuthorization: // auth.authorization#cd050916
//  case *AuthAuthorizationSignUpRequired: // auth.authorizationSignUpRequired#44747e9a
//  default: panic(v)
//  }
type AuthAuthorizationClass interface {
	bin.Encoder
	bin.Decoder
	construct() AuthAuthorizationClass
	fmt.Stringer
}

// DecodeAuthAuthorization implements binary de-serialization for AuthAuthorizationClass.
func DecodeAuthAuthorization(buf *bin.Buffer) (AuthAuthorizationClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case AuthAuthorizationTypeID:
		// Decoding auth.authorization#cd050916.
		v := AuthAuthorization{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AuthAuthorizationClass: %w", err)
		}
		return &v, nil
	case AuthAuthorizationSignUpRequiredTypeID:
		// Decoding auth.authorizationSignUpRequired#44747e9a.
		v := AuthAuthorizationSignUpRequired{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AuthAuthorizationClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode AuthAuthorizationClass: %w", bin.NewUnexpectedID(id))
	}
}

// AuthAuthorization boxes the AuthAuthorizationClass providing a helper.
type AuthAuthorizationBox struct {
	Authorization AuthAuthorizationClass
}

// Decode implements bin.Decoder for AuthAuthorizationBox.
func (b *AuthAuthorizationBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode AuthAuthorizationBox to nil")
	}
	v, err := DecodeAuthAuthorization(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.Authorization = v
	return nil
}

// Encode implements bin.Encode for AuthAuthorizationBox.
func (b *AuthAuthorizationBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.Authorization == nil {
		return fmt.Errorf("unable to encode AuthAuthorizationClass as nil")
	}
	return b.Authorization.Encode(buf)
}
