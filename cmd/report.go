package cmd

import (
	"github.com/spf13/cobra"
	"github.com/ycanty/slack-stats/report"
)

func NewReportCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "report",
		Short: "Generate reports from the statistics database",
		Long:  ``,
	}

	command.AddCommand(NewReportOverviewCommand())

	// TODO Add --from <date> --to <date>

	return command
}

func reportApi() (*report.Api, error) {
	dbApi, err := dbApi()
	if err != nil {
		return nil, err
	}
	return report.NewApi(dbApi)
}
