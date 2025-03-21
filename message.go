package disgomsg

import (
	"errors"

	"github.com/bwmarrin/discordgo"
)

// Message is a Discord message representation used for sending and editing messages in a channel.
type Message struct {
	messageID       string
	channelID       string
	AllowedMentions *discordgo.MessageAllowedMentions
	Components      []discordgo.MessageComponent
	Content         string
	Embeds          []*discordgo.MessageEmbed
	Files           []*discordgo.File
	Flags           discordgo.MessageFlags
	Reference       *discordgo.MessageReference
	StickerIDs      []string
	TTS             bool
}

// Send s the message to the specified channel using the provided Discord session.
func (m *Message) Send(s *discordgo.Session, channelID string, options ...discordgo.RequestOption) error {
	message := &discordgo.MessageSend{
		AllowedMentions: m.AllowedMentions,
		Components:      m.Components,
		Content:         m.Content,
		Embeds:          m.Embeds,
		Files:           m.Files,
		Flags:           m.Flags,
		Reference:       m.Reference,
		StickerIDs:      m.StickerIDs,
		TTS:             m.TTS,
	}
	sent, err := s.ChannelMessageSendComplex(channelID, message, options...)
	if err != nil {
		return err
	}
	m.messageID = sent.ID
	m.channelID = channelID

	return nil
}

// SendEphemeral sends the message as an ephemeral message to the specified channel using the provided Discord session.
func (m *Message) SendEphemeral(s *discordgo.Session, channelID string, options ...discordgo.RequestOption) error {
	m.Flags ^= discordgo.MessageFlagsEphemeral
	return m.Send(s, channelID, options...)
}

// Edit edits the existing message using the provided Discord session and updates its content, components, embeds, and flags.
func (m *Message) Edit(s *discordgo.Session, options ...discordgo.RequestOption) error {
	if m.channelID == "" {
		return errors.New("missing channelID") // No message to edit
	}
	if m.messageID == "" {
		return errors.New("missing messageID") // No message to edit
	}
	message := &discordgo.MessageEdit{
		ID:         m.messageID,
		Channel:    m.channelID,
		Content:    &m.Content,
		Components: &m.Components,
		Embeds:     &m.Embeds,
		Flags:      m.Flags,
	}
	_, err := s.ChannelMessageEditComplex(message, options...)
	if err != nil {
		return err
	}

	return nil
}

// Delete deletes the message using the provided Discord session and clears the MessageID to indicate it has been deleted.
func (m *Message) Delete(s *discordgo.Session, options ...discordgo.RequestOption) error {
	if m.channelID == "" {
		return errors.New("missing channelID") // No message to edit
	}
	if m.messageID == "" {
		return errors.New("missing messageID") // No message to edit
	}
	err := s.ChannelMessageDelete(m.channelID, m.messageID, options...)
	if err != nil {
		return err
	}
	m.messageID = "" // Clear the ID after deletion
	return nil
}

// WithMessageID sets the message ID that corresponds to the message to be edited or deleted.
func (m *Message) WithMessageID(messageID string) *Message {
	m.messageID = messageID
	return m

}

// WithChannelID sets the channel ID that corresponds to the message to be edited or deleted.
func (m *Message) WithChannelID(channelID string) *Message {
	m.channelID = channelID
	return m
}
