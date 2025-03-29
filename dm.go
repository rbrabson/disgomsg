package disgomsg

import (
	"errors"

	"github.com/bwmarrin/discordgo"
)

// DirectMessage is a Discord message representation used for sending a message to a user in a private channel.
type DirectMessage struct {
	AllowedMentions *discordgo.MessageAllowedMentions
	Components      []discordgo.MessageComponent
	Content         string
	Embeds          []*discordgo.MessageEmbed
	Files           []*discordgo.File
	Flags           discordgo.MessageFlags
	Reference       *discordgo.MessageReference
	StickerIDs      []string
	TTS             bool
	messageID       string
	channelID       string
}

// Send sends a direct message to the specified member using the provided Discord session.
func (dm *DirectMessage) Send(s *discordgo.Session, memberID string, options ...discordgo.RequestOption) error {
	channel, err := s.UserChannelCreate(memberID)
	if err != nil {
		return err
	}

	message := &discordgo.MessageSend{
		AllowedMentions: dm.AllowedMentions,
		Components:      dm.Components,
		Content:         dm.Content,
		Embeds:          dm.Embeds,
		Files:           dm.Files,
		Flags:           dm.Flags,
		Reference:       dm.Reference,
		StickerIDs:      dm.StickerIDs,
		TTS:             dm.TTS,
	}

	sent, err := s.ChannelMessageSendComplex(channel.ID, message, options...)
	if err != nil {
		return err
	}
	dm.messageID = sent.ID
	dm.channelID = channel.ID

	return nil
}

// Edit edits the existing message using the provided Discord session and updates its content, components, embeds, and flags.
func (dm *DirectMessage) Edit(s *discordgo.Session, options ...discordgo.RequestOption) error {
	if dm.channelID == "" {
		return errors.New("missing channelID") // No message to edit
	}
	if dm.messageID == "" {
		return errors.New("missing messageID") // No message to edit
	}
	message := &discordgo.MessageEdit{
		ID:         dm.messageID,
		Channel:    dm.channelID,
		Content:    &dm.Content,
		Components: &dm.Components,
		Embeds:     &dm.Embeds,
		Flags:      dm.Flags,
	}
	_, err := s.ChannelMessageEditComplex(message, options...)
	if err != nil {
		return err
	}

	return nil
}

// Delete deletes the message using the provided Discord session and clears the MessageID to indicate it has been deleted.
func (dm *DirectMessage) Delete(s *discordgo.Session, options ...discordgo.RequestOption) error {
	if dm.channelID == "" {
		return errors.New("missing channelID") // No message to edit
	}
	if dm.messageID == "" {
		return errors.New("missing messageID") // No message to edit
	}
	err := s.ChannelMessageDelete(dm.channelID, dm.messageID, options...)
	if err != nil {
		return err
	}
	dm.messageID = "" // Clear the ID after deletion
	return nil
}

// WithMessageID sets the message ID that corresponds to the direct message to be edited or deleted.
func (dm *DirectMessage) WithMessageID(messageID string) *DirectMessage {
	dm.messageID = messageID
	return dm

}

// WithChannelID sets the channel ID that corresponds to the direct message to be edited or deleted.
func (dm *DirectMessage) WithChannelID(channelID string) *DirectMessage {
	dm.channelID = channelID
	return dm
}
