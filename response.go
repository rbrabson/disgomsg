package disgomsg

import (
	"errors"

	"github.com/bwmarrin/discordgo"
)

// Response is a Discord interaction response representation used for sending and editing responses to interactions.
type Response message

// NewResponse creates a new message instance with the provided options that may be sent as a response to an interaction.
func NewResponse(opts ...Option) *Response {
	message := newMessage(opts...)
	return (*Response)(message)
}

// SetOptions sets the options for the message.
func (r *Response) SetOptions(opts ...Option) {
	message := (*message)(r)
	message.setOptions(opts...)
}

// Send sends the interaction response to the specified channel using the provided Discord session.
func (r *Response) Send(s *discordgo.Session, i *discordgo.Interaction, options ...discordgo.RequestOption) error {
	var respType discordgo.InteractionResponseType
	if r.responseType == nil {
		respType = discordgo.InteractionResponseChannelMessageWithSource
	} else {
		respType = *r.responseType
	}
	response := &discordgo.InteractionResponse{
		Type: respType,
		Data: &discordgo.InteractionResponseData{
			AllowedMentions: r.allowedMentions,
			Attachments:     &r.attachments,
			Components:      r.components,
			Content:         r.content,
			Embeds:          r.embeds,
			Files:           r.files,
			Flags:           r.flags,
			Choices:         r.choices,
			CustomID:        r.customID,
			Title:           r.title,
		},
	}
	r.interaction = i
	err := s.InteractionRespond(r.interaction, response, options...)
	if err != nil {
		return err
	}

	return nil
}

// SendEphemeral sends the interaction response as an ephemeral message to the specified channel using the provided Discord session.
func (r *Response) SendEphemeral(s *discordgo.Session, i *discordgo.Interaction, options ...discordgo.RequestOption) error {
	r.flags ^= discordgo.MessageFlagsEphemeral
	return r.Send(s, i, options...)
}

// Edit edits the existing interaction response using the provided Discord session and updates its content, components, embeds, and attachments.
func (r *Response) Edit(s *discordgo.Session, options ...discordgo.RequestOption) error {
	if r.interaction == nil {
		return errors.New("missing interaction") // No interaction to delete
	}

	webhookEdit := &discordgo.WebhookEdit{
		Content:         &r.content,
		Components:      &r.components,
		Embeds:          &r.embeds,
		Attachments:     &r.attachments,
		AllowedMentions: r.allowedMentions,
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

// WithInteraction sets the interaction for the response.
func (r *Response) WithInteraction(i *discordgo.Interaction) *Response {
	r.interaction = i
	return r
}
