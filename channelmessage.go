package disgomsg

import (
	"github.com/bwmarrin/discordgo"
)

// Message is a Discord message representation used for sending and editing messages in a channel.
type Message message

// NewMessage creates a new message instance with the provided options that may be sent to a channel.
func NewMessage(opts ...Option) *Message {
	message := newMessage(opts...)
	return (*Message)(message)
}

// SetOptions sets the options for the message.
func (m *Message) SetOptions(opts ...Option) {
	message := (*message)(m)
	message.setOptions(opts...)
}

// Send s the message to the specified channel using the provided Discord session.
func (m *Message) Send(s *discordgo.Session, options ...discordgo.RequestOption) error {
	message := &discordgo.MessageSend{
		AllowedMentions: m.allowedMentions,
		Components:      m.components,
		Content:         m.content,
		Embeds:          m.embeds,
		Files:           m.files,
		Flags:           m.flags,
		Reference:       m.reference,
		StickerIDs:      m.stickerIDs,
		TTS:             m.tts,
	}
	sent, err := s.ChannelMessageSendComplex(m.channelID, message, options...)
	if err != nil {
		return err
	}
	m.messageID = sent.ID

	return nil
}

// Edit edits the existing message using the provided Discord session and updates its content, components, embeds, and flags.
func (m *Message) Edit(s *discordgo.Session, options ...discordgo.RequestOption) error {
	if m.channelID == "" {
		return ErrMissingChannelID
	}
	if m.messageID == "" {
		return ErrMissingMessageID
	}
	message := &discordgo.MessageEdit{
		ID:         m.messageID,
		Channel:    m.channelID,
		Content:    &m.content,
		Components: &m.components,
		Embeds:     &m.embeds,
		Flags:      m.flags,
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
		return ErrMissingChannelID
	}
	if m.messageID == "" {
		return ErrMissingMessageID
	}
	err := s.ChannelMessageDelete(m.channelID, m.messageID, options...)
	if err != nil {
		return err
	}
	m.messageID = "" // Clear the ID after deletion
	return nil
}
