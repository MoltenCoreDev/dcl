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
		fmt.Printf("dcl >> ")
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

			fmt.Printf("Switched to the Guild %v. \nWhat text channel do you want to use? ", guild.Name)
			currentGuild = guild.ID

			for i, channel := range guild.Channels {
				fmt.Printf("%d. %v \n", i, channel.Name)
			}

			scanner.Scan()

			c, _ = strconv.Atoi(scanner.Text())

			currentChannel = guild.Channels[c].ID

			fmt.Printf("Switched the channel to %v\n", guild.Channels[c].Name)

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
