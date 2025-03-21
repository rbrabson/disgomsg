package disgomsg

import (
	"errors"

	"github.com/bwmarrin/discordgo"
)

// Response is a Discord interaction response representation used for sending and editing responses to interactions.
type Response struct {
	interaction     *discordgo.Interaction
	Type            discordgo.InteractionResponseType
	AllowedMentions *discordgo.MessageAllowedMentions
	Attachments     []*discordgo.MessageAttachment
	Components      []discordgo.MessageComponent
	Content         string
	Embeds          []*discordgo.MessageEmbed
	Files           []*discordgo.File
	TTS             bool
	Choices         []*discordgo.ApplicationCommandOptionChoice // Autocomplete interaction only.
	Flags           discordgo.MessageFlags                      // Only MessageFlagsSuppressEmbeds and MessageFlagsEphemeral are valid.
	CustomID        string                                      // Modal interaction only.
	Title           string                                      // Modal interaction only.
}

// Send sends the interaction response to the specified channel using the provided Discord session.
func (r *Response) Send(s *discordgo.Session, i *discordgo.Interaction, options ...discordgo.RequestOption) error {
	response := &discordgo.InteractionResponse{
		Type: r.Type,
		Data: &discordgo.InteractionResponseData{
			AllowedMentions: r.AllowedMentions,
			Attachments:     &r.Attachments,
			Components:      r.Components,
			Content:         r.Content,
			Embeds:          r.Embeds,
			Files:           r.Files,
			Flags:           r.Flags,
			Choices:         r.Choices,
			CustomID:        r.CustomID,
			Title:           r.Title,
		},
	}
	err := s.InteractionRespond(i, response, options...)
	if err != nil {
		return err
	}
	r.interaction = i // Store the interaction for future reference

	return nil
}

// SendEphemeral sends the interaction response as an ephemeral message to the specified channel using the provided Discord session.
func (r *Response) SendEphemeral(s *discordgo.Session, i *discordgo.Interaction, options ...discordgo.RequestOption) error {
	r.Flags ^= discordgo.MessageFlagsEphemeral
	return r.Send(s, i, options...)
}

// Edit edits the existing interaction response using the provided Discord session and updates its content, components, embeds, and attachments.
func (r *Response) Edit(s *discordgo.Session, options ...discordgo.RequestOption) error {
	if r.interaction == nil {
		return errors.New("missing interaction") // No interaction to delete
	}

	webhookEdit := &discordgo.WebhookEdit{
		Content:         &r.Content,
		Components:      &r.Components,
		Embeds:          &r.Embeds,
		Attachments:     &r.Attachments,
		AllowedMentions: r.AllowedMentions,
	}
	_, err := s.InteractionResponseEdit(r.interaction, webhookEdit, options...)
	if err != nil {
		return err
	}

	return nil
}

// Delete deletes the interaction response using the provided Discord session.
func (r *Response) Delete(s *discordgo.Session, options ...discordgo.RequestOption) error {
	if r.interaction == nil {
		return errors.New("missing interaction") // No interaction to delete
	}
	err := s.InteractionResponseDelete(r.interaction, options...)
	if err != nil {
		return err
	}

	return nil
}

// WithInteraction sets the interaction that corresponds to the response to be edited or deleted.
func (r *Response) WithInteraction(i *discordgo.Interaction) *Response {
	r.interaction = i
	return r
}
