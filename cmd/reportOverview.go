package cmd

import (
	"github.com/spf13/cobra"
	"github.com/ycanty/slack-stats/json"
)

func NewReportOverviewCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "overview",
		Short: "Overview over a given period.",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			api, err := reportApi()
			if err != nil {
				return err
			}

			report, err := api.Overview()
			if err != nil {
				return err
			}
			return json.PrintJSON(cmd.OutOrStdout(), report)
		},
	}

	return command
}
