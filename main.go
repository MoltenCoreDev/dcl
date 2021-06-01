package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/MoltenCoreDev/dcl/commands"
	"github.com/bwmarrin/discordgo"
)

var (
	currentGuild   string
	currentChannel string
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	token := os.Getenv("dcl_token")
	dg, err := discordgo.New(token)
	if err != nil {
		fmt.Println(err)
		return
	}

	dg.AddHandler(MessageCreate)

	dg.Identify.Intents = discordgo.IntentsAll

	err = dg.Open()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer dg.Close()

	for {
		fmt.Printf("dcl >>")
		scanner.Scan()
		args := strings.Split(scanner.Text(), " ")
		cmd := args[0]

		switch cmd {
		case "switchGuild":
			commands.ChangeGuild(&currentGuild, args[1])
		case "switchChannel":
			commands.ChangeChannel(&currentChannel, args[1])
		case "send":
			commands.SendMessage(dg, currentChannel, strings.Join(args[1:], " "))
		}

		if cmd == "quit" { // We need to break the for loop not the switch statement
			break
		}
	}

	fmt.Scanln()

}

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
}
