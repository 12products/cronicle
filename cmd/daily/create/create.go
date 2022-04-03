package create

import (
	"cronicle/utils"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "create a new daily entry",
		Long:  "create a new daily entry in your cronicle journal.",
		Run:   run,
	}

	cmd.Flags().StringP("message", "m", "", "content of your daily entry")
	cmd.Flags().StringP("tags", "t", "", "comma separated tags of your daily entry")

	return cmd
}

func run(cmd *cobra.Command, args []string) {
	m, _ := cmd.Flags().GetString("message")
	t, _ := cmd.Flags().GetString("tags")
	utils.WriteOrCreateDaily(utils.WriteDailyParams{Message: m, Tags: t})
}