package disgomsg

import (
	"errors"

	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Response is a Discord interaction response representation used for sending and editing responses to interactions.
type Response struct {
	ID              primitive.ObjectID                          `json:"_id,omitempty" bson:"_id,omitempty"`
	Interaction     *discordgo.Interaction                      `json:"interaction" bson:"interaction"`
	Type            discordgo.InteractionResponseType           `json:"type,omitempty" bson:"type,omitempty"`
	AllowedMentions *discordgo.MessageAllowedMentions           `json:"allowed_mentions,omitempty" bson:"allowed_mentions,omitempty"`
	Attachments     []*discordgo.MessageAttachment              `json:"attachments,omitempty" bson:"attachments,omitempty"`
	Components      []discordgo.MessageComponent                `json:"components" bson:"components,omitempty"`
	Content         string                                      `json:"content" bson:"content,omitempty"`
	Embeds          []*discordgo.MessageEmbed                   `json:"embeds" bson:"embeds,omitempty"`
	Files           []*discordgo.File                           `json:"-" bson:"-"`
	TTS             bool                                        `json:"tts" bson:"tts"`
	Choices         []*discordgo.ApplicationCommandOptionChoice `json:"choices,omitempty" bson:"choices,omitempty"`     // Autocomplete interaction only.
	Flags           discordgo.MessageFlags                      `json:"flags,omitempty"`                                // Only MessageFlagsSuppressEmbeds and MessageFlagsEphemeral are valid.
	CustomID        string                                      `json:"custom_id,omitempty" bson:"custom_id,omitempty"` // Modal interaction only.
	Title           string                                      `json:"title,omitempty" bson:"title,omitempty"`         // Modal interaction only.
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
	r.Interaction = i // Store the interaction for future reference

	return nil
}

// SendEphemeral sends the interaction response as an ephemeral message to the specified channel using the provided Discord session.
func (r *Response) SendEphemeral(s *discordgo.Session, i *discordgo.Interaction, options ...discordgo.RequestOption) error {
	r.Flags ^= ^discordgo.MessageFlagsEphemeral
	return r.Send(s, i, options...)
}

// Edit edits the existing interaction response using the provided Discord session and updates its content, components, embeds, and attachments.
func (r *Response) Edit(s *discordgo.Session, options ...discordgo.RequestOption) error {
	if r.Interaction == nil {
		return errors.New("missing interaction") // No interaction to delete
	}

	webhookEdit := &discordgo.WebhookEdit{
		Content:         &r.Content,
		Components:      &r.Components,
		Embeds:          &r.Embeds,
		Attachments:     &r.Attachments,
		AllowedMentions: r.AllowedMentions,
	}
	_, err := s.InteractionResponseEdit(r.Interaction, webhookEdit, options...)
	if err != nil {
		return err
	}

	return nil
}

// Delete deletes the interaction response using the provided Discord session.
func (r *Response) Delete(s *discordgo.Session, options ...discordgo.RequestOption) error {
	if r.Interaction == nil {
		return errors.New("missing interaction") // No interaction to delete
	}
	err := s.InteractionResponseDelete(r.Interaction, options...)
	if err != nil {
		return err
	}

	return nil
}
