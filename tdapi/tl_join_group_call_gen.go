// Code generated by gotdgen, DO NOT EDIT.

package tdapi

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// JoinGroupCallRequest represents TL type `joinGroupCall#c1c947e5`.
type JoinGroupCallRequest struct {
	// Group call identifier
	GroupCallID int32
	// Identifier of a group call participant, which will be used to join the call; pass null
	// to join as self; video chats only
	ParticipantID MessageSenderClass
	// Caller audio channel synchronization source identifier; received from tgcalls
	AudioSourceID int32
	// Group call join payload; received from tgcalls
	Payload string
	// True, if the user's microphone is muted
	IsMuted bool
	// True, if the user's video is enabled
	IsMyVideoEnabled bool
	// If non-empty, invite hash to be used to join the group call without being muted by
	// administrators
	InviteHash string
}

// JoinGroupCallRequestTypeID is TL type id of JoinGroupCallRequest.
const JoinGroupCallRequestTypeID = 0xc1c947e5

// Ensuring interfaces in compile-time for JoinGroupCallRequest.
var (
	_ bin.Encoder     = &JoinGroupCallRequest{}
	_ bin.Decoder     = &JoinGroupCallRequest{}
	_ bin.BareEncoder = &JoinGroupCallRequest{}
	_ bin.BareDecoder = &JoinGroupCallRequest{}
)

func (j *JoinGroupCallRequest) Zero() bool {
	if j == nil {
		return true
	}
	if !(j.GroupCallID == 0) {
		return false
	}
	if !(j.ParticipantID == nil) {
		return false
	}
	if !(j.AudioSourceID == 0) {
		return false
	}
	if !(j.Payload == "") {
		return false
	}
	if !(j.IsMuted == false) {
		return false
	}
	if !(j.IsMyVideoEnabled == false) {
		return false
	}
	if !(j.InviteHash == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (j *JoinGroupCallRequest) String() string {
	if j == nil {
		return "JoinGroupCallRequest(nil)"
	}
	type Alias JoinGroupCallRequest
	return fmt.Sprintf("JoinGroupCallRequest%+v", Alias(*j))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*JoinGroupCallRequest) TypeID() uint32 {
	return JoinGroupCallRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*JoinGroupCallRequest) TypeName() string {
	return "joinGroupCall"
}

// TypeInfo returns info about TL type.
func (j *JoinGroupCallRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "joinGroupCall",
		ID:   JoinGroupCallRequestTypeID,
	}
	if j == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "GroupCallID",
			SchemaName: "group_call_id",
		},
		{
			Name:       "ParticipantID",
			SchemaName: "participant_id",
		},
		{
			Name:       "AudioSourceID",
			SchemaName: "audio_source_id",
		},
		{
			Name:       "Payload",
			SchemaName: "payload",
		},
		{
			Name:       "IsMuted",
			SchemaName: "is_muted",
		},
		{
			Name:       "IsMyVideoEnabled",
			SchemaName: "is_my_video_enabled",
		},
		{
			Name:       "InviteHash",
			SchemaName: "invite_hash",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (j *JoinGroupCallRequest) Encode(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't encode joinGroupCall#c1c947e5 as nil")
	}
	b.PutID(JoinGroupCallRequestTypeID)
	return j.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (j *JoinGroupCallRequest) EncodeBare(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't encode joinGroupCall#c1c947e5 as nil")
	}
	b.PutInt32(j.GroupCallID)
	if j.ParticipantID == nil {
		return fmt.Errorf("unable to encode joinGroupCall#c1c947e5: field participant_id is nil")
	}
	if err := j.ParticipantID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode joinGroupCall#c1c947e5: field participant_id: %w", err)
	}
	b.PutInt32(j.AudioSourceID)
	b.PutString(j.Payload)
	b.PutBool(j.IsMuted)
	b.PutBool(j.IsMyVideoEnabled)
	b.PutString(j.InviteHash)
	return nil
}

// Decode implements bin.Decoder.
func (j *JoinGroupCallRequest) Decode(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't decode joinGroupCall#c1c947e5 to nil")
	}
	if err := b.ConsumeID(JoinGroupCallRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode joinGroupCall#c1c947e5: %w", err)
	}
	return j.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (j *JoinGroupCallRequest) DecodeBare(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't decode joinGroupCall#c1c947e5 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode joinGroupCall#c1c947e5: field group_call_id: %w", err)
		}
		j.GroupCallID = value
	}
	{
		value, err := DecodeMessageSender(b)
		if err != nil {
			return fmt.Errorf("unable to decode joinGroupCall#c1c947e5: field participant_id: %w", err)
		}
		j.ParticipantID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode joinGroupCall#c1c947e5: field audio_source_id: %w", err)
		}
		j.AudioSourceID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode joinGroupCall#c1c947e5: field payload: %w", err)
		}
		j.Payload = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode joinGroupCall#c1c947e5: field is_muted: %w", err)
		}
		j.IsMuted = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode joinGroupCall#c1c947e5: field is_my_video_enabled: %w", err)
		}
		j.IsMyVideoEnabled = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode joinGroupCall#c1c947e5: field invite_hash: %w", err)
		}
		j.InviteHash = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (j *JoinGroupCallRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if j == nil {
		return fmt.Errorf("can't encode joinGroupCall#c1c947e5 as nil")
	}
	b.ObjStart()
	b.PutID("joinGroupCall")
	b.FieldStart("group_call_id")
	b.PutInt32(j.GroupCallID)
	b.FieldStart("participant_id")
	if j.ParticipantID == nil {
		return fmt.Errorf("unable to encode joinGroupCall#c1c947e5: field participant_id is nil")
	}
	if err := j.ParticipantID.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode joinGroupCall#c1c947e5: field participant_id: %w", err)
	}
	b.FieldStart("audio_source_id")
	b.PutInt32(j.AudioSourceID)
	b.FieldStart("payload")
	b.PutString(j.Payload)
	b.FieldStart("is_muted")
	b.PutBool(j.IsMuted)
	b.FieldStart("is_my_video_enabled")
	b.PutBool(j.IsMyVideoEnabled)
	b.FieldStart("invite_hash")
	b.PutString(j.InviteHash)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (j *JoinGroupCallRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if j == nil {
		return fmt.Errorf("can't decode joinGroupCall#c1c947e5 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("joinGroupCall"); err != nil {
				return fmt.Errorf("unable to decode joinGroupCall#c1c947e5: %w", err)
			}
		case "group_call_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode joinGroupCall#c1c947e5: field group_call_id: %w", err)
			}
			j.GroupCallID = value
		case "participant_id":
			value, err := DecodeTDLibJSONMessageSender(b)
			if err != nil {
				return fmt.Errorf("unable to decode joinGroupCall#c1c947e5: field participant_id: %w", err)
			}
			j.ParticipantID = value
		case "audio_source_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode joinGroupCall#c1c947e5: field audio_source_id: %w", err)
			}
			j.AudioSourceID = value
		case "payload":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode joinGroupCall#c1c947e5: field payload: %w", err)
			}
			j.Payload = value
		case "is_muted":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode joinGroupCall#c1c947e5: field is_muted: %w", err)
			}
			j.IsMuted = value
		case "is_my_video_enabled":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode joinGroupCall#c1c947e5: field is_my_video_enabled: %w", err)
			}
			j.IsMyVideoEnabled = value
		case "invite_hash":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode joinGroupCall#c1c947e5: field invite_hash: %w", err)
			}
			j.InviteHash = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetGroupCallID returns value of GroupCallID field.
func (j *JoinGroupCallRequest) GetGroupCallID() (value int32) {
	return j.GroupCallID
}

// GetParticipantID returns value of ParticipantID field.
func (j *JoinGroupCallRequest) GetParticipantID() (value MessageSenderClass) {
	return j.ParticipantID
}

// GetAudioSourceID returns value of AudioSourceID field.
func (j *JoinGroupCallRequest) GetAudioSourceID() (value int32) {
	return j.AudioSourceID
}

// GetPayload returns value of Payload field.
func (j *JoinGroupCallRequest) GetPayload() (value string) {
	return j.Payload
}

// GetIsMuted returns value of IsMuted field.
func (j *JoinGroupCallRequest) GetIsMuted() (value bool) {
	return j.IsMuted
}

// GetIsMyVideoEnabled returns value of IsMyVideoEnabled field.
func (j *JoinGroupCallRequest) GetIsMyVideoEnabled() (value bool) {
	return j.IsMyVideoEnabled
}

// GetInviteHash returns value of InviteHash field.
func (j *JoinGroupCallRequest) GetInviteHash() (value string) {
	return j.InviteHash
}

// JoinGroupCall invokes method joinGroupCall#c1c947e5 returning error if any.
func (c *Client) JoinGroupCall(ctx context.Context, request *JoinGroupCallRequest) (*Text, error) {
	var result Text

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
