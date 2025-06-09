package disgomsg

import (
	"testing"

	"github.com/bwmarrin/discordgo"
)

func TestNewMessage(t *testing.T) {
	// Test creating a new message with no options
	msg := newMessage()
	if msg == nil {
		t.Fatal("Expected non-nil message")
	}
	if msg.content != "" {
		t.Errorf("Expected empty content, got %q", msg.content)
	}
	if len(msg.embeds) != 0 {
		t.Errorf("Expected empty embeds, got %d items", len(msg.embeds))
	}
	if len(msg.components) != 0 {
		t.Errorf("Expected empty components, got %d items", len(msg.components))
	}

	// Test creating a new message with content option
	content := "Test message content"
	msg = newMessage(WithContent(content))
	if msg.content != content {
		t.Errorf("Expected content %q, got %q", content, msg.content)
	}

	// Test creating a new message with multiple options
	embed := &discordgo.MessageEmbed{Title: "Test Embed"}
	component := discordgo.Button{Label: "Test Button"}
	msg = newMessage(
		WithContent(content),
		WithEmbeds([]*discordgo.MessageEmbed{embed}),
		WithComponents([]discordgo.MessageComponent{component}),
	)
	if msg.content != content {
		t.Errorf("Expected content %q, got %q", content, msg.content)
	}
	if len(msg.embeds) != 1 {
		t.Errorf("Expected 1 embed, got %d", len(msg.embeds))
	}
	if msg.embeds[0] != embed {
		t.Errorf("Expected embed %v, got %v", embed, msg.embeds[0])
	}
	if len(msg.components) != 1 {
		t.Errorf("Expected 1 component, got %d", len(msg.components))
	}
	if msg.components[0] != component {
		t.Errorf("Expected component %v, got %v", component, msg.components[0])
	}
}

func TestWithAllowedMentions(t *testing.T) {
	allowedMentions := &discordgo.MessageAllowedMentions{
		Parse: []discordgo.AllowedMentionType{discordgo.AllowedMentionTypeUsers},
	}
	msg := newMessage(WithAllowedMentions(allowedMentions))
	if msg.allowedMentions != allowedMentions {
		t.Errorf("Expected allowedMentions %v, got %v", allowedMentions, msg.allowedMentions)
	}
}

func TestWithAttachments(t *testing.T) {
	attachments := []*discordgo.MessageAttachment{
		{ID: "123", URL: "https://example.com/file.png"},
	}
	msg := newMessage(WithAttachments(attachments))
	if len(msg.attachments) != len(attachments) {
		t.Errorf("Expected attachments length %d, got %d", len(attachments), len(msg.attachments))
		return
	}
	for i, att := range attachments {
		if msg.attachments[i] != att {
			t.Errorf("Expected attachment at index %d to be %v, got %v", i, att, msg.attachments[i])
		}
	}
}

func TestWithChannelID(t *testing.T) {
	channelID := "123456789"
	msg := newMessage(WithChannelID(channelID))
	if msg.channelID != channelID {
		t.Errorf("Expected channelID %s, got %s", channelID, msg.channelID)
	}
}

func TestWithChoices(t *testing.T) {
	choices := []*discordgo.ApplicationCommandOptionChoice{
		{Name: "Option 1", Value: "value1"},
	}
	msg := newMessage(WithChoices(choices))
	if len(msg.choices) != len(choices) {
		t.Errorf("Expected choices length %d, got %d", len(choices), len(msg.choices))
		return
	}
	for i, choice := range choices {
		if msg.choices[i] != choice {
			t.Errorf("Expected choice at index %d to be %v, got %v", i, choice, msg.choices[i])
		}
	}
}

func TestWithComponents(t *testing.T) {
	components := []discordgo.MessageComponent{
		discordgo.Button{Label: "Click Me"},
	}
	msg := newMessage(WithComponents(components))
	if len(msg.components) != len(components) {
		t.Errorf("Expected components length %d, got %d", len(components), len(msg.components))
		return
	}
	// Note: For complex types like MessageComponent, a deep comparison might be needed
	// This is a simple check that might not catch all differences
	for i := range components {
		if msg.components[i] != components[i] {
			t.Errorf("Expected component at index %d to match", i)
		}
	}
}

func TestWithContent(t *testing.T) {
	content := "Test content"
	msg := newMessage(WithContent(content))
	if msg.content != content {
		t.Errorf("Expected content %s, got %s", content, msg.content)
	}
}

func TestWithCustomID(t *testing.T) {
	customID := "test-custom-id"
	msg := newMessage(WithCustomID(customID))
	if msg.customID != customID {
		t.Errorf("Expected customID %s, got %s", customID, msg.customID)
	}
}

func TestWithEmbeds(t *testing.T) {
	embeds := []*discordgo.MessageEmbed{
		{Title: "Test Embed", Description: "This is a test embed"},
	}
	msg := newMessage(WithEmbeds(embeds))
	if len(msg.embeds) != len(embeds) {
		t.Errorf("Expected embeds length %d, got %d", len(embeds), len(msg.embeds))
		return
	}
	for i, embed := range embeds {
		if msg.embeds[i] != embed {
			t.Errorf("Expected embed at index %d to be %v, got %v", i, embed, msg.embeds[i])
		}
	}
}

func TestWithFiles(t *testing.T) {
	files := []*discordgo.File{
		{Name: "test.txt", ContentType: "text/plain"},
	}
	msg := newMessage(WithFiles(files))
	if len(msg.files) != len(files) {
		t.Errorf("Expected files length %d, got %d", len(files), len(msg.files))
		return
	}
	for i, file := range files {
		if msg.files[i] != file {
			t.Errorf("Expected file at index %d to be %v, got %v", i, file, msg.files[i])
		}
	}
}

func TestWithFlags(t *testing.T) {
	flags := discordgo.MessageFlagsEphemeral
	msg := newMessage(WithFlags(flags))
	if msg.flags != flags {
		t.Errorf("Expected flags %v, got %v", flags, msg.flags)
	}
}

func TestWithInteraction(t *testing.T) {
	interaction := &discordgo.Interaction{ID: "123456"}
	msg := newMessage(WithInteraction(interaction))
	if msg.interaction != interaction {
		t.Errorf("Expected interaction %v, got %v", interaction, msg.interaction)
	}
}

func TestWithMessageID(t *testing.T) {
	messageID := "987654321"
	msg := newMessage(WithMessageID(messageID))
	if msg.messageID != messageID {
		t.Errorf("Expected messageID %s, got %s", messageID, msg.messageID)
	}
}

func TestWithReference(t *testing.T) {
	reference := &discordgo.MessageReference{MessageID: "123456"}
	msg := newMessage(WithReference(reference))
	if msg.reference != reference {
		t.Errorf("Expected reference %v, got %v", reference, msg.reference)
	}
}

func TestWithResponseType(t *testing.T) {
	responseType := discordgo.InteractionResponseChannelMessageWithSource
	msg := newMessage(WithResponseType(&responseType))
	if msg.responseType != &responseType {
		t.Errorf("Expected responseType %v, got %v", &responseType, msg.responseType)
	}
}

func TestWithStickerIDs(t *testing.T) {
	stickerIDs := []string{"123", "456"}
	msg := newMessage(WithStickerIDs(stickerIDs))
	if len(msg.stickerIDs) != len(stickerIDs) {
		t.Errorf("Expected stickerIDs length %d, got %d", len(stickerIDs), len(msg.stickerIDs))
	}
	for i, id := range stickerIDs {
		if msg.stickerIDs[i] != id {
			t.Errorf("Expected stickerID at index %d to be %s, got %s", i, id, msg.stickerIDs[i])
		}
	}
}

func TestWithTitle(t *testing.T) {
	title := "Test Title"
	msg := newMessage(WithTitle(title))
	if msg.title != title {
		t.Errorf("Expected title %s, got %s", title, msg.title)
	}
}

func TestWithTTS(t *testing.T) {
	msg := newMessage(WithTTS(true))
	if !msg.tts {
		t.Errorf("Expected tts to be true, got false")
	}
}
