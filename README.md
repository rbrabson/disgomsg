# DisGoMsg

DisGoMsg is a Go package that provides a clean and flexible API for creating, sending, and managing Discord messages using the [discordgo](https://github.com/bwmarrin/discordgo) library. It simplifies the process of working with different types of Discord messages including channel messages, direct messages, and interaction responses.

## Features

- Functional options pattern for flexible message creation
- Fluent interface for message configuration
- Support for all Discord message features:
  - Text content
  - Embeds
  - Components (buttons, select menus, etc.)
  - Files and attachments
  - Message flags
  - Allowed mentions
- Specialized message types:
  - Channel messages
  - Direct messages
  - Interaction responses

## Installation

```bash
go get github.com/rbrabson/disgomsg
```

## Usage

### Creating and Sending a Channel Message

```go
package main

import (
    "github.com/bwmarrin/discordgo"
    "github.com/rbrabson/disgomsg"
)

func main() {
    // Initialize Discord session
    session, _ := discordgo.New("Bot YOUR_BOT_TOKEN")
    
    // Create a new message with content and an embed
    msg := disgomsg.NewMessage(
        disgomsg.WithContent("Hello, Discord!"),
        disgomsg.WithEmbeds([]*discordgo.MessageEmbed{
            {
                Title:       "Example Embed",
                Description: "This is an example embed",
                Color:       0x00ff00,
            },
        }),
    )
    
    // Send the message to a channel
    channelID := "YOUR_CHANNEL_ID"
    _, err := msg.WithChannelID(channelID).Send(session)
    if err != nil {
        // Handle error
    }
}
```

### Creating and Sending an Interaction Response

```go
func handleInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
    // Create a response with content and a button
    response := disgomsg.NewResponse(
        disgomsg.WithContent("Here's a button!"),
        disgomsg.WithComponents([]discordgo.MessageComponent{
            discordgo.Button{
                Label:    "Click Me",
                Style:    discordgo.PrimaryButton,
                CustomID: "button_click",
            },
        }),
    )
    
    // Send the response
    err := response.Send(s, i.Interaction)
    if err != nil {
        // Handle error
    }
}
```

### Editing a Message

```go
// Edit an existing message
msg := disgomsg.NewMessage(
    disgomsg.WithContent("Updated content"),
)
_, err := msg.WithChannelID(channelID).WithMessageID(messageID).Edit(session)
```

## License

This project is licensed under the GNU General Public License v3.0 - see the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.