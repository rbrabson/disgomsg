package disgomsg

import (
	"github.com/bwmarrin/discordgo"
)

// DirectMessage is a Discord message representation used for sending a message to a user in a private channel.
type DirectMessage message

// NewDirectMessage creates a new message instance with the provided options that may be sent as a direct message to
// a given member.
func NewDirectMessage(opts ...Option) *DirectMessage {
	message := newMessage(opts...)
	return (*DirectMessage)(message)
}

// Send sends a direct message to the specified member using the provided Discord session.
func (dm *DirectMessage) Send(s *discordgo.Session, memberID string, options ...discordgo.RequestOption) (messagID string, err error) {
	channel, err := s.UserChannelCreate(memberID)
	if err != nil {
		return "", err
	}

	message := &discordgo.MessageSend{
		AllowedMentions: dm.allowedMentions,
		Components:      dm.components,
		Content:         dm.content,
		Embeds:          dm.embeds,
		Files:           dm.files,
		Flags:           dm.flags,
		Reference:       dm.reference,
		StickerIDs:      dm.stickerIDs,
		TTS:             dm.tts,
	}

	sent, err := s.ChannelMessageSendComplex(channel.ID, message, options...)
	if err != nil {
		return "", err
	}
	dm.messageID = sent.ID
	dm.channelID = channel.ID

	return dm.messageID, nil
}

// Edit edits the existing message using the provided Discord session and updates its content, components, embeds, and flags.
func (dm *DirectMessage) Edit(s *discordgo.Session, options ...discordgo.RequestOption) error {
	if dm.channelID == "" {
		return ErrMissingChannelID
	}
	if dm.messageID == "" {
		return ErrMissingMessageID
	}
	message := &discordgo.MessageEdit{
		ID:         dm.messageID,
		Channel:    dm.channelID,
		Content:    &dm.content,
		Components: &dm.components,
		Embeds:     &dm.embeds,
		Flags:      dm.flags,
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
		return ErrMissingChannelID
	}
	if dm.messageID == "" {
		return ErrMissingMessageID
	}
	err := s.ChannelMessageDelete(dm.channelID, dm.messageID, options...)
	if err != nil {
		return err
	}
	dm.messageID = "" // Clear the ID after deletion
	return nil
}

// WithMemberID uses the member ID to create a new channel to the member and sets the channel ID
// for the message.
func (dm *DirectMessage) WithMemberID(s *discordgo.Session, memberID string) *DirectMessage {
	channel, err := s.UserChannelCreate(memberID)
	if err == nil {
		dm.channelID = channel.ID
	}
	return dm
}

// WithMessageID sets the message ID for the message.
func (dm *DirectMessage) WithMessageID(channelID string) *DirectMessage {
	dm.channelID = channelID
	return dm
}

// WtthContent sets the content for the message.
func (dm *DirectMessage) WithContent(content string) *DirectMessage {
	dm.content = content
	return dm
}

// WithEmbeds sets the embeds for the message.
func (dm *DirectMessage) WithEmbeds(embeds []*discordgo.MessageEmbed) *DirectMessage {
	dm.embeds = embeds
	return dm
}

// WithComponents sets the components for the message.
func (dm *DirectMessage) WithComponents(components []discordgo.MessageComponent) *DirectMessage {
	dm.components = components
	return dm
}
