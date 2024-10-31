package bot

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/theboredkb/bot-template/internal/commands"
)

type Bot struct {
	Session  *discordgo.Session
	Commands []*discordgo.ApplicationCommand
	Handlers map[string]commands.Handler
}

func (b *Bot) RegisterCommands() error {
	s := b.Session
	cmds := b.Commands

	log.Println("Registering commands...")
	for _, v := range cmds {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, "", v)
		if err != nil {
			return fmt.Errorf("Cannot create command %v: %v", v.Name, err)
		}
		log.Printf("Added command %v\n", v.Name)
	}

	return nil
}

func (b *Bot) RegisterHandlers() {
	s := b.Session

	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v", s.State.User.Username)
	})

	handlers := b.Handlers
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := handlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
}

func (b *Bot) PrintCommands() {
	s := b.Session
	regCmds, err := s.ApplicationCommands(s.State.User.ID, "")
	if err != nil {
		log.Printf("Error when reading registered commands %v", err)
	}

	for _, v := range regCmds {
		fmt.Printf("Command: %v", v.Name)
		fmt.Printf("Description: %v", v.Description)
	}
}

func New(token string) (*Bot, error) {
	if token == "" {
		return nil, fmt.Errorf("Token is missing or was not loaded properly.")
	}
	s, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, fmt.Errorf("Invalid bot parameters: %v", err)
	}

	cmds, err := commands.LoadCommands()
	if err != nil {
		return nil, fmt.Errorf("Error encountered while loading commands: %v", err)
	}

	h := commands.GetHandlers()

	return &Bot{
		Session:  s,
		Commands: cmds,
		Handlers: h,
	}, nil
}
