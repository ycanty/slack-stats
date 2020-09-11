package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

func newFindChannelIDCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "find-channel-id",
		Short: "Find the channel ID from a channel name",
		Long:  ``,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			channels, err := Api().GetChannels(true)
			if err != nil {
				return err
			}

			type response struct {
				Name string
				ID   string
			}

			responses := make([]response, 0)
			for _, channel := range channels {
				if strings.Contains(channel.Name, args[0]) {
					responses = append(responses, response{
						Name: channel.Name,
						ID:   channel.ID,
					})
				}
			}
			jsonBytes, err := json.Marshal(responses)

			if err != nil {
				return err
			}
			fmt.Println(string(jsonBytes))

			return nil
		},
	}

	return command
}
