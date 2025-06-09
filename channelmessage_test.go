package disgomsg

import (
	"testing"

	"github.com/bwmarrin/discordgo"
)

func TestNewChannelMessage(t *testing.T) {
	// Test creating a new message with no options
	msg := NewMessage()
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
	msg = NewMessage(WithContent(content))
	if msg.content != content {
		t.Errorf("Expected content %q, got %q", content, msg.content)
	}

	// Test creating a new message with multiple options
	embed := &discordgo.MessageEmbed{Title: "Test Embed"}
	component := discordgo.Button{Label: "Test Button"}
	msg = NewMessage(
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

func TestMessageWithMethods(t *testing.T) {
	// Test WithChannelID
	msg := NewMessage()
	channelID := "123456789"
	msg = msg.WithChannelID(channelID)
	if msg.channelID != channelID {
		t.Errorf("Expected channelID %s, got %s", channelID, msg.channelID)
	}

	// Test WithMessageID
	messageID := "987654321"
	msg = msg.WithMessageID(messageID)
	if msg.messageID != messageID {
		t.Errorf("Expected messageID %s, got %s", messageID, msg.messageID)
	}

	// Test WithContent
	content := "Test content"
	msg = msg.WithContent(content)
	if msg.content != content {
		t.Errorf("Expected content %s, got %s", content, msg.content)
	}

	// Test WithEmbeds
	embeds := []*discordgo.MessageEmbed{{Title: "Test Embed"}}
	msg = msg.WithEmbeds(embeds)
	if len(msg.embeds) != len(embeds) {
		t.Errorf("Expected embeds length %d, got %d", len(embeds), len(msg.embeds))
	} else if msg.embeds[0] != embeds[0] {
		t.Errorf("Expected embed %v, got %v", embeds[0], msg.embeds[0])
	}

	// Test WithComponents
	components := []discordgo.MessageComponent{discordgo.Button{Label: "Test Button"}}
	msg = msg.WithComponents(components)
	if len(msg.components) != len(components) {
		t.Errorf("Expected components length %d, got %d", len(components), len(msg.components))
	} else if msg.components[0] != components[0] {
		t.Errorf("Expected component %v, got %v", components[0], msg.components[0])
	}
}

// Note: We can't easily test Send, Edit, and Delete methods
// without mocking the discordgo.Session interface, which is quite large.
// In a real-world scenario, you might use a mocking library or create a test
// wrapper around the discordgo.Session interface.
