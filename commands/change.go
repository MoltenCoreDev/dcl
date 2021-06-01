package commands

import "github.com/bwmarrin/discordgo"

func ChangeGuild(currentGuild *string, id string) {
	*currentGuild = id
}

func ChangeChannel(currentChannel *string, id string) {
	*currentChannel = id
}

func SendMessage(session *discordgo.Session, channelID string, msgContent string) {
	session.ChannelMessageSend(channelID, msgContent)
}
