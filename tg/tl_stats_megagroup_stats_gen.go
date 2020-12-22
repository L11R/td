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

// StatsMegagroupStats represents TL type `stats.megagroupStats#ef7ff916`.
// Supergroup statistics¹
//
// Links:
//  1) https://core.telegram.org/api/stats
//
// See https://core.telegram.org/constructor/stats.megagroupStats for reference.
type StatsMegagroupStats struct {
	// Period in consideration
	Period StatsDateRangeDays
	// Member count change for period in consideration
	Members StatsAbsValueAndPrev
	// Message number change for period in consideration
	Messages StatsAbsValueAndPrev
	// Number of users that viewed messages, for range in consideration
	Viewers StatsAbsValueAndPrev
	// Number of users that posted messages, for range in consideration
	Posters StatsAbsValueAndPrev
	// Supergroup growth graph (absolute subscriber count)
	GrowthGraph StatsGraphClass
	// Members growth (relative subscriber count)
	MembersGraph StatsGraphClass
	// New members by source graph
	NewMembersBySourceGraph StatsGraphClass
	// Subscriber language graph (piechart)
	LanguagesGraph StatsGraphClass
	// Message activity graph (stacked bar graph, message type)
	MessagesGraph StatsGraphClass
	// Group activity graph (deleted, modified messages, blocked users)
	ActionsGraph StatsGraphClass
	// Activity per hour graph (absolute)
	TopHoursGraph StatsGraphClass
	// Activity per day of week graph (absolute)
	WeekdaysGraph StatsGraphClass
	// Info about most active group members
	TopPosters []StatsGroupTopPoster
	// Info about most active group admins
	TopAdmins []StatsGroupTopAdmin
	// Info about most active group inviters
	TopInviters []StatsGroupTopInviter
	// Info about users mentioned in statistics
	Users []UserClass
}

// StatsMegagroupStatsTypeID is TL type id of StatsMegagroupStats.
const StatsMegagroupStatsTypeID = 0xef7ff916

// String implements fmt.Stringer.
func (m *StatsMegagroupStats) String() string {
	if m == nil {
		return "StatsMegagroupStats(nil)"
	}
	var sb strings.Builder
	sb.WriteString("StatsMegagroupStats")
	sb.WriteString("{\n")
	sb.WriteString("\tPeriod: ")
	sb.WriteString(m.Period.String())
	sb.WriteString(",\n")
	sb.WriteString("\tMembers: ")
	sb.WriteString(m.Members.String())
	sb.WriteString(",\n")
	sb.WriteString("\tMessages: ")
	sb.WriteString(m.Messages.String())
	sb.WriteString(",\n")
	sb.WriteString("\tViewers: ")
	sb.WriteString(m.Viewers.String())
	sb.WriteString(",\n")
	sb.WriteString("\tPosters: ")
	sb.WriteString(m.Posters.String())
	sb.WriteString(",\n")
	sb.WriteString("\tGrowthGraph: ")
	sb.WriteString(m.GrowthGraph.String())
	sb.WriteString(",\n")
	sb.WriteString("\tMembersGraph: ")
	sb.WriteString(m.MembersGraph.String())
	sb.WriteString(",\n")
	sb.WriteString("\tNewMembersBySourceGraph: ")
	sb.WriteString(m.NewMembersBySourceGraph.String())
	sb.WriteString(",\n")
	sb.WriteString("\tLanguagesGraph: ")
	sb.WriteString(m.LanguagesGraph.String())
	sb.WriteString(",\n")
	sb.WriteString("\tMessagesGraph: ")
	sb.WriteString(m.MessagesGraph.String())
	sb.WriteString(",\n")
	sb.WriteString("\tActionsGraph: ")
	sb.WriteString(m.ActionsGraph.String())
	sb.WriteString(",\n")
	sb.WriteString("\tTopHoursGraph: ")
	sb.WriteString(m.TopHoursGraph.String())
	sb.WriteString(",\n")
	sb.WriteString("\tWeekdaysGraph: ")
	sb.WriteString(m.WeekdaysGraph.String())
	sb.WriteString(",\n")
	sb.WriteByte('[')
	for _, v := range m.TopPosters {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteByte('[')
	for _, v := range m.TopAdmins {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteByte('[')
	for _, v := range m.TopInviters {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteByte('[')
	for _, v := range m.Users {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (m *StatsMegagroupStats) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode stats.megagroupStats#ef7ff916 as nil")
	}
	b.PutID(StatsMegagroupStatsTypeID)
	if err := m.Period.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field period: %w", err)
	}
	if err := m.Members.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field members: %w", err)
	}
	if err := m.Messages.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field messages: %w", err)
	}
	if err := m.Viewers.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field viewers: %w", err)
	}
	if err := m.Posters.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field posters: %w", err)
	}
	if m.GrowthGraph == nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field growth_graph is nil")
	}
	if err := m.GrowthGraph.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field growth_graph: %w", err)
	}
	if m.MembersGraph == nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field members_graph is nil")
	}
	if err := m.MembersGraph.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field members_graph: %w", err)
	}
	if m.NewMembersBySourceGraph == nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field new_members_by_source_graph is nil")
	}
	if err := m.NewMembersBySourceGraph.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field new_members_by_source_graph: %w", err)
	}
	if m.LanguagesGraph == nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field languages_graph is nil")
	}
	if err := m.LanguagesGraph.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field languages_graph: %w", err)
	}
	if m.MessagesGraph == nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field messages_graph is nil")
	}
	if err := m.MessagesGraph.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field messages_graph: %w", err)
	}
	if m.ActionsGraph == nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field actions_graph is nil")
	}
	if err := m.ActionsGraph.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field actions_graph: %w", err)
	}
	if m.TopHoursGraph == nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field top_hours_graph is nil")
	}
	if err := m.TopHoursGraph.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field top_hours_graph: %w", err)
	}
	if m.WeekdaysGraph == nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field weekdays_graph is nil")
	}
	if err := m.WeekdaysGraph.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field weekdays_graph: %w", err)
	}
	b.PutVectorHeader(len(m.TopPosters))
	for idx, v := range m.TopPosters {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field top_posters element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(m.TopAdmins))
	for idx, v := range m.TopAdmins {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field top_admins element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(m.TopInviters))
	for idx, v := range m.TopInviters {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field top_inviters element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(m.Users))
	for idx, v := range m.Users {
		if v == nil {
			return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (m *StatsMegagroupStats) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode stats.megagroupStats#ef7ff916 to nil")
	}
	if err := b.ConsumeID(StatsMegagroupStatsTypeID); err != nil {
		return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: %w", err)
	}
	{
		if err := m.Period.Decode(b); err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field period: %w", err)
		}
	}
	{
		if err := m.Members.Decode(b); err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field members: %w", err)
		}
	}
	{
		if err := m.Messages.Decode(b); err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field messages: %w", err)
		}
	}
	{
		if err := m.Viewers.Decode(b); err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field viewers: %w", err)
		}
	}
	{
		if err := m.Posters.Decode(b); err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field posters: %w", err)
		}
	}
	{
		value, err := DecodeStatsGraph(b)
		if err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field growth_graph: %w", err)
		}
		m.GrowthGraph = value
	}
	{
		value, err := DecodeStatsGraph(b)
		if err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field members_graph: %w", err)
		}
		m.MembersGraph = value
	}
	{
		value, err := DecodeStatsGraph(b)
		if err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field new_members_by_source_graph: %w", err)
		}
		m.NewMembersBySourceGraph = value
	}
	{
		value, err := DecodeStatsGraph(b)
		if err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field languages_graph: %w", err)
		}
		m.LanguagesGraph = value
	}
	{
		value, err := DecodeStatsGraph(b)
		if err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field messages_graph: %w", err)
		}
		m.MessagesGraph = value
	}
	{
		value, err := DecodeStatsGraph(b)
		if err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field actions_graph: %w", err)
		}
		m.ActionsGraph = value
	}
	{
		value, err := DecodeStatsGraph(b)
		if err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field top_hours_graph: %w", err)
		}
		m.TopHoursGraph = value
	}
	{
		value, err := DecodeStatsGraph(b)
		if err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field weekdays_graph: %w", err)
		}
		m.WeekdaysGraph = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field top_posters: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value StatsGroupTopPoster
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field top_posters: %w", err)
			}
			m.TopPosters = append(m.TopPosters, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field top_admins: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value StatsGroupTopAdmin
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field top_admins: %w", err)
			}
			m.TopAdmins = append(m.TopAdmins, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field top_inviters: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value StatsGroupTopInviter
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field top_inviters: %w", err)
			}
			m.TopInviters = append(m.TopInviters, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field users: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field users: %w", err)
			}
			m.Users = append(m.Users, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for StatsMegagroupStats.
var (
	_ bin.Encoder = &StatsMegagroupStats{}
	_ bin.Decoder = &StatsMegagroupStats{}
)
