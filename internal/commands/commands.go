package commands

import (
	"encoding/json"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

type Handler func(s *discordgo.Session, i *discordgo.InteractionCreate)
type CmdList struct {
	Commands []*discordgo.ApplicationCommand `json:"commands"`
}

func LoadCommands() ([]*discordgo.ApplicationCommand, error) {
	file, err := os.ReadFile("./config/commands.json")
	if err != nil {
		log.Panicf("Commands could not be loaded in: %v", err)
	}

	var cmdList CmdList
	err = json.Unmarshal(file, &cmdList)
	if err != nil {
		log.Panicf("Command error unmarshaling JSON: %v", err)
	}

	return cmdList.Commands, nil
}

func GetHandlers() map[string]Handler {
	return map[string]Handler{
		"ping": ping,
	}
}
