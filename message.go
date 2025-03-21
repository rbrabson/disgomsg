package disgomsg

import (
	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Message is a Discord message representation used for sending and editing messages in a channel.
type Message struct {
	ID              primitive.ObjectID                `json:"_id,omitempty" bson:"_id,omitempty"`
	MessageID       string                            `json:"message_id,omitempty" bson:"message_id,omitempty"`
	ChannelID       string                            `json:"channel_id,omitempty" bson:"channel_id,omitempty"`
	AllowedMentions *discordgo.MessageAllowedMentions `json:"allowed_mentions,omitempty" bson:"allowed_mentions,omitempty"`
	Components      []discordgo.MessageComponent      `json:"components" bson:"components,omitempty"`
	Content         string                            `json:"content,omitempty" bson:"content,omitempty"`
	Embeds          []*discordgo.MessageEmbed         `json:"embeds" bson:"embeds,omitempty"`
	Files           []*discordgo.File                 `json:"-" bson:"-"`
	Flags           discordgo.MessageFlags            `json:"flags,omitempty" bson:"flags,omitempty"`
	Reference       *discordgo.MessageReference       `json:"message_reference,omitempty" bson:"message_reference,omitempty"`
	StickerIDs      []string                          `json:"sticker_ids" bson:"sticker_ids,omitempty"`
	TTS             bool                              `json:"tts" bson:"tts"`
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
	m.MessageID = sent.ID
	m.ChannelID = channelID

	return nil
}

// SendEphemeral sends the message as an ephemeral message to the specified channel using the provided Discord session.
func (m *Message) SendEphemeral(s *discordgo.Session, channelID string, options ...discordgo.RequestOption) error {
	m.Flags ^= ^discordgo.MessageFlagsEphemeral
	return m.Send(s, channelID, options...)
}

// Edit edits the existing message using the provided Discord session and updates its content, components, embeds, and flags.
func (m *Message) Edit(s *discordgo.Session, channelID string, options ...discordgo.RequestOption) error {
	message := &discordgo.MessageEdit{
		ID:         m.MessageID,
		Channel:    m.ChannelID,
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
func (m *Message) Delete(s *discordgo.Session, channelID string, options ...discordgo.RequestOption) error {
	if m.MessageID == "" {
		return nil // No message to delete
	}
	err := s.ChannelMessageDelete(m.ChannelID, m.MessageID, options...)
	if err != nil {
		return err
	}
	m.MessageID = "" // Clear the ID after deletion
	return nil
}
