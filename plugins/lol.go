package plugins

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/booyaa/phoenix"
	"os/exec"
	"strings"
)

type (
	LolPlugin struct {
		name        string
		version     string
		author      string
		description string
	}
)

func Lol() Plugin {
	plugin := LolPlugin{
		name:        "lol",
		version:     "0.0",
		author:      "booyaa",
		description: "lol tool",
	}
	return plugin
}

func (plugin LolPlugin) Name() string {
	return plugin.name
}

func (plugin LolPlugin) Version() string {
	return plugin.version
}

func (plugin LolPlugin) Author() string {
	return plugin.author
}

func (plugin LolPlugin) Description() string {
	return plugin.description
}

func (plugin LolPlugin) Handle(message *phoenix.Message) (string, error) {
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
	}).Info("Lol Handler")

	data := message.Text
	if data == "" {
		return "provide a command", nil
	}

	command, args, err := parseCommand(data)
	output := ""
	if err != nil {
		output = fmt.Sprintf("Error parsing  %s %v", data, err)
	}

	// output = fmt.Sprintf("executing... %s %s", command, args)
	output = executeCommand(command, args)

	return output, nil
}

func parseCommand(message string) (string, string, error) {
	explode := strings.Split(message, " ")
	command := explode[0]
	args := strings.Join(explode[1:], " ")
	logger.WithFields(logrus.Fields{
		"message": message,
		"command": message, "args": args,
	}).Info("parsing command")
	return command, args, nil
}

func executeCommand(command string, args string) (output string) {
	var result []byte
	var err error
	if args == "" {
		result, err = exec.Command(command).CombinedOutput()
	} else {
		result, err = exec.Command(command, args).CombinedOutput()
	}
	output = string(result)

	if err != nil {
		output = fmt.Sprintf("Error occurred running %s: %s", command, result)
	}
	logger.WithFields(logrus.Fields{
		"command": command,
		"args":    args,
		"output":  output,
	}).Info("executeCommand")
	return string(output)
}
