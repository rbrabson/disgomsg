package disgomsg

import (
	"testing"

	"github.com/bwmarrin/discordgo"
)

func TestNewDirectMessage(t *testing.T) {
	// Test creating a new direct message with no options
	dm := NewDirectMessage()
	if dm == nil {
		t.Fatal("Expected non-nil direct message")
	}
	if dm.content != "" {
		t.Errorf("Expected empty content, got %q", dm.content)
	}
	if len(dm.embeds) != 0 {
		t.Errorf("Expected empty embeds, got %d items", len(dm.embeds))
	}
	if len(dm.components) != 0 {
		t.Errorf("Expected empty components, got %d items", len(dm.components))
	}

	// Test creating a new direct message with content option
	content := "Test direct message content"
	dm = NewDirectMessage(WithContent(content))
	if dm.content != content {
		t.Errorf("Expected content %q, got %q", content, dm.content)
	}

	// Test creating a new direct message with multiple options
	embed := &discordgo.MessageEmbed{Title: "Test Embed"}
	component := discordgo.Button{Label: "Test Button"}
	dm = NewDirectMessage(
		WithContent(content),
		WithEmbeds([]*discordgo.MessageEmbed{embed}),
		WithComponents([]discordgo.MessageComponent{component}),
	)
	if dm.content != content {
		t.Errorf("Expected content %q, got %q", content, dm.content)
	}
	if len(dm.embeds) != 1 {
		t.Errorf("Expected 1 embed, got %d", len(dm.embeds))
	}
	if dm.embeds[0] != embed {
		t.Errorf("Expected embed %v, got %v", embed, dm.embeds[0])
	}
	if len(dm.components) != 1 {
		t.Errorf("Expected 1 component, got %d", len(dm.components))
	}
	if dm.components[0] != component {
		t.Errorf("Expected component %v, got %v", component, dm.components[0])
	}
}

func TestDirectMessageWithMethods(t *testing.T) {
	// Test WithMessageID
	dm := NewDirectMessage()
	channelID := "123456789"
	dm = dm.WithMessageID(channelID)
	if dm.channelID != channelID {
		t.Errorf("Expected channelID %s, got %s", channelID, dm.channelID)
	}

	// Test WithContent
	content := "Test content"
	dm = dm.WithContent(content)
	if dm.content != content {
		t.Errorf("Expected content %s, got %s", content, dm.content)
	}

	// Test WithEmbeds
	embeds := []*discordgo.MessageEmbed{{Title: "Test Embed"}}
	dm = dm.WithEmbeds(embeds)
	if len(dm.embeds) != len(embeds) {
		t.Errorf("Expected embeds length %d, got %d", len(embeds), len(dm.embeds))
	} else if dm.embeds[0] != embeds[0] {
		t.Errorf("Expected embed %v, got %v", embeds[0], dm.embeds[0])
	}

	// Test WithComponents
	components := []discordgo.MessageComponent{discordgo.Button{Label: "Test Button"}}
	dm = dm.WithComponents(components)
	if len(dm.components) != len(components) {
		t.Errorf("Expected components length %d, got %d", len(components), len(dm.components))
	} else if dm.components[0] != components[0] {
		t.Errorf("Expected component %v, got %v", components[0], dm.components[0])
	}
}

// Note: We can't easily test Send, Edit, and Delete methods
// without mocking the discordgo.Session interface, which is quite large.
// In a real-world scenario, you might use a mocking library or create a test
// wrapper around the discordgo.Session interface.

// Note: We also can't easily test WithMemberID method since it requires
// a real discordgo.Session to create a user channel.