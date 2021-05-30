package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

var (
  content string
  guildID string
)

func main() {
  scanner := bufio.NewScanner(os.Stdin)
  token := ""
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
    fmt.Println("Give me the guild id")
    scanner.Scan()
    guildID = scanner.Text()
    guild, _ := dg.State.Guild(guildID)
    for i,channel := range guild.Channels {
      fmt.Printf("%d. %v \n", i, channel.Name)
    }

    fmt.Println("Give me the channel number")
    scanner.Scan()
    channelID := scanner.Text()
    channelIndex, _ := strconv.Atoi(channelID)
    channel := guild.Channels[channelIndex]
    for {
      fmt.Printf("Give me the message to send in %v ", channel.Name)
      scanner.Scan()
      content = scanner.Text()
      if content == "/exit" {
        break
      }
      dg.ChannelMessageSend(channel.ID, content)
    }
  }

}



func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
  fmt.Printf("\n%v said %v in %v\n", m.Author.Username, m.Content, m.GuildID)
}
