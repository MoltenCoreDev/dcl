package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/MoltenCoreDev/dcl/commands"
	"github.com/bwmarrin/discordgo"
)

var (
	currentGuild   string // TODO: make a use of it
	currentChannel string
	PS1            string
)

func init() {
	// I want to use init method to keep the var cleaner
	PS1 = "dcl >> "
}

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
		commands.DrawPrompt(PS1)
		scanner.Scan()
		args := strings.Split(scanner.Text(), " ")
		cmd := args[0]

		switch cmd {
		case "switch":
			guilds := dg.State.Guilds
			for i, guild := range guilds {
				fmt.Printf("%d. %v \n", i, guild.Name)
			}

			fmt.Println("Select the guild.")

			scanner.Scan()

			c, _ := strconv.Atoi(scanner.Text())

			guild := guilds[c]

			fmt.Printf("Switched to the Guild %v. \nWhat text channel do you want to use? \n", guild.Name)
			currentGuild = guild.ID

			var channels []discordgo.Channel

			for _, channel := range guild.Channels {
				if channel.Type != 0 {
					continue
				} else {
					channels = append(channels, *channel)
					fmt.Printf("%d. %v \n", len(channels), channel.Name)
				}
			}

			scanner.Scan()

			c, _ = strconv.Atoi(scanner.Text())

			currentChannel = channels[c-1].ID

			fmt.Printf("Switched the channel to %v\n", channels[c-1].Name)

		case "send":
			go commands.SendMessage(dg, currentChannel, strings.Join(args[1:], " "))
		case "PS1":
			PS1 = strings.Join(args[1:], " ")
		}

		if cmd == "quit" { // We need to break the for loop not the switch statement
			break
		}
	}

	fmt.Scanln()

}

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.ChannelID == currentChannel && m.Author.ID != s.State.User.ID {
		fmt.Printf("\n%v >> %v \n", m.Author.Username, m.Content)
	}
}
