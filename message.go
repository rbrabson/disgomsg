package disgomsg

import "github.com/bwmarrin/discordgo"

// message is the common struct for all direct messages, channel messages and responses
type message struct {
	allowedMentions *discordgo.MessageAllowedMentions
	attachments     []*discordgo.MessageAttachment
	channelID       string
	choices         []*discordgo.ApplicationCommandOptionChoice // Autocomplete interaction only.
	components      []discordgo.MessageComponent
	content         string
	customID        string // Modal interaction only.
	embeds          []*discordgo.MessageEmbed
	files           []*discordgo.File
	flags           discordgo.MessageFlags // Only MessageFlagsSuppressEmbeds and MessageFlagsEphemeral are valid.
	interaction     *discordgo.Interaction
	messageID       string
	reference       *discordgo.MessageReference
	responseType    *discordgo.InteractionResponseType
	stickerIDs      []string
	title           string
	tts             bool
}

// newMessage creates a new message with the given options
func newMessage(opts ...Option) *message {
	f := &message{}
	for _, opt := range opts {
		opt(f)
	}
	return f
}

// Option is a function that modifies a message.
type Option func(*message)

// WithAllowedMentions sets the allowed mentions for the message.
func WithAllowedMentions(allowedMentions *discordgo.MessageAllowedMentions) Option {
	return func(f *message) {
		f.allowedMentions = allowedMentions
	}
}

// WithAttachments sets the attachments for the message.
func WithAttachments(attachments []*discordgo.MessageAttachment) Option {
	return func(f *message) {
		f.attachments = attachments
	}
}

// WithChannelID sets the channel ID for the message.
func WithChannelID(channelID string) Option {
	return func(f *message) {
		f.channelID = channelID
	}
}

// WithChoices sets the choices for the message.
func WithChoices(choices []*discordgo.ApplicationCommandOptionChoice) Option {
	return func(f *message) {
		f.choices = choices
	}
}

// WithComponents sets the components for the message.
func WithComponents(components []discordgo.MessageComponent) Option {
	return func(f *message) {
		f.components = components
	}
}

// WithContent sets the content for the message.
func WithContent(content string) Option {
	return func(f *message) {
		f.content = content
	}
}

// WithCustomID sets the custom ID for the message.
func WithCustomID(customID string) Option {
	return func(f *message) {
		f.customID = customID
	}
}

// WithEmbeds sets the embeds for the message.
func WithEmbeds(embeds []*discordgo.MessageEmbed) Option {
	return func(f *message) {
		f.embeds = embeds
	}
}

// WithFiles sets the files for the message.
func WithFiles(files []*discordgo.File) Option {
	return func(f *message) {
		f.files = files
	}
}

// WithFlags sets the flags for the message.
func WithFlags(flags discordgo.MessageFlags) Option {
	return func(f *message) {
		f.flags = flags
	}
}

// WithInteraction sets the interaction for the message.
func WithInteraction(interaction *discordgo.Interaction) Option {
	return func(f *message) {
		f.interaction = interaction
	}
}

// WithMessageID sets the message ID for the message.
func WithMessageID(messageID string) Option {
	return func(f *message) {
		f.messageID = messageID
	}
}

// WithReference sets the reference for the message.
func WithReference(reference *discordgo.MessageReference) Option {
	return func(f *message) {
		f.reference = reference
	}
}

// WithResponseType sets the response type for the message.
func WithResponseType(responseType *discordgo.InteractionResponseType) Option {
	return func(f *message) {
		f.responseType = responseType
	}
}

// WithStickerIDs sets the sticker IDs for the message.
func WithStickerIDs(stickerIDs []string) Option {
	return func(f *message) {
		f.stickerIDs = stickerIDs
	}
}

// WithTitle sets the title for the message.
func WithTitle(title string) Option {
	return func(f *message) {
		f.title = title
	}
}

// WithTTS sets the tts for the message.
func WithTTS(tts bool) Option {
	return func(f *message) {
		f.tts = tts
	}
}
