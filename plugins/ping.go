package plugins

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/booyaa/phoenix"
)

type (
	PingPlugin struct {
		name        string
		version     string
		author      string
		description string
	}
)

func Ping() Plugin {
	plugin := PingPlugin{
		name:        "ping",
		version:     "0.0",
		author:      "booyaa",
		description: "ping tool",
	}
	return plugin
}

func (plugin PingPlugin) Name() string {
	return plugin.name
}

func (plugin PingPlugin) Version() string {
	return plugin.version
}

func (plugin PingPlugin) Author() string {
	return plugin.author
}

func (plugin PingPlugin) Description() string {
	return plugin.description
}

func (plugin PingPlugin) Handle(message *phoenix.Message) (string, error) {
	logger.WithFields(logrus.Fields{
		"token":       message.Token,
		"teamId":      message.TeamId,
		"channelId":   message.ChannelId,
		"channelName": message.ChannelName,
		"timestamp":   message.Timestamp,
		"userId":      message.UserId,
		"username":    message.Username,
		"text":        message.Text,
		"triggerWord": message.TriggerWord,
	}).Info("Ping Handler")

	host := message.Text
	if host == "" {
		return "provide a host", nil
	}

	return fmt.Sprintf("pinging %s", host), nil
}
