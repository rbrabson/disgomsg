package disgomsg

import (
	"testing"

	"github.com/bwmarrin/discordgo"
)

func TestNewResponse(t *testing.T) {
	// Test creating a new response with no options
	resp := NewResponse()
	if resp == nil {
		t.Fatal("Expected non-nil response")
	}
	if resp.content != "" {
		t.Errorf("Expected empty content, got %q", resp.content)
	}

	// Test creating a new response with content option
	content := "Test response content"
	resp = NewResponse(WithContent(content))
	if resp.content != content {
		t.Errorf("Expected content %q, got %q", content, resp.content)
	}

	// Test creating a new response with multiple options
	embed := &discordgo.MessageEmbed{Title: "Test Embed"}
	component := discordgo.Button{Label: "Test Button"}
	resp = NewResponse(
		WithContent(content),
		WithEmbeds([]*discordgo.MessageEmbed{embed}),
		WithComponents([]discordgo.MessageComponent{component}),
	)
	if resp.content != content {
		t.Errorf("Expected content %q, got %q", content, resp.content)
	}
	if len(resp.embeds) != 1 {
		t.Errorf("Expected 1 embed, got %d", len(resp.embeds))
	}
	if resp.embeds[0] != embed {
		t.Errorf("Expected embed %v, got %v", embed, resp.embeds[0])
	}
	if len(resp.components) != 1 {
		t.Errorf("Expected 1 component, got %d", len(resp.components))
	}
	if resp.components[0] != component {
		t.Errorf("Expected component %v, got %v", component, resp.components[0])
	}
}

func TestResponseWithMethods(t *testing.T) {
	// Test WithInteraction
	resp := NewResponse()
	interaction := &discordgo.Interaction{ID: "123456"}
	resp = resp.WithInteraction(interaction)
	if resp.interaction != interaction {
		t.Errorf("Expected interaction %v, got %v", interaction, resp.interaction)
	}

	// Test WithContent
	content := "Test content"
	resp = resp.WithContent(content)
	if resp.content != content {
		t.Errorf("Expected content %s, got %s", content, resp.content)
	}

	// Test WithEmbeds
	embeds := []*discordgo.MessageEmbed{{Title: "Test Embed"}}
	resp = resp.WithEmbeds(embeds)
	if len(resp.embeds) != len(embeds) {
		t.Errorf("Expected embeds length %d, got %d", len(embeds), len(resp.embeds))
	} else if resp.embeds[0] != embeds[0] {
		t.Errorf("Expected embed %v, got %v", embeds[0], resp.embeds[0])
	}

	// Test WithComponents
	components := []discordgo.MessageComponent{discordgo.Button{Label: "Test Button"}}
	resp = resp.WithComponents(components)
	if len(resp.components) != len(components) {
		t.Errorf("Expected components length %d, got %d", len(components), len(resp.components))
	} else if resp.components[0] != components[0] {
		t.Errorf("Expected component %v, got %v", components[0], resp.components[0])
	}
}

// Note: We can't easily test Send, SendEphemeral, Edit, and Delete methods
// without mocking the discordgo.Session interface, which is quite large.
// In a real-world scenario, you might use a mocking library or create a test
// wrapper around the discordgo.Session interface.
